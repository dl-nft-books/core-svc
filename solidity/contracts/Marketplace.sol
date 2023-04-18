// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts-upgradeable/utils/cryptography/draft-EIP712Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

import "@dlsl/dev-modules/contracts-registry/AbstractDependant.sol";
import "@dlsl/dev-modules/utils/Globals.sol";
import "@dlsl/dev-modules/libs/decimals/DecimalsConverter.sol";
import "@dlsl/dev-modules/libs/arrays/Paginator.sol";

import "./interfaces/IMarketplace.sol";
import "./interfaces/IRoleManager.sol";
import "./interfaces/IContractsRegistry.sol";
import "./interfaces/ITokenFactory.sol";
import "./interfaces/tokens/IERC721MintableToken.sol";

contract Marketplace is
IMarketplace,
ERC721Holder,
AbstractDependant,
EIP712Upgradeable,
PausableUpgradeable
{
    using EnumerableSet for EnumerableSet.AddressSet;
    using Paginator for EnumerableSet.AddressSet;
    using DecimalsConverter for uint256;
    using SafeERC20 for IERC20Metadata;

    string public baseTokenContractsURI;

    bytes32 internal constant _BUY_TYPEHASH =
    keccak256(
        "Buy(address tokenContract,uint256 futureTokenId,address paymentTokenAddress,uint256 paymentTokenPrice,uint256 discount,uint256 endTimestamp,bytes32 tokenURI)"
    );
    bytes32 internal constant _BUY_WITH_REQUEST_TYPEHASH =
    keccak256(
        "BuyWithRequest(uint256 requestId,uint256 futureTokenId,uint256 endTimestamp,bytes32 tokenURI)"
    );
    bytes32 internal constant _BUY_WITH_VOUCHER_TYPEHASH =
    keccak256(
        "BuyWithVoucher(address requester,address tokenContract,uint256 futureTokenId,address voucherTokenContract,uint256 voucherTokensAmount,uint256 endTimestamp,bytes32 tokenURI)"
    );

    IRoleManager internal _roleManager;
    ITokenFactory internal _tokenFactory;

    EnumerableSet.AddressSet internal _tokenContracts;
    mapping(address => TokenParams) internal _tokenParams;
    NFTRequestInfo[] internal _nftRequests;

    modifier onlyMarketplaceManager() {
        _onlyMarketplaceManager();
        _;
    }

    modifier onlyWithdrawalManager() {
        _onlyWithdrawalManager();
        _;
    }

    function __Marketplace_init(
        string memory baseTokenContractsURI_
    ) external override initializer {
        __EIP712_init("Marketplace", "1");

        baseTokenContractsURI = baseTokenContractsURI_;
    }

    function setDependencies(
        address contractsRegistry_,
        bytes calldata
    ) external override dependant {
        IContractsRegistry registry_ = IContractsRegistry(contractsRegistry_);

        _roleManager = IRoleManager(registry_.getRoleManagerContract());
        _tokenFactory = ITokenFactory(registry_.getTokenFactoryContract());
    }

    function pause() external override onlyMarketplaceManager {
        _pause();
    }

    function unpause() external override onlyMarketplaceManager {
        _unpause();
    }

    function setBaseTokenContractsURI(
        string memory baseTokenContractsURI_
    ) external override whenNotPaused onlyMarketplaceManager {
        baseTokenContractsURI = baseTokenContractsURI_;

        emit BaseTokenContractsURIUpdated(baseTokenContractsURI_);
    }

    function addToken(
        string memory name_,
        string memory symbol_,
        TokenParams memory tokenParams_
    ) external override whenNotPaused onlyMarketplaceManager returns (address tokenProxy_) {
        _validateTokenParams(name_, symbol_);

        require(!tokenParams_.isDisabled, "Marketplace: Token can not be disabled on creation.");

        tokenProxy_ = _tokenFactory.deployToken(name_, symbol_);

        if (tokenParams_.isVoucherBuyable && tokenParams_.voucherTokenContract == address(0)) {
            tokenParams_.voucherTokenContract = _tokenFactory.deployVoucher(
                string.concat(name_, "_Voucher"),
                string.concat(symbol_, "_V")
            );
        }

        _validateVoucherParams(
            tokenParams_.isVoucherBuyable,
            tokenParams_.voucherTokensAmount,
            tokenParams_.voucherTokenContract
        );

        _tokenParams[tokenProxy_] = tokenParams_;

        _tokenContracts.add(tokenProxy_);

        emit TokenContractDeployed(tokenProxy_, name_, symbol_, tokenParams_);
    }

    function updateAllParams(
        address tokenContract_,
        string memory name_,
        string memory symbol_,
        TokenParams memory newTokenParams_
    ) external override whenNotPaused onlyMarketplaceManager {
        _checkTokenContractExists(tokenContract_);

        _validateTokenParams(name_, symbol_);
        _validateVoucherParams(
            newTokenParams_.isVoucherBuyable,
            newTokenParams_.voucherTokensAmount,
            newTokenParams_.voucherTokenContract
        );

        _tokenParams[tokenContract_] = newTokenParams_;

        IERC721MintableToken(tokenContract_).updateTokenParams(name_, symbol_);

        emit TokenContractParamsUpdated(tokenContract_, name_, symbol_, newTokenParams_);
    }

    function withdrawCurrency(
        address tokenAddr_,
        address recipient_,
        uint256 desiredAmount_,
        bool withdrawAll_
    ) external override onlyWithdrawalManager {
        IERC20Metadata token_ = IERC20Metadata(tokenAddr_);
        bool isNativeCurrency_ = tokenAddr_ == address(0);

        uint256 amount_ = isNativeCurrency_
        ? address(this).balance
        : token_.balanceOf(address(this));

        if (!withdrawAll_) {
            amount_ = Math.min(amount_, desiredAmount_);
        }

        require(amount_ > 0, "Marketplace: Nothing to withdraw.");

        if (isNativeCurrency_) {
            _sendNativeCurrency(recipient_, amount_);
        } else {
            token_.safeTransfer(recipient_, amount_);

            amount_ = amount_.to18(token_.decimals());
        }

        emit PaidTokensWithdrawn(tokenAddr_, recipient_, amount_);
    }

    function buyTokenWithETH(
        BuyParams memory buyParams_,
        Sig memory sig_
    ) external payable override whenNotPaused {
        _beforeBuyTokenCheck(buyParams_, sig_);

        require(
            buyParams_.paymentDetails.paymentTokenAddress == address(0),
            "Marketplace: Invalid payment token address"
        );

        TokenParams storage _currentTokenParams = _tokenParams[buyParams_.tokenContract];

        uint256 amountToPay_ = _getAmountToPay(
            buyParams_.paymentDetails,
            _currentTokenParams.pricePerOneToken
        );

        require(msg.value >= amountToPay_, "Marketplace: Invalid currency amount.");

        address fundsRecipient_ = _getFundsRecipient(_currentTokenParams.fundsRecipient);

        if (fundsRecipient_ != address(this)) {
            _sendNativeCurrency(fundsRecipient_, amountToPay_);
        }

        uint256 extraCurrencyAmount_ = msg.value - amountToPay_;

        if (extraCurrencyAmount_ > 0) {
            _sendNativeCurrency(msg.sender, extraCurrencyAmount_);
        }

        _mintToken(
            buyParams_,
            PaymentType.NATIVE,
            _currentTokenParams.pricePerOneToken,
            amountToPay_
        );
    }

    function buyTokenWithERC20(
        BuyParams memory buyParams_,
        Sig memory sig_
    ) external override whenNotPaused {
        _beforeBuyTokenCheck(buyParams_, sig_);

        TokenParams storage _currentTokenParams = _tokenParams[buyParams_.tokenContract];

        uint256 amountToPay_ = _getAmountToPay(
            buyParams_.paymentDetails,
            _currentTokenParams.pricePerOneToken
        );

        _sendERC20(
            IERC20Metadata(buyParams_.paymentDetails.paymentTokenAddress),
            msg.sender,
            _getFundsRecipient(_currentTokenParams.fundsRecipient),
            amountToPay_
        );

        _mintToken(
            buyParams_,
            PaymentType.ERC20,
            _currentTokenParams.pricePerOneToken,
            amountToPay_
        );
    }

    function buyTokenWithVoucher(
        BuyParams memory buyParams_,
        Sig memory sig_,
        Sig memory permitSig_
    ) external override whenNotPaused {
        _beforeBuyTokenCheck(buyParams_, sig_);

        TokenParams storage _currentTokenParams = _tokenParams[buyParams_.tokenContract];

        require(
            _currentTokenParams.isVoucherBuyable,
            "Marketplace: Unable to buy token with voucher"
        );
        require(
            buyParams_.paymentDetails.paymentTokenAddress ==
            _currentTokenParams.voucherTokenContract,
            "Marketplace: Invalid payment token address"
        );

        IERC20Permit(_currentTokenParams.voucherTokenContract).permit(
            buyParams_.recipient,
            address(this),
            _currentTokenParams.voucherTokensAmount,
            permitSig_.endTimestamp,
            permitSig_.v,
            permitSig_.r,
            permitSig_.s
        );

        _sendERC20(
            IERC20Metadata(_currentTokenParams.voucherTokenContract),
            buyParams_.recipient,
            _getFundsRecipient(_currentTokenParams.fundsRecipient),
            _currentTokenParams.voucherTokensAmount
        );

        _mintToken(
            buyParams_,
            PaymentType.VOUCHER,
            _currentTokenParams.pricePerOneToken,
            _currentTokenParams.voucherTokensAmount
        );
    }

    function buyTokenWithNFT(
        BuyParams memory buyParams_,
        Sig memory sig_
    ) external override whenNotPaused {
        _beforeBuyTokenCheck(buyParams_, sig_);

        TokenParams storage _currentTokenParams = _tokenParams[buyParams_.tokenContract];

        require(_currentTokenParams.isNFTBuyable, "Marketplace: Unable to buy token with NFT");

        require(
            buyParams_.paymentDetails.paymentTokenPrice >= _currentTokenParams.minNFTFloorPrice,
            "Marketplace: NFT floor price is less than the minimal."
        );

        _tranferNFT(
            buyParams_.paymentDetails.paymentTokenAddress,
            msg.sender,
            _getFundsRecipient(_currentTokenParams.fundsRecipient),
            buyParams_.paymentDetails.nftTokenId
        );

        _mintToken(buyParams_, PaymentType.NFT, _currentTokenParams.minNFTFloorPrice, 1);
    }

    function buyTokenWithRequest(
        RequestBuyParams memory requestBuyParams_,
        Sig memory sig_
    ) external override {
        _beforeBuyTokenWithRequestCheck(requestBuyParams_, sig_);

        NFTRequestInfo storage _nftRequest = _nftRequests[requestBuyParams_.requestId];

        TokenParams storage _currentTokenParams = _tokenParams[_nftRequest.tokenContract];

        address fundsRecipient_ = _getFundsRecipient(_currentTokenParams.fundsRecipient);
        if (fundsRecipient_ != address(this)) {
            _tranferNFT(
                _nftRequest.nftContract,
                address(this),
                fundsRecipient_,
                _nftRequest.nftId
            );
        }

        _updateRequestStatus(requestBuyParams_.requestId, NFTRequestStatus.MINTED);

        _mint(
            _nftRequest.tokenContract,
            msg.sender,
            requestBuyParams_.futureTokenId,
            requestBuyParams_.tokenURI
        );

        emit TokenSuccessfullyExchanged(msg.sender, requestBuyParams_, _nftRequest);
    }

    function createNFTRequest(
        address nftContract_,
        uint256 nftId_,
        address tokenContract_
    ) external override returns (uint256 requestId_) {
        _checkTokenContractExists(tokenContract_);

        TokenParams storage _currentTokenParams = _tokenParams[tokenContract_];

        require(!_currentTokenParams.isDisabled, "Marketplace: Token is disabled.");

        require(
            _currentTokenParams.isNFTBuyable,
            "Marketplace: This token cannot be purchased with NFT."
        );

        _tranferNFT(nftContract_, msg.sender, address(this), nftId_);

        requestId_ = _nftRequests.length;

        _nftRequests.push(
            NFTRequestInfo(
                tokenContract_,
                nftContract_,
                nftId_,
                msg.sender,
                NFTRequestStatus.PENDING
            )
        );

        emit NFTRequestCreated(requestId_, msg.sender, nftContract_, nftId_, tokenContract_);
    }

    function cancelNFTRequest(uint256 requestId_) external override {
        require(requestId_ < _nftRequests.length, "Marketplace: Request ID is not valid.");

        NFTRequestInfo storage _nftRequest = _nftRequests[requestId_];

        require(_nftRequest.requester == msg.sender, "Marketplace: Sender is not the requester.");

        _updateRequestStatus(requestId_, NFTRequestStatus.CANCELED);

        _tranferNFT(_nftRequest.nftContract, address(this), msg.sender, _nftRequest.nftId);

        emit NFTRequestCanceled(requestId_);
    }

    function getTokenContractsPart(
        uint256 offset_,
        uint256 limit_
    ) public view override returns (address[] memory) {
        return _tokenContracts.part(offset_, limit_);
    }

    function getBaseTokenParams(
        address[] memory tokenContract_
    ) public view override returns (BaseTokenParams[] memory baseTokenParams_) {
        baseTokenParams_ = new BaseTokenParams[](tokenContract_.length);
        for (uint256 i; i < tokenContract_.length; i++) {
            TokenParams memory _currentTokenParams = _tokenParams[tokenContract_[i]];
            baseTokenParams_[i] = BaseTokenParams(
                tokenContract_[i],
                _currentTokenParams.isDisabled,
                _currentTokenParams.pricePerOneToken,
                IERC721Metadata(tokenContract_[i]).name()
            );
        }
    }

    function getDetailedTokenParams(
        address[] memory tokenContracts_
    ) public view override returns (DetailedTokenParams[] memory detailedTokenParams_) {
        detailedTokenParams_ = new DetailedTokenParams[](tokenContracts_.length);

        for (uint256 i; i < tokenContracts_.length; i++) {
            detailedTokenParams_[i] = DetailedTokenParams(
                tokenContracts_[i],
                _tokenParams[tokenContracts_[i]],
                IERC721Metadata(tokenContracts_[i]).name(),
                IERC721Metadata(tokenContracts_[i]).symbol()
            );
        }
    }

    function getUserTokensPart(
        address userAddr_,
        uint256 offset_,
        uint256 limit_
    ) external view override returns (UserTokens[] memory userTokens_) {
        address[] memory _tokenContractsPart = _tokenContracts.part(offset_, limit_);

        userTokens_ = new UserTokens[](_tokenContractsPart.length);

        for (uint256 i = 0; i < _tokenContractsPart.length; i++) {
            userTokens_[i] = UserTokens(
                _tokenContractsPart[i],
                IERC721MintableToken(_tokenContractsPart[i]).getUserTokenIDs(userAddr_)
            );
        }
    }

    function getTokenContractsCount() external view override returns (uint256) {
        return _tokenContracts.length();
    }

    function getActiveTokenContractsCount() external view override returns (uint256 count_) {
        for (uint256 i = 0; i < _tokenContracts.length(); i++) {
            if (!_tokenParams[_tokenContracts.at(i)].isDisabled) {
                count_++;
            }
        }
    }

    function getBaseTokenParamsPart(
        uint256 offset_,
        uint256 limit_
    ) external view override returns (BaseTokenParams[] memory) {
        return getBaseTokenParams(getTokenContractsPart(offset_, limit_));
    }

    function getDetailedTokenParamsPart(
        uint256 offset_,
        uint256 limit_
    ) external view override returns (DetailedTokenParams[] memory) {
        return getDetailedTokenParams(getTokenContractsPart(offset_, limit_));
    }

    function getNFTRequestsCount() external view override returns (uint256) {
        return _nftRequests.length;
    }

    function getNFTRequestsPart(
        uint256 offset_,
        uint256 limit_
    ) external view override returns (NFTRequestInfo[] memory nftRequests_) {
        uint256 to_ = _handleIncomingParametersForPart(_nftRequests.length, offset_, limit_);
        nftRequests_ = new NFTRequestInfo[](to_ - offset_);

        for (uint256 i = offset_; i < to_; i++) {
            nftRequests_[i - offset_] = _nftRequests[i];
        }
    }

    function _sendNativeCurrency(address recipient_, uint256 amountToSend_) internal {
        (bool success_, ) = recipient_.call{value: amountToSend_}("");

        require(success_, "Marketplace: Failed to send currency to the recipient.");
    }

    function _sendERC20(
        IERC20Metadata token_,
        address sender_,
        address recipient_,
        uint256 amountToSend_
    ) internal {
        token_.safeTransferFrom(sender_, recipient_, amountToSend_.from18(token_.decimals()));
    }

    function _mintToken(
        BuyParams memory buyParams_,
        PaymentType paymentType_,
        uint256 pricePerOneToken_,
        uint256 paidTokensAmount_
    ) internal {
        _mint(buyParams_.tokenContract, msg.sender, buyParams_.futureTokenId, buyParams_.tokenURI);

        emit TokenSuccessfullyPurchased(
            msg.sender,
            pricePerOneToken_,
            paidTokensAmount_,
            buyParams_,
            paymentType_
        );
    }

    function _mint(
        address tokenContract_,
        address recipient_,
        uint256 tokenId_,
        string memory tokenURI_
    ) internal {
        IERC721MintableToken(tokenContract_).mint(recipient_, tokenId_, tokenURI_);
    }

    function _tranferNFT(
        address nftContract_,
        address from_,
        address to_,
        uint256 nftId_
    ) internal {
        IERC721 nft_ = IERC721(nftContract_);
        require(nft_.ownerOf(nftId_) == from_, "Marketplace: Sender is not the owner.");
        nft_.safeTransferFrom(from_, to_, nftId_);
    }

    function _updateRequestStatus(uint256 requestId_, NFTRequestStatus status_) internal {
        NFTRequestInfo storage _nftRequest = _nftRequests[requestId_];
        require(
            _nftRequest.status == NFTRequestStatus.PENDING,
            "Marketplace: Request status is not valid."
        );

        _nftRequest.status = status_;
    }

    function _beforeBuyTokenCheck(BuyParams memory buyParams_, Sig memory sig_) internal view {
        _checkTokenContractExists(buyParams_.tokenContract);

        TokenParams storage _currentTokenParams = _tokenParams[buyParams_.tokenContract];

        require(!_currentTokenParams.isDisabled, "Marketplace: Unable to buy disabled token");

        _verifySignature(
            sig_.endTimestamp,
            sig_,
            keccak256(
                abi.encode(
                    _BUY_TYPEHASH,
                    buyParams_.tokenContract,
                    buyParams_.futureTokenId,
                    buyParams_.paymentDetails.paymentTokenAddress,
                    buyParams_.paymentDetails.paymentTokenPrice,
                    buyParams_.paymentDetails.discount,
                    sig_.endTimestamp,
                    keccak256(abi.encodePacked(buyParams_.tokenURI))
                )
            )
        );
    }

    function _beforeBuyTokenWithRequestCheck(
        RequestBuyParams memory requestBuyParams_,
        Sig memory sig_
    ) internal view {
        require(
            requestBuyParams_.requestId < _nftRequests.length,
            "Marketplace: Request ID is not valid."
        );

        _verifySignature(
            sig_.endTimestamp,
            sig_,
            keccak256(
                abi.encode(
                    _BUY_WITH_REQUEST_TYPEHASH,
                    requestBuyParams_.requestId,
                    requestBuyParams_.futureTokenId,
                    sig_.endTimestamp,
                    keccak256(abi.encodePacked(requestBuyParams_.tokenURI))
                )
            )
        );

        NFTRequestInfo storage _nftRequest = _nftRequests[requestBuyParams_.requestId];

        require(_nftRequest.requester == msg.sender, "Marketplace: Sender is not the requester.");

        TokenParams storage _currentTokenParams = _tokenParams[_nftRequest.tokenContract];

        require(!_currentTokenParams.isDisabled, "Marketplace: Token is disabled.");

        require(
            _currentTokenParams.isNFTBuyable,
            "Marketplace: This token cannot be purchased with NFT."
        );
    }

    function _verifySignature(
        uint256 endTimestamp_,
        Sig memory sig_,
        bytes32 structHash_
    ) internal view {
        address signer_ = ECDSAUpgradeable.recover(
            _hashTypedDataV4(structHash_),
            sig_.v,
            sig_.r,
            sig_.s
        );

        require(_roleManager.isSignatureManager(signer_), "Marketplace: Invalid signature.");
        require(block.timestamp <= endTimestamp_, "Marketplace: Signature expired.");
    }

    function _getFundsRecipient(address fundsRecipient_) internal view returns (address) {
        return fundsRecipient_ == address(0) ? address(this) : fundsRecipient_;
    }

    function _checkTokenContractExists(address tokenContract_) internal view {
        require(
            _tokenContracts.contains(tokenContract_),
            "Marketplace: Token contract not found."
        );
    }

    function _onlyMarketplaceManager() internal view {
        require(
            _roleManager.isMarketplaceManager(msg.sender),
            "Marketplace: Caller is not a marketplace manager."
        );
    }

    function _onlyWithdrawalManager() internal view {
        require(
            _roleManager.isWithdrawalManager(msg.sender),
            "Marketplace: Caller is not a withdrawal manager."
        );
    }

    function _validateTokenParams(string memory name_, string memory symbol_) internal pure {
        require(
            bytes(name_).length > 0 && bytes(symbol_).length > 0,
            "Marketplace: Token name or symbol is empty."
        );
    }

    function _validateVoucherParams(
        bool isVoucherBuyable,
        uint256 voucherTokensAmount,
        address voucherAddress
    ) internal pure {
        require(
            !isVoucherBuyable || (voucherTokensAmount > 0 && voucherAddress != address(0)),
            "Marketplace: Invalid voucher params."
        );
    }

    function _getAmountToPay(
        PaymentDetails memory paymentDetails_,
        uint256 pricePerOneToken_
    ) internal pure returns (uint256) {
        uint256 amountWithoutDiscount_ = (pricePerOneToken_ * DECIMAL) /
        paymentDetails_.paymentTokenPrice;

        return
        (amountWithoutDiscount_ * (PERCENTAGE_100 - paymentDetails_.discount)) /
        PERCENTAGE_100;
    }

    function _handleIncomingParametersForPart(
        uint256 length_,
        uint256 offset_,
        uint256 limit_
    ) private pure returns (uint256 to_) {
        to_ = offset_ + limit_;

        if (to_ > length_) {
            to_ = length_;
        }

        if (offset_ > to_) {
            to_ = offset_;
        }
    }
}
