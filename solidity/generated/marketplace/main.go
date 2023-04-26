// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC721MintableTokenTokenMintData is an auto generated low-level Go binding around an user-defined struct.
type IERC721MintableTokenTokenMintData struct {
	TokenId  *big.Int
	TokenURI string
}

// IMarketplaceAcceptRequestParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceAcceptRequestParams struct {
	RequestId *big.Int
	Recipient common.Address
	TokenData IERC721MintableTokenTokenMintData
}

// IMarketplaceBaseTokenData is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBaseTokenData struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
}

// IMarketplaceBriefTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBriefTokenInfo struct {
	BaseTokenData    IMarketplaceBaseTokenData
	PricePerOneToken *big.Int
	IsDisabled       bool
}

// IMarketplaceBuyParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBuyParams struct {
	TokenContract  common.Address
	Recipient      common.Address
	PaymentDetails IMarketplacePaymentDetails
	TokenData      IERC721MintableTokenTokenMintData
}

// IMarketplaceDetailedTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceDetailedTokenInfo struct {
	BaseTokenData IMarketplaceBaseTokenData
	TokenParams   IMarketplaceTokenParams
}

// IMarketplaceNFTRequestInfo is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceNFTRequestInfo struct {
	Requester     common.Address
	TokenContract common.Address
	NftContract   common.Address
	NftId         *big.Int
	Status        uint8
}

// IMarketplacePaymentDetails is an auto generated low-level Go binding around an user-defined struct.
type IMarketplacePaymentDetails struct {
	PaymentTokenAddress common.Address
	PaymentTokenPrice   *big.Int
	Discount            *big.Int
	NftTokenId          *big.Int
}

// IMarketplaceSigData is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceSigData struct {
	EndSigTimestamp *big.Int
	R               [32]byte
	S               [32]byte
	V               uint8
}

// IMarketplaceTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceTokenParams struct {
	PricePerOneToken     *big.Int
	MinNFTFloorPrice     *big.Int
	VoucherTokensAmount  *big.Int
	VoucherTokenContract common.Address
	FundsRecipient       common.Address
	IsNFTBuyable         bool
	IsDisabled           bool
	IsVoucherBuyable     bool
}

// IMarketplaceUserTokens is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceUserTokens struct {
	TokenContract common.Address
	TokenIds      []*big.Int
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"NFTRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NFTRequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"}],\"name\":\"NFTTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.AcceptRequestParams\",\"name\":\"acceptRequestParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.NFTRequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.NFTRequestInfo\",\"name\":\"nftRequestInfo\",\"type\":\"tuple\"}],\"name\":\"TokenSuccessfullyExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedTokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidTokensAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumIMarketplace.PaymentType\",\"name\":\"paymentType\",\"type\":\"uint8\"}],\"name\":\"TokenSuccessfullyPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"__Marketplace_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.AcceptRequestParams\",\"name\":\"requestParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endSigTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.SigData\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"acceptRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams_\",\"type\":\"tuple\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenProxy_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endSigTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.SigData\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endSigTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.SigData\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endSigTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.SigData\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIERC721MintableToken.TokenMintData\",\"name\":\"tokenData\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endSigTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.SigData\",\"name\":\"sig_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endSigTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.SigData\",\"name\":\"permitSig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"cancelNFTRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"createNFTRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPendingRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContracts_\",\"type\":\"address[]\"}],\"name\":\"getBriefTokenInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenData\",\"name\":\"baseTokenData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.BriefTokenInfo[]\",\"name\":\"baseTokenParams_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getBriefTokenInfoPart\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenData\",\"name\":\"baseTokenData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.BriefTokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContracts_\",\"type\":\"address[]\"}],\"name\":\"getDetailedTokenInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenData\",\"name\":\"baseTokenData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.DetailedTokenInfo[]\",\"name\":\"detailedTokenParams_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getDetailedTokenInfoPart\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenData\",\"name\":\"baseTokenData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"internalType\":\"structIMarketplace.DetailedTokenInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestsId_\",\"type\":\"uint256[]\"}],\"name\":\"getNFTRequestsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.NFTRequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.NFTRequestInfo[]\",\"name\":\"nftRequestsInfo_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getPendingRequestsPart\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"}],\"name\":\"getUserPendingRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getUserPendingRequestsPart\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getUserTokensPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIMarketplace.UserTokens[]\",\"name\":\"userTokens_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractsRegistry_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"newTokenParams_\",\"type\":\"tuple\"}],\"name\":\"updateTokenParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"desiredAmount_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"withdrawAll_\",\"type\":\"bool\"}],\"name\":\"withdrawCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceCaller) BaseTokenContractsURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "baseTokenContractsURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceSession) BaseTokenContractsURI() (string, error) {
	return _Marketplace.Contract.BaseTokenContractsURI(&_Marketplace.CallOpts)
}

// BaseTokenContractsURI is a free data retrieval call binding the contract method 0x72148c1c.
//
// Solidity: function baseTokenContractsURI() view returns(string)
func (_Marketplace *MarketplaceCallerSession) BaseTokenContractsURI() (string, error) {
	return _Marketplace.Contract.BaseTokenContractsURI(&_Marketplace.CallOpts)
}

// GetActiveTokenContractsCount is a free data retrieval call binding the contract method 0x206a3da6.
//
// Solidity: function getActiveTokenContractsCount() view returns(uint256 count_)
func (_Marketplace *MarketplaceCaller) GetActiveTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getActiveTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveTokenContractsCount is a free data retrieval call binding the contract method 0x206a3da6.
//
// Solidity: function getActiveTokenContractsCount() view returns(uint256 count_)
func (_Marketplace *MarketplaceSession) GetActiveTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetActiveTokenContractsCount(&_Marketplace.CallOpts)
}

// GetActiveTokenContractsCount is a free data retrieval call binding the contract method 0x206a3da6.
//
// Solidity: function getActiveTokenContractsCount() view returns(uint256 count_)
func (_Marketplace *MarketplaceCallerSession) GetActiveTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetActiveTokenContractsCount(&_Marketplace.CallOpts)
}

// GetAllPendingRequestsCount is a free data retrieval call binding the contract method 0x41a0bc19.
//
// Solidity: function getAllPendingRequestsCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetAllPendingRequestsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getAllPendingRequestsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllPendingRequestsCount is a free data retrieval call binding the contract method 0x41a0bc19.
//
// Solidity: function getAllPendingRequestsCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetAllPendingRequestsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetAllPendingRequestsCount(&_Marketplace.CallOpts)
}

// GetAllPendingRequestsCount is a free data retrieval call binding the contract method 0x41a0bc19.
//
// Solidity: function getAllPendingRequestsCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetAllPendingRequestsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetAllPendingRequestsCount(&_Marketplace.CallOpts)
}

// GetBriefTokenInfo is a free data retrieval call binding the contract method 0x6bb299e3.
//
// Solidity: function getBriefTokenInfo(address[] tokenContracts_) view returns(((address,string,string),uint256,bool)[] baseTokenParams_)
func (_Marketplace *MarketplaceCaller) GetBriefTokenInfo(opts *bind.CallOpts, tokenContracts_ []common.Address) ([]IMarketplaceBriefTokenInfo, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBriefTokenInfo", tokenContracts_)

	if err != nil {
		return *new([]IMarketplaceBriefTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBriefTokenInfo)).(*[]IMarketplaceBriefTokenInfo)

	return out0, err

}

// GetBriefTokenInfo is a free data retrieval call binding the contract method 0x6bb299e3.
//
// Solidity: function getBriefTokenInfo(address[] tokenContracts_) view returns(((address,string,string),uint256,bool)[] baseTokenParams_)
func (_Marketplace *MarketplaceSession) GetBriefTokenInfo(tokenContracts_ []common.Address) ([]IMarketplaceBriefTokenInfo, error) {
	return _Marketplace.Contract.GetBriefTokenInfo(&_Marketplace.CallOpts, tokenContracts_)
}

// GetBriefTokenInfo is a free data retrieval call binding the contract method 0x6bb299e3.
//
// Solidity: function getBriefTokenInfo(address[] tokenContracts_) view returns(((address,string,string),uint256,bool)[] baseTokenParams_)
func (_Marketplace *MarketplaceCallerSession) GetBriefTokenInfo(tokenContracts_ []common.Address) ([]IMarketplaceBriefTokenInfo, error) {
	return _Marketplace.Contract.GetBriefTokenInfo(&_Marketplace.CallOpts, tokenContracts_)
}

// GetBriefTokenInfoPart is a free data retrieval call binding the contract method 0x6b7ac758.
//
// Solidity: function getBriefTokenInfoPart(uint256 offset_, uint256 limit_) view returns(((address,string,string),uint256,bool)[])
func (_Marketplace *MarketplaceCaller) GetBriefTokenInfoPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBriefTokenInfo, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBriefTokenInfoPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceBriefTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBriefTokenInfo)).(*[]IMarketplaceBriefTokenInfo)

	return out0, err

}

// GetBriefTokenInfoPart is a free data retrieval call binding the contract method 0x6b7ac758.
//
// Solidity: function getBriefTokenInfoPart(uint256 offset_, uint256 limit_) view returns(((address,string,string),uint256,bool)[])
func (_Marketplace *MarketplaceSession) GetBriefTokenInfoPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBriefTokenInfo, error) {
	return _Marketplace.Contract.GetBriefTokenInfoPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetBriefTokenInfoPart is a free data retrieval call binding the contract method 0x6b7ac758.
//
// Solidity: function getBriefTokenInfoPart(uint256 offset_, uint256 limit_) view returns(((address,string,string),uint256,bool)[])
func (_Marketplace *MarketplaceCallerSession) GetBriefTokenInfoPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBriefTokenInfo, error) {
	return _Marketplace.Contract.GetBriefTokenInfoPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetDetailedTokenInfo is a free data retrieval call binding the contract method 0x2195554e.
//
// Solidity: function getDetailedTokenInfo(address[] tokenContracts_) view returns(((address,string,string),(uint256,uint256,uint256,address,address,bool,bool,bool))[] detailedTokenParams_)
func (_Marketplace *MarketplaceCaller) GetDetailedTokenInfo(opts *bind.CallOpts, tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenInfo, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDetailedTokenInfo", tokenContracts_)

	if err != nil {
		return *new([]IMarketplaceDetailedTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceDetailedTokenInfo)).(*[]IMarketplaceDetailedTokenInfo)

	return out0, err

}

// GetDetailedTokenInfo is a free data retrieval call binding the contract method 0x2195554e.
//
// Solidity: function getDetailedTokenInfo(address[] tokenContracts_) view returns(((address,string,string),(uint256,uint256,uint256,address,address,bool,bool,bool))[] detailedTokenParams_)
func (_Marketplace *MarketplaceSession) GetDetailedTokenInfo(tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenInfo, error) {
	return _Marketplace.Contract.GetDetailedTokenInfo(&_Marketplace.CallOpts, tokenContracts_)
}

// GetDetailedTokenInfo is a free data retrieval call binding the contract method 0x2195554e.
//
// Solidity: function getDetailedTokenInfo(address[] tokenContracts_) view returns(((address,string,string),(uint256,uint256,uint256,address,address,bool,bool,bool))[] detailedTokenParams_)
func (_Marketplace *MarketplaceCallerSession) GetDetailedTokenInfo(tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenInfo, error) {
	return _Marketplace.Contract.GetDetailedTokenInfo(&_Marketplace.CallOpts, tokenContracts_)
}

// GetDetailedTokenInfoPart is a free data retrieval call binding the contract method 0x2b8ee271.
//
// Solidity: function getDetailedTokenInfoPart(uint256 offset_, uint256 limit_) view returns(((address,string,string),(uint256,uint256,uint256,address,address,bool,bool,bool))[])
func (_Marketplace *MarketplaceCaller) GetDetailedTokenInfoPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenInfo, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDetailedTokenInfoPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceDetailedTokenInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceDetailedTokenInfo)).(*[]IMarketplaceDetailedTokenInfo)

	return out0, err

}

// GetDetailedTokenInfoPart is a free data retrieval call binding the contract method 0x2b8ee271.
//
// Solidity: function getDetailedTokenInfoPart(uint256 offset_, uint256 limit_) view returns(((address,string,string),(uint256,uint256,uint256,address,address,bool,bool,bool))[])
func (_Marketplace *MarketplaceSession) GetDetailedTokenInfoPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenInfo, error) {
	return _Marketplace.Contract.GetDetailedTokenInfoPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetDetailedTokenInfoPart is a free data retrieval call binding the contract method 0x2b8ee271.
//
// Solidity: function getDetailedTokenInfoPart(uint256 offset_, uint256 limit_) view returns(((address,string,string),(uint256,uint256,uint256,address,address,bool,bool,bool))[])
func (_Marketplace *MarketplaceCallerSession) GetDetailedTokenInfoPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenInfo, error) {
	return _Marketplace.Contract.GetDetailedTokenInfoPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Marketplace *MarketplaceCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Marketplace *MarketplaceSession) GetInjector() (common.Address, error) {
	return _Marketplace.Contract.GetInjector(&_Marketplace.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Marketplace *MarketplaceCallerSession) GetInjector() (common.Address, error) {
	return _Marketplace.Contract.GetInjector(&_Marketplace.CallOpts)
}

// GetNFTRequestsInfo is a free data retrieval call binding the contract method 0x9458c67e.
//
// Solidity: function getNFTRequestsInfo(uint256[] requestsId_) view returns((address,address,address,uint256,uint8)[] nftRequestsInfo_)
func (_Marketplace *MarketplaceCaller) GetNFTRequestsInfo(opts *bind.CallOpts, requestsId_ []*big.Int) ([]IMarketplaceNFTRequestInfo, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getNFTRequestsInfo", requestsId_)

	if err != nil {
		return *new([]IMarketplaceNFTRequestInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceNFTRequestInfo)).(*[]IMarketplaceNFTRequestInfo)

	return out0, err

}

// GetNFTRequestsInfo is a free data retrieval call binding the contract method 0x9458c67e.
//
// Solidity: function getNFTRequestsInfo(uint256[] requestsId_) view returns((address,address,address,uint256,uint8)[] nftRequestsInfo_)
func (_Marketplace *MarketplaceSession) GetNFTRequestsInfo(requestsId_ []*big.Int) ([]IMarketplaceNFTRequestInfo, error) {
	return _Marketplace.Contract.GetNFTRequestsInfo(&_Marketplace.CallOpts, requestsId_)
}

// GetNFTRequestsInfo is a free data retrieval call binding the contract method 0x9458c67e.
//
// Solidity: function getNFTRequestsInfo(uint256[] requestsId_) view returns((address,address,address,uint256,uint8)[] nftRequestsInfo_)
func (_Marketplace *MarketplaceCallerSession) GetNFTRequestsInfo(requestsId_ []*big.Int) ([]IMarketplaceNFTRequestInfo, error) {
	return _Marketplace.Contract.GetNFTRequestsInfo(&_Marketplace.CallOpts, requestsId_)
}

// GetPendingRequestsPart is a free data retrieval call binding the contract method 0x5ef83bc5.
//
// Solidity: function getPendingRequestsPart(uint256 offset_, uint256 limit_) view returns(uint256[])
func (_Marketplace *MarketplaceCaller) GetPendingRequestsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getPendingRequestsPart", offset_, limit_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingRequestsPart is a free data retrieval call binding the contract method 0x5ef83bc5.
//
// Solidity: function getPendingRequestsPart(uint256 offset_, uint256 limit_) view returns(uint256[])
func (_Marketplace *MarketplaceSession) GetPendingRequestsPart(offset_ *big.Int, limit_ *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetPendingRequestsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetPendingRequestsPart is a free data retrieval call binding the contract method 0x5ef83bc5.
//
// Solidity: function getPendingRequestsPart(uint256 offset_, uint256 limit_) view returns(uint256[])
func (_Marketplace *MarketplaceCallerSession) GetPendingRequestsPart(offset_ *big.Int, limit_ *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetPendingRequestsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetTokenContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetTokenContractsCount(&_Marketplace.CallOpts)
}

// GetTokenContractsCount is a free data retrieval call binding the contract method 0x11b2bf29.
//
// Solidity: function getTokenContractsCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetTokenContractsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetTokenContractsCount(&_Marketplace.CallOpts)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCaller) GetTokenContractsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getTokenContractsPart", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetTokenContractsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetTokenContractsPart is a free data retrieval call binding the contract method 0xcc7c925a.
//
// Solidity: function getTokenContractsPart(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCallerSession) GetTokenContractsPart(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetTokenContractsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetUserPendingRequestsCount is a free data retrieval call binding the contract method 0xdbdb9635.
//
// Solidity: function getUserPendingRequestsCount(address userAddr_) view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetUserPendingRequestsCount(opts *bind.CallOpts, userAddr_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getUserPendingRequestsCount", userAddr_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserPendingRequestsCount is a free data retrieval call binding the contract method 0xdbdb9635.
//
// Solidity: function getUserPendingRequestsCount(address userAddr_) view returns(uint256)
func (_Marketplace *MarketplaceSession) GetUserPendingRequestsCount(userAddr_ common.Address) (*big.Int, error) {
	return _Marketplace.Contract.GetUserPendingRequestsCount(&_Marketplace.CallOpts, userAddr_)
}

// GetUserPendingRequestsCount is a free data retrieval call binding the contract method 0xdbdb9635.
//
// Solidity: function getUserPendingRequestsCount(address userAddr_) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetUserPendingRequestsCount(userAddr_ common.Address) (*big.Int, error) {
	return _Marketplace.Contract.GetUserPendingRequestsCount(&_Marketplace.CallOpts, userAddr_)
}

// GetUserPendingRequestsPart is a free data retrieval call binding the contract method 0x024a72bb.
//
// Solidity: function getUserPendingRequestsPart(address userAddr_, uint256 offset_, uint256 limit_) view returns(uint256[])
func (_Marketplace *MarketplaceCaller) GetUserPendingRequestsPart(opts *bind.CallOpts, userAddr_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getUserPendingRequestsPart", userAddr_, offset_, limit_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserPendingRequestsPart is a free data retrieval call binding the contract method 0x024a72bb.
//
// Solidity: function getUserPendingRequestsPart(address userAddr_, uint256 offset_, uint256 limit_) view returns(uint256[])
func (_Marketplace *MarketplaceSession) GetUserPendingRequestsPart(userAddr_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserPendingRequestsPart(&_Marketplace.CallOpts, userAddr_, offset_, limit_)
}

// GetUserPendingRequestsPart is a free data retrieval call binding the contract method 0x024a72bb.
//
// Solidity: function getUserPendingRequestsPart(address userAddr_, uint256 offset_, uint256 limit_) view returns(uint256[])
func (_Marketplace *MarketplaceCallerSession) GetUserPendingRequestsPart(userAddr_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]*big.Int, error) {
	return _Marketplace.Contract.GetUserPendingRequestsPart(&_Marketplace.CallOpts, userAddr_, offset_, limit_)
}

// GetUserTokensPart is a free data retrieval call binding the contract method 0x49049282.
//
// Solidity: function getUserTokensPart(address userAddr_, uint256 offset_, uint256 limit_) view returns((address,uint256[])[] userTokens_)
func (_Marketplace *MarketplaceCaller) GetUserTokensPart(opts *bind.CallOpts, userAddr_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceUserTokens, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getUserTokensPart", userAddr_, offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceUserTokens), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceUserTokens)).(*[]IMarketplaceUserTokens)

	return out0, err

}

// GetUserTokensPart is a free data retrieval call binding the contract method 0x49049282.
//
// Solidity: function getUserTokensPart(address userAddr_, uint256 offset_, uint256 limit_) view returns((address,uint256[])[] userTokens_)
func (_Marketplace *MarketplaceSession) GetUserTokensPart(userAddr_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceUserTokens, error) {
	return _Marketplace.Contract.GetUserTokensPart(&_Marketplace.CallOpts, userAddr_, offset_, limit_)
}

// GetUserTokensPart is a free data retrieval call binding the contract method 0x49049282.
//
// Solidity: function getUserTokensPart(address userAddr_, uint256 offset_, uint256 limit_) view returns((address,uint256[])[] userTokens_)
func (_Marketplace *MarketplaceCallerSession) GetUserTokensPart(userAddr_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceUserTokens, error) {
	return _Marketplace.Contract.GetUserTokensPart(&_Marketplace.CallOpts, userAddr_, offset_, limit_)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Marketplace *MarketplaceCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Marketplace *MarketplaceSession) NextRequestId() (*big.Int, error) {
	return _Marketplace.Contract.NextRequestId(&_Marketplace.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) NextRequestId() (*big.Int, error) {
	return _Marketplace.Contract.NextRequestId(&_Marketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceSession) Paused() (bool, error) {
	return _Marketplace.Contract.Paused(&_Marketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Marketplace *MarketplaceCallerSession) Paused() (bool, error) {
	return _Marketplace.Contract.Paused(&_Marketplace.CallOpts)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactor) MarketplaceInit(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "__Marketplace_init", baseTokenContractsURI_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceSession) MarketplaceInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xe709d541.
//
// Solidity: function __Marketplace_init(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactorSession) MarketplaceInit(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0xcd534ad3.
//
// Solidity: function acceptRequest((uint256,address,(uint256,string)) requestParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) AcceptRequest(opts *bind.TransactOpts, requestParams_ IMarketplaceAcceptRequestParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "acceptRequest", requestParams_, sig_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0xcd534ad3.
//
// Solidity: function acceptRequest((uint256,address,(uint256,string)) requestParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) AcceptRequest(requestParams_ IMarketplaceAcceptRequestParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.AcceptRequest(&_Marketplace.TransactOpts, requestParams_, sig_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0xcd534ad3.
//
// Solidity: function acceptRequest((uint256,address,(uint256,string)) requestParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) AcceptRequest(requestParams_ IMarketplaceAcceptRequestParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.AcceptRequest(&_Marketplace.TransactOpts, requestParams_, sig_)
}

// AddToken is a paid mutator transaction binding the contract method 0xa169fd46.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams_) returns(address tokenProxy_)
func (_Marketplace *MarketplaceTransactor) AddToken(opts *bind.TransactOpts, name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "addToken", name_, symbol_, tokenParams_)
}

// AddToken is a paid mutator transaction binding the contract method 0xa169fd46.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams_) returns(address tokenProxy_)
func (_Marketplace *MarketplaceSession) AddToken(name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.AddToken(&_Marketplace.TransactOpts, name_, symbol_, tokenParams_)
}

// AddToken is a paid mutator transaction binding the contract method 0xa169fd46.
//
// Solidity: function addToken(string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams_) returns(address tokenProxy_)
func (_Marketplace *MarketplaceTransactorSession) AddToken(name_ string, symbol_ string, tokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.AddToken(&_Marketplace.TransactOpts, name_, symbol_, tokenParams_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0xf68e448f.
//
// Solidity: function buyTokenWithERC20((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithERC20(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithERC20", buyParams_, sig_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0xf68e448f.
//
// Solidity: function buyTokenWithERC20((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithERC20(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithERC20(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0xf68e448f.
//
// Solidity: function buyTokenWithERC20((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithERC20(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithERC20(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0xfb278ce4.
//
// Solidity: function buyTokenWithETH((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithETH(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithETH", buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0xfb278ce4.
//
// Solidity: function buyTokenWithETH((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithETH(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithETH(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0xfb278ce4.
//
// Solidity: function buyTokenWithETH((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithETH(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithETH(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0xad2cf8af.
//
// Solidity: function buyTokenWithNFT((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithNFT(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithNFT", buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0xad2cf8af.
//
// Solidity: function buyTokenWithNFT((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithNFT(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithNFT(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0xad2cf8af.
//
// Solidity: function buyTokenWithNFT((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithNFT(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithNFT(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x1994c325.
//
// Solidity: function buyTokenWithVoucher((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_, (uint256,bytes32,bytes32,uint8) permitSig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithVoucher(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData, permitSig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithVoucher", buyParams_, sig_, permitSig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x1994c325.
//
// Solidity: function buyTokenWithVoucher((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_, (uint256,bytes32,bytes32,uint8) permitSig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithVoucher(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData, permitSig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithVoucher(&_Marketplace.TransactOpts, buyParams_, sig_, permitSig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x1994c325.
//
// Solidity: function buyTokenWithVoucher((address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams_, (uint256,bytes32,bytes32,uint8) sig_, (uint256,bytes32,bytes32,uint8) permitSig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithVoucher(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSigData, permitSig_ IMarketplaceSigData) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithVoucher(&_Marketplace.TransactOpts, buyParams_, sig_, permitSig_)
}

// CancelNFTRequest is a paid mutator transaction binding the contract method 0x5ec50810.
//
// Solidity: function cancelNFTRequest(uint256 requestId_) returns()
func (_Marketplace *MarketplaceTransactor) CancelNFTRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelNFTRequest", requestId_)
}

// CancelNFTRequest is a paid mutator transaction binding the contract method 0x5ec50810.
//
// Solidity: function cancelNFTRequest(uint256 requestId_) returns()
func (_Marketplace *MarketplaceSession) CancelNFTRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelNFTRequest(&_Marketplace.TransactOpts, requestId_)
}

// CancelNFTRequest is a paid mutator transaction binding the contract method 0x5ec50810.
//
// Solidity: function cancelNFTRequest(uint256 requestId_) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelNFTRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelNFTRequest(&_Marketplace.TransactOpts, requestId_)
}

// CreateNFTRequest is a paid mutator transaction binding the contract method 0x62002896.
//
// Solidity: function createNFTRequest(address tokenContract_, address nftContract_, uint256 nftId_) returns(uint256 requestId_)
func (_Marketplace *MarketplaceTransactor) CreateNFTRequest(opts *bind.TransactOpts, tokenContract_ common.Address, nftContract_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createNFTRequest", tokenContract_, nftContract_, nftId_)
}

// CreateNFTRequest is a paid mutator transaction binding the contract method 0x62002896.
//
// Solidity: function createNFTRequest(address tokenContract_, address nftContract_, uint256 nftId_) returns(uint256 requestId_)
func (_Marketplace *MarketplaceSession) CreateNFTRequest(tokenContract_ common.Address, nftContract_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateNFTRequest(&_Marketplace.TransactOpts, tokenContract_, nftContract_, nftId_)
}

// CreateNFTRequest is a paid mutator transaction binding the contract method 0x62002896.
//
// Solidity: function createNFTRequest(address tokenContract_, address nftContract_, uint256 nftId_) returns(uint256 requestId_)
func (_Marketplace *MarketplaceTransactorSession) CreateNFTRequest(tokenContract_ common.Address, nftContract_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateNFTRequest(&_Marketplace.TransactOpts, tokenContract_, nftContract_, nftId_)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Marketplace *MarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _Marketplace.Contract.Pause(&_Marketplace.TransactOpts)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactor) SetBaseTokenContractsURI(opts *bind.TransactOpts, baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setBaseTokenContractsURI", baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBaseTokenContractsURI(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// SetBaseTokenContractsURI is a paid mutator transaction binding the contract method 0x30dcd8aa.
//
// Solidity: function setBaseTokenContractsURI(string baseTokenContractsURI_) returns()
func (_Marketplace *MarketplaceTransactorSession) SetBaseTokenContractsURI(baseTokenContractsURI_ string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetBaseTokenContractsURI(&_Marketplace.TransactOpts, baseTokenContractsURI_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes ) returns()
func (_Marketplace *MarketplaceTransactor) SetDependencies(opts *bind.TransactOpts, contractsRegistry_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setDependencies", contractsRegistry_, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes ) returns()
func (_Marketplace *MarketplaceSession) SetDependencies(contractsRegistry_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SetDependencies(&_Marketplace.TransactOpts, contractsRegistry_, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address contractsRegistry_, bytes ) returns()
func (_Marketplace *MarketplaceTransactorSession) SetDependencies(contractsRegistry_ common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.SetDependencies(&_Marketplace.TransactOpts, contractsRegistry_, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Marketplace *MarketplaceTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Marketplace *MarketplaceSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetInjector(&_Marketplace.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Marketplace *MarketplaceTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetInjector(&_Marketplace.TransactOpts, injector_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Marketplace *MarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _Marketplace.Contract.Unpause(&_Marketplace.TransactOpts)
}

// UpdateTokenParams is a paid mutator transaction binding the contract method 0x7b68e174.
//
// Solidity: function updateTokenParams(address tokenContract_, (uint256,uint256,uint256,address,address,bool,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactor) UpdateTokenParams(opts *bind.TransactOpts, tokenContract_ common.Address, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateTokenParams", tokenContract_, newTokenParams_)
}

// UpdateTokenParams is a paid mutator transaction binding the contract method 0x7b68e174.
//
// Solidity: function updateTokenParams(address tokenContract_, (uint256,uint256,uint256,address,address,bool,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceSession) UpdateTokenParams(tokenContract_ common.Address, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateTokenParams(&_Marketplace.TransactOpts, tokenContract_, newTokenParams_)
}

// UpdateTokenParams is a paid mutator transaction binding the contract method 0x7b68e174.
//
// Solidity: function updateTokenParams(address tokenContract_, (uint256,uint256,uint256,address,address,bool,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateTokenParams(tokenContract_ common.Address, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateTokenParams(&_Marketplace.TransactOpts, tokenContract_, newTokenParams_)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x8099d792.
//
// Solidity: function withdrawCurrency(address tokenAddr_, address recipient_, uint256 desiredAmount_, bool withdrawAll_) returns()
func (_Marketplace *MarketplaceTransactor) WithdrawCurrency(opts *bind.TransactOpts, tokenAddr_ common.Address, recipient_ common.Address, desiredAmount_ *big.Int, withdrawAll_ bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdrawCurrency", tokenAddr_, recipient_, desiredAmount_, withdrawAll_)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x8099d792.
//
// Solidity: function withdrawCurrency(address tokenAddr_, address recipient_, uint256 desiredAmount_, bool withdrawAll_) returns()
func (_Marketplace *MarketplaceSession) WithdrawCurrency(tokenAddr_ common.Address, recipient_ common.Address, desiredAmount_ *big.Int, withdrawAll_ bool) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawCurrency(&_Marketplace.TransactOpts, tokenAddr_, recipient_, desiredAmount_, withdrawAll_)
}

// WithdrawCurrency is a paid mutator transaction binding the contract method 0x8099d792.
//
// Solidity: function withdrawCurrency(address tokenAddr_, address recipient_, uint256 desiredAmount_, bool withdrawAll_) returns()
func (_Marketplace *MarketplaceTransactorSession) WithdrawCurrency(tokenAddr_ common.Address, recipient_ common.Address, desiredAmount_ *big.Int, withdrawAll_ bool) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawCurrency(&_Marketplace.TransactOpts, tokenAddr_, recipient_, desiredAmount_, withdrawAll_)
}

// WithdrawNFTs is a paid mutator transaction binding the contract method 0xa684e1fd.
//
// Solidity: function withdrawNFTs(address nft_, address recipient_, uint256[] tokenIds_) returns()
func (_Marketplace *MarketplaceTransactor) WithdrawNFTs(opts *bind.TransactOpts, nft_ common.Address, recipient_ common.Address, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdrawNFTs", nft_, recipient_, tokenIds_)
}

// WithdrawNFTs is a paid mutator transaction binding the contract method 0xa684e1fd.
//
// Solidity: function withdrawNFTs(address nft_, address recipient_, uint256[] tokenIds_) returns()
func (_Marketplace *MarketplaceSession) WithdrawNFTs(nft_ common.Address, recipient_ common.Address, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawNFTs(&_Marketplace.TransactOpts, nft_, recipient_, tokenIds_)
}

// WithdrawNFTs is a paid mutator transaction binding the contract method 0xa684e1fd.
//
// Solidity: function withdrawNFTs(address nft_, address recipient_, uint256[] tokenIds_) returns()
func (_Marketplace *MarketplaceTransactorSession) WithdrawNFTs(nft_ common.Address, recipient_ common.Address, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawNFTs(&_Marketplace.TransactOpts, nft_, recipient_, tokenIds_)
}

// MarketplaceBaseTokenContractsURIUpdatedIterator is returned from FilterBaseTokenContractsURIUpdated and is used to iterate over the raw logs and unpacked data for BaseTokenContractsURIUpdated events raised by the Marketplace contract.
type MarketplaceBaseTokenContractsURIUpdatedIterator struct {
	Event *MarketplaceBaseTokenContractsURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceBaseTokenContractsURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceBaseTokenContractsURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceBaseTokenContractsURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceBaseTokenContractsURIUpdated represents a BaseTokenContractsURIUpdated event raised by the Marketplace contract.
type MarketplaceBaseTokenContractsURIUpdated struct {
	NewBaseTokenContractsURI string
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenContractsURIUpdated is a free log retrieval operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) FilterBaseTokenContractsURIUpdated(opts *bind.FilterOpts) (*MarketplaceBaseTokenContractsURIUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceBaseTokenContractsURIUpdatedIterator{contract: _Marketplace.contract, event: "BaseTokenContractsURIUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseTokenContractsURIUpdated is a free log subscription operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) WatchBaseTokenContractsURIUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceBaseTokenContractsURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "BaseTokenContractsURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceBaseTokenContractsURIUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaseTokenContractsURIUpdated is a log parse operation binding the contract event 0x35221ce24b65bb230797e03080154779f50580793d0f60505bb8a6d15c507b80.
//
// Solidity: event BaseTokenContractsURIUpdated(string newBaseTokenContractsURI)
func (_Marketplace *MarketplaceFilterer) ParseBaseTokenContractsURIUpdated(log types.Log) (*MarketplaceBaseTokenContractsURIUpdated, error) {
	event := new(MarketplaceBaseTokenContractsURIUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "BaseTokenContractsURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTRequestCanceledIterator is returned from FilterNFTRequestCanceled and is used to iterate over the raw logs and unpacked data for NFTRequestCanceled events raised by the Marketplace contract.
type MarketplaceNFTRequestCanceledIterator struct {
	Event *MarketplaceNFTRequestCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceNFTRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTRequestCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceNFTRequestCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceNFTRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTRequestCanceled represents a NFTRequestCanceled event raised by the Marketplace contract.
type MarketplaceNFTRequestCanceled struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTRequestCanceled is a free log retrieval operation binding the contract event 0x4c7582395e3269090aa4d10a422ccbc61f8704f6c3175ea9e6a848570bd0cf50.
//
// Solidity: event NFTRequestCanceled(uint256 indexed requestId)
func (_Marketplace *MarketplaceFilterer) FilterNFTRequestCanceled(opts *bind.FilterOpts, requestId []*big.Int) (*MarketplaceNFTRequestCanceledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTRequestCanceled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTRequestCanceledIterator{contract: _Marketplace.contract, event: "NFTRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchNFTRequestCanceled is a free log subscription operation binding the contract event 0x4c7582395e3269090aa4d10a422ccbc61f8704f6c3175ea9e6a848570bd0cf50.
//
// Solidity: event NFTRequestCanceled(uint256 indexed requestId)
func (_Marketplace *MarketplaceFilterer) WatchNFTRequestCanceled(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTRequestCanceled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTRequestCanceled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTRequestCanceled)
				if err := _Marketplace.contract.UnpackLog(event, "NFTRequestCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTRequestCanceled is a log parse operation binding the contract event 0x4c7582395e3269090aa4d10a422ccbc61f8704f6c3175ea9e6a848570bd0cf50.
//
// Solidity: event NFTRequestCanceled(uint256 indexed requestId)
func (_Marketplace *MarketplaceFilterer) ParseNFTRequestCanceled(log types.Log) (*MarketplaceNFTRequestCanceled, error) {
	event := new(MarketplaceNFTRequestCanceled)
	if err := _Marketplace.contract.UnpackLog(event, "NFTRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTRequestCreatedIterator is returned from FilterNFTRequestCreated and is used to iterate over the raw logs and unpacked data for NFTRequestCreated events raised by the Marketplace contract.
type MarketplaceNFTRequestCreatedIterator struct {
	Event *MarketplaceNFTRequestCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceNFTRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTRequestCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceNFTRequestCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceNFTRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTRequestCreated represents a NFTRequestCreated event raised by the Marketplace contract.
type MarketplaceNFTRequestCreated struct {
	RequestId     *big.Int
	Requester     common.Address
	TokenContract common.Address
	NftContract   common.Address
	NftId         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNFTRequestCreated is a free log retrieval operation binding the contract event 0x3f8f712ef49e6784ca3e1cf25f4e7d3876b97ff96c11c24926d7164dabf246b7.
//
// Solidity: event NFTRequestCreated(uint256 indexed requestId, address indexed requester, address indexed tokenContract, address nftContract, uint256 nftId)
func (_Marketplace *MarketplaceFilterer) FilterNFTRequestCreated(opts *bind.FilterOpts, requestId []*big.Int, requester []common.Address, tokenContract []common.Address) (*MarketplaceNFTRequestCreatedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTRequestCreated", requestIdRule, requesterRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTRequestCreatedIterator{contract: _Marketplace.contract, event: "NFTRequestCreated", logs: logs, sub: sub}, nil
}

// WatchNFTRequestCreated is a free log subscription operation binding the contract event 0x3f8f712ef49e6784ca3e1cf25f4e7d3876b97ff96c11c24926d7164dabf246b7.
//
// Solidity: event NFTRequestCreated(uint256 indexed requestId, address indexed requester, address indexed tokenContract, address nftContract, uint256 nftId)
func (_Marketplace *MarketplaceFilterer) WatchNFTRequestCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTRequestCreated, requestId []*big.Int, requester []common.Address, tokenContract []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTRequestCreated", requestIdRule, requesterRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTRequestCreated)
				if err := _Marketplace.contract.UnpackLog(event, "NFTRequestCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTRequestCreated is a log parse operation binding the contract event 0x3f8f712ef49e6784ca3e1cf25f4e7d3876b97ff96c11c24926d7164dabf246b7.
//
// Solidity: event NFTRequestCreated(uint256 indexed requestId, address indexed requester, address indexed tokenContract, address nftContract, uint256 nftId)
func (_Marketplace *MarketplaceFilterer) ParseNFTRequestCreated(log types.Log) (*MarketplaceNFTRequestCreated, error) {
	event := new(MarketplaceNFTRequestCreated)
	if err := _Marketplace.contract.UnpackLog(event, "NFTRequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTTokensWithdrawnIterator is returned from FilterNFTTokensWithdrawn and is used to iterate over the raw logs and unpacked data for NFTTokensWithdrawn events raised by the Marketplace contract.
type MarketplaceNFTTokensWithdrawnIterator struct {
	Event *MarketplaceNFTTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceNFTTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceNFTTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceNFTTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTTokensWithdrawn represents a NFTTokensWithdrawn event raised by the Marketplace contract.
type MarketplaceNFTTokensWithdrawn struct {
	NftAddr   common.Address
	Recipient common.Address
	TokenIDs  []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTTokensWithdrawn is a free log retrieval operation binding the contract event 0x532e4a6fa1bc2d0da3910834e7b66e23cfac61111db6e1b37b4986cb16ba4614.
//
// Solidity: event NFTTokensWithdrawn(address indexed nftAddr, address recipient, uint256[] tokenIDs)
func (_Marketplace *MarketplaceFilterer) FilterNFTTokensWithdrawn(opts *bind.FilterOpts, nftAddr []common.Address) (*MarketplaceNFTTokensWithdrawnIterator, error) {

	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTTokensWithdrawn", nftAddrRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTTokensWithdrawnIterator{contract: _Marketplace.contract, event: "NFTTokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNFTTokensWithdrawn is a free log subscription operation binding the contract event 0x532e4a6fa1bc2d0da3910834e7b66e23cfac61111db6e1b37b4986cb16ba4614.
//
// Solidity: event NFTTokensWithdrawn(address indexed nftAddr, address recipient, uint256[] tokenIDs)
func (_Marketplace *MarketplaceFilterer) WatchNFTTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTTokensWithdrawn, nftAddr []common.Address) (event.Subscription, error) {

	var nftAddrRule []interface{}
	for _, nftAddrItem := range nftAddr {
		nftAddrRule = append(nftAddrRule, nftAddrItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTTokensWithdrawn", nftAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTTokensWithdrawn)
				if err := _Marketplace.contract.UnpackLog(event, "NFTTokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTTokensWithdrawn is a log parse operation binding the contract event 0x532e4a6fa1bc2d0da3910834e7b66e23cfac61111db6e1b37b4986cb16ba4614.
//
// Solidity: event NFTTokensWithdrawn(address indexed nftAddr, address recipient, uint256[] tokenIDs)
func (_Marketplace *MarketplaceFilterer) ParseNFTTokensWithdrawn(log types.Log) (*MarketplaceNFTTokensWithdrawn, error) {
	event := new(MarketplaceNFTTokensWithdrawn)
	if err := _Marketplace.contract.UnpackLog(event, "NFTTokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePaidTokensWithdrawnIterator is returned from FilterPaidTokensWithdrawn and is used to iterate over the raw logs and unpacked data for PaidTokensWithdrawn events raised by the Marketplace contract.
type MarketplacePaidTokensWithdrawnIterator struct {
	Event *MarketplacePaidTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplacePaidTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePaidTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplacePaidTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplacePaidTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePaidTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePaidTokensWithdrawn represents a PaidTokensWithdrawn event raised by the Marketplace contract.
type MarketplacePaidTokensWithdrawn struct {
	TokenAddr common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaidTokensWithdrawn is a free log retrieval operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) FilterPaidTokensWithdrawn(opts *bind.FilterOpts, tokenAddr []common.Address) (*MarketplacePaidTokensWithdrawnIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePaidTokensWithdrawnIterator{contract: _Marketplace.contract, event: "PaidTokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPaidTokensWithdrawn is a free log subscription operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) WatchPaidTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MarketplacePaidTokensWithdrawn, tokenAddr []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PaidTokensWithdrawn", tokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePaidTokensWithdrawn)
				if err := _Marketplace.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaidTokensWithdrawn is a log parse operation binding the contract event 0xae2408273a2b19d84df6e31c6b33cda24768f5e01118d06cbe3b098e231868f1.
//
// Solidity: event PaidTokensWithdrawn(address indexed tokenAddr, address recipient, uint256 amount)
func (_Marketplace *MarketplaceFilterer) ParsePaidTokensWithdrawn(log types.Log) (*MarketplacePaidTokensWithdrawn, error) {
	event := new(MarketplacePaidTokensWithdrawn)
	if err := _Marketplace.contract.UnpackLog(event, "PaidTokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Marketplace contract.
type MarketplacePausedIterator struct {
	Event *MarketplacePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplacePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePaused represents a Paused event raised by the Marketplace contract.
type MarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketplacePausedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketplacePausedIterator{contract: _Marketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePaused)
				if err := _Marketplace.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Marketplace *MarketplaceFilterer) ParsePaused(log types.Log) (*MarketplacePaused, error) {
	event := new(MarketplacePaused)
	if err := _Marketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenContractDeployedIterator is returned from FilterTokenContractDeployed and is used to iterate over the raw logs and unpacked data for TokenContractDeployed events raised by the Marketplace contract.
type MarketplaceTokenContractDeployedIterator struct {
	Event *MarketplaceTokenContractDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenContractDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenContractDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenContractDeployed represents a TokenContractDeployed event raised by the Marketplace contract.
type MarketplaceTokenContractDeployed struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenContractDeployed is a free log retrieval operation binding the contract event 0x8bc16697985bc295d4dca70f21f8acd84cd075238dc1fabd9bb9a563cf3d2a18.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenContractDeployed(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenContractDeployedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenContractDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenContractDeployedIterator{contract: _Marketplace.contract, event: "TokenContractDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenContractDeployed is a free log subscription operation binding the contract event 0x8bc16697985bc295d4dca70f21f8acd84cd075238dc1fabd9bb9a563cf3d2a18.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenContractDeployed(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenContractDeployed, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenContractDeployed", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenContractDeployed)
				if err := _Marketplace.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenContractDeployed is a log parse operation binding the contract event 0x8bc16697985bc295d4dca70f21f8acd84cd075238dc1fabd9bb9a563cf3d2a18.
//
// Solidity: event TokenContractDeployed(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenContractDeployed(log types.Log) (*MarketplaceTokenContractDeployed, error) {
	event := new(MarketplaceTokenContractDeployed)
	if err := _Marketplace.contract.UnpackLog(event, "TokenContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenParamsUpdatedIterator is returned from FilterTokenParamsUpdated and is used to iterate over the raw logs and unpacked data for TokenParamsUpdated events raised by the Marketplace contract.
type MarketplaceTokenParamsUpdatedIterator struct {
	Event *MarketplaceTokenParamsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenParamsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenParamsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenParamsUpdated represents a TokenParamsUpdated event raised by the Marketplace contract.
type MarketplaceTokenParamsUpdated struct {
	TokenContract common.Address
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenParamsUpdated is a free log retrieval operation binding the contract event 0xd37221c35049c28d88e4fcc76bda0a4ede57343f131a0ebf1c3bca09baf6f1e1.
//
// Solidity: event TokenParamsUpdated(address indexed tokenContract, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenParamsUpdated(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenParamsUpdatedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenParamsUpdatedIterator{contract: _Marketplace.contract, event: "TokenParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenParamsUpdated is a free log subscription operation binding the contract event 0xd37221c35049c28d88e4fcc76bda0a4ede57343f131a0ebf1c3bca09baf6f1e1.
//
// Solidity: event TokenParamsUpdated(address indexed tokenContract, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenParamsUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenParamsUpdated, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenParamsUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "TokenParamsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenParamsUpdated is a log parse operation binding the contract event 0xd37221c35049c28d88e4fcc76bda0a4ede57343f131a0ebf1c3bca09baf6f1e1.
//
// Solidity: event TokenParamsUpdated(address indexed tokenContract, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenParamsUpdated(log types.Log) (*MarketplaceTokenParamsUpdated, error) {
	event := new(MarketplaceTokenParamsUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "TokenParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenSuccessfullyExchangedIterator is returned from FilterTokenSuccessfullyExchanged and is used to iterate over the raw logs and unpacked data for TokenSuccessfullyExchanged events raised by the Marketplace contract.
type MarketplaceTokenSuccessfullyExchangedIterator struct {
	Event *MarketplaceTokenSuccessfullyExchanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenSuccessfullyExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenSuccessfullyExchanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenSuccessfullyExchanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenSuccessfullyExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenSuccessfullyExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenSuccessfullyExchanged represents a TokenSuccessfullyExchanged event raised by the Marketplace contract.
type MarketplaceTokenSuccessfullyExchanged struct {
	AcceptRequestParams IMarketplaceAcceptRequestParams
	NftRequestInfo      IMarketplaceNFTRequestInfo
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokenSuccessfullyExchanged is a free log retrieval operation binding the contract event 0xeb314bbb2d6fa7e2d75c3225d2d867d89b3dba641c935cced0b29b006c1323b0.
//
// Solidity: event TokenSuccessfullyExchanged((uint256,address,(uint256,string)) acceptRequestParams, (address,address,address,uint256,uint8) nftRequestInfo)
func (_Marketplace *MarketplaceFilterer) FilterTokenSuccessfullyExchanged(opts *bind.FilterOpts) (*MarketplaceTokenSuccessfullyExchangedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenSuccessfullyExchanged")
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenSuccessfullyExchangedIterator{contract: _Marketplace.contract, event: "TokenSuccessfullyExchanged", logs: logs, sub: sub}, nil
}

// WatchTokenSuccessfullyExchanged is a free log subscription operation binding the contract event 0xeb314bbb2d6fa7e2d75c3225d2d867d89b3dba641c935cced0b29b006c1323b0.
//
// Solidity: event TokenSuccessfullyExchanged((uint256,address,(uint256,string)) acceptRequestParams, (address,address,address,uint256,uint8) nftRequestInfo)
func (_Marketplace *MarketplaceFilterer) WatchTokenSuccessfullyExchanged(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenSuccessfullyExchanged) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenSuccessfullyExchanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenSuccessfullyExchanged)
				if err := _Marketplace.contract.UnpackLog(event, "TokenSuccessfullyExchanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSuccessfullyExchanged is a log parse operation binding the contract event 0xeb314bbb2d6fa7e2d75c3225d2d867d89b3dba641c935cced0b29b006c1323b0.
//
// Solidity: event TokenSuccessfullyExchanged((uint256,address,(uint256,string)) acceptRequestParams, (address,address,address,uint256,uint8) nftRequestInfo)
func (_Marketplace *MarketplaceFilterer) ParseTokenSuccessfullyExchanged(log types.Log) (*MarketplaceTokenSuccessfullyExchanged, error) {
	event := new(MarketplaceTokenSuccessfullyExchanged)
	if err := _Marketplace.contract.UnpackLog(event, "TokenSuccessfullyExchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceTokenSuccessfullyPurchasedIterator is returned from FilterTokenSuccessfullyPurchased and is used to iterate over the raw logs and unpacked data for TokenSuccessfullyPurchased events raised by the Marketplace contract.
type MarketplaceTokenSuccessfullyPurchasedIterator struct {
	Event *MarketplaceTokenSuccessfullyPurchased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceTokenSuccessfullyPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenSuccessfullyPurchased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceTokenSuccessfullyPurchased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceTokenSuccessfullyPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenSuccessfullyPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenSuccessfullyPurchased represents a TokenSuccessfullyPurchased event raised by the Marketplace contract.
type MarketplaceTokenSuccessfullyPurchased struct {
	MintedTokenPrice *big.Int
	PaidTokensAmount *big.Int
	BuyParams        IMarketplaceBuyParams
	PaymentType      uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenSuccessfullyPurchased is a free log retrieval operation binding the contract event 0x693a26d2a31857c9285cd7d58481fcd72b3b1dd2f2a60b0e990c5d15bab6b8b9.
//
// Solidity: event TokenSuccessfullyPurchased(uint256 mintedTokenPrice, uint256 paidTokensAmount, (address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) FilterTokenSuccessfullyPurchased(opts *bind.FilterOpts) (*MarketplaceTokenSuccessfullyPurchasedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenSuccessfullyPurchased")
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenSuccessfullyPurchasedIterator{contract: _Marketplace.contract, event: "TokenSuccessfullyPurchased", logs: logs, sub: sub}, nil
}

// WatchTokenSuccessfullyPurchased is a free log subscription operation binding the contract event 0x693a26d2a31857c9285cd7d58481fcd72b3b1dd2f2a60b0e990c5d15bab6b8b9.
//
// Solidity: event TokenSuccessfullyPurchased(uint256 mintedTokenPrice, uint256 paidTokensAmount, (address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) WatchTokenSuccessfullyPurchased(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenSuccessfullyPurchased) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenSuccessfullyPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenSuccessfullyPurchased)
				if err := _Marketplace.contract.UnpackLog(event, "TokenSuccessfullyPurchased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSuccessfullyPurchased is a log parse operation binding the contract event 0x693a26d2a31857c9285cd7d58481fcd72b3b1dd2f2a60b0e990c5d15bab6b8b9.
//
// Solidity: event TokenSuccessfullyPurchased(uint256 mintedTokenPrice, uint256 paidTokensAmount, (address,address,(address,uint256,uint256,uint256),(uint256,string)) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) ParseTokenSuccessfullyPurchased(log types.Log) (*MarketplaceTokenSuccessfullyPurchased, error) {
	event := new(MarketplaceTokenSuccessfullyPurchased)
	if err := _Marketplace.contract.UnpackLog(event, "TokenSuccessfullyPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Marketplace contract.
type MarketplaceUnpausedIterator struct {
	Event *MarketplaceUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceUnpaused represents a Unpaused event raised by the Marketplace contract.
type MarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketplaceUnpausedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketplaceUnpausedIterator{contract: _Marketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceUnpaused)
				if err := _Marketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Marketplace *MarketplaceFilterer) ParseUnpaused(log types.Log) (*MarketplaceUnpaused, error) {
	event := new(MarketplaceUnpaused)
	if err := _Marketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
