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

// IMarketplaceBaseTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBaseTokenParams struct {
	TokenContract    common.Address
	IsDisabled       bool
	PricePerOneToken *big.Int
	TokenName        string
}

// IMarketplaceBuyParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceBuyParams struct {
	PaymentDetails IMarketplacePaymentDetails
	TokenContract  common.Address
	Recipient      common.Address
	FutureTokenId  *big.Int
	TokenURI       string
}

// IMarketplaceDetailedTokenParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceDetailedTokenParams struct {
	TokenContract common.Address
	TokenParams   IMarketplaceTokenParams
	TokenName     string
	TokenSymbol   string
}

// IMarketplaceNFTRequestInfo is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceNFTRequestInfo struct {
	TokenContract common.Address
	NftContract   common.Address
	NftId         *big.Int
	Requester     common.Address
	Status        uint8
}

// IMarketplacePaymentDetails is an auto generated low-level Go binding around an user-defined struct.
type IMarketplacePaymentDetails struct {
	PaymentTokenAddress common.Address
	PaymentTokenPrice   *big.Int
	Discount            *big.Int
	NftTokenId          *big.Int
}

// IMarketplaceRequestBuyParams is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceRequestBuyParams struct {
	RequestId     *big.Int
	FutureTokenId *big.Int
	TokenURI      string
}

// IMarketplaceSig is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceSig struct {
	EndTimestamp *big.Int
	R            [32]byte
	S            [32]byte
	V            uint8
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newBaseTokenContractsURI\",\"type\":\"string\"}],\"name\":\"BaseTokenContractsURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"NFTRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"NFTRequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaidTokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"}],\"name\":\"TokenContractParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.RequestBuyParams\",\"name\":\"buyParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"enumIMarketplace.NFTRequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.NFTRequestInfo\",\"name\":\"nftRequestInfo\",\"type\":\"tuple\"}],\"name\":\"TokenSuccessfullyExchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedTokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paidTokensAmount\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumIMarketplace.PaymentType\",\"name\":\"paymentType\",\"type\":\"uint8\"}],\"name\":\"TokenSuccessfullyPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"__Marketplace_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams_\",\"type\":\"tuple\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenProxy_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenContractsURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.RequestBuyParams\",\"name\":\"requestBuyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIMarketplace.PaymentDetails\",\"name\":\"paymentDetails\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"futureTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BuyParams\",\"name\":\"buyParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"sig_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.Sig\",\"name\":\"permitSig_\",\"type\":\"tuple\"}],\"name\":\"buyTokenWithVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"cancelNFTRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"}],\"name\":\"createNFTRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContract_\",\"type\":\"address[]\"}],\"name\":\"getBaseTokenParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenParams[]\",\"name\":\"baseTokenParams_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getBaseTokenParamsPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.BaseTokenParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokenContracts_\",\"type\":\"address[]\"}],\"name\":\"getDetailedTokenParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.DetailedTokenParams[]\",\"name\":\"detailedTokenParams_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getDetailedTokenParamsPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"tokenParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"internalType\":\"structIMarketplace.DetailedTokenParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTRequestsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getNFTRequestsPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"enumIMarketplace.NFTRequestStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.NFTRequestInfo[]\",\"name\":\"nftRequests_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getTokenContractsPart\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddr_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getUserTokensPart\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIMarketplace.UserTokens[]\",\"name\":\"userTokens_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseTokenContractsURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenContractsURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractsRegistry_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pricePerOneToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minNFTFloorPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voucherTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voucherTokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNFTBuyable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isVoucherBuyable\",\"type\":\"bool\"}],\"internalType\":\"structIMarketplace.TokenParams\",\"name\":\"newTokenParams_\",\"type\":\"tuple\"}],\"name\":\"updateAllParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"desiredAmount_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"withdrawAll_\",\"type\":\"bool\"}],\"name\":\"withdrawCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetBaseTokenParams is a free data retrieval call binding the contract method 0x20acff14.
//
// Solidity: function getBaseTokenParams(address[] tokenContract_) view returns((address,bool,uint256,string)[] baseTokenParams_)
func (_Marketplace *MarketplaceCaller) GetBaseTokenParams(opts *bind.CallOpts, tokenContract_ []common.Address) ([]IMarketplaceBaseTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBaseTokenParams", tokenContract_)

	if err != nil {
		return *new([]IMarketplaceBaseTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBaseTokenParams)).(*[]IMarketplaceBaseTokenParams)

	return out0, err

}

// GetBaseTokenParams is a free data retrieval call binding the contract method 0x20acff14.
//
// Solidity: function getBaseTokenParams(address[] tokenContract_) view returns((address,bool,uint256,string)[] baseTokenParams_)
func (_Marketplace *MarketplaceSession) GetBaseTokenParams(tokenContract_ []common.Address) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParams(&_Marketplace.CallOpts, tokenContract_)
}

// GetBaseTokenParams is a free data retrieval call binding the contract method 0x20acff14.
//
// Solidity: function getBaseTokenParams(address[] tokenContract_) view returns((address,bool,uint256,string)[] baseTokenParams_)
func (_Marketplace *MarketplaceCallerSession) GetBaseTokenParams(tokenContract_ []common.Address) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParams(&_Marketplace.CallOpts, tokenContract_)
}

// GetBaseTokenParamsPart is a free data retrieval call binding the contract method 0x27a99cf7.
//
// Solidity: function getBaseTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,bool,uint256,string)[])
func (_Marketplace *MarketplaceCaller) GetBaseTokenParamsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBaseTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBaseTokenParamsPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceBaseTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceBaseTokenParams)).(*[]IMarketplaceBaseTokenParams)

	return out0, err

}

// GetBaseTokenParamsPart is a free data retrieval call binding the contract method 0x27a99cf7.
//
// Solidity: function getBaseTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,bool,uint256,string)[])
func (_Marketplace *MarketplaceSession) GetBaseTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetBaseTokenParamsPart is a free data retrieval call binding the contract method 0x27a99cf7.
//
// Solidity: function getBaseTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,bool,uint256,string)[])
func (_Marketplace *MarketplaceCallerSession) GetBaseTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceBaseTokenParams, error) {
	return _Marketplace.Contract.GetBaseTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetDetailedTokenParams is a free data retrieval call binding the contract method 0xcc1270cb.
//
// Solidity: function getDetailedTokenParams(address[] tokenContracts_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool,bool),string,string)[] detailedTokenParams_)
func (_Marketplace *MarketplaceCaller) GetDetailedTokenParams(opts *bind.CallOpts, tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDetailedTokenParams", tokenContracts_)

	if err != nil {
		return *new([]IMarketplaceDetailedTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceDetailedTokenParams)).(*[]IMarketplaceDetailedTokenParams)

	return out0, err

}

// GetDetailedTokenParams is a free data retrieval call binding the contract method 0xcc1270cb.
//
// Solidity: function getDetailedTokenParams(address[] tokenContracts_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool,bool),string,string)[] detailedTokenParams_)
func (_Marketplace *MarketplaceSession) GetDetailedTokenParams(tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParams(&_Marketplace.CallOpts, tokenContracts_)
}

// GetDetailedTokenParams is a free data retrieval call binding the contract method 0xcc1270cb.
//
// Solidity: function getDetailedTokenParams(address[] tokenContracts_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool,bool),string,string)[] detailedTokenParams_)
func (_Marketplace *MarketplaceCallerSession) GetDetailedTokenParams(tokenContracts_ []common.Address) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParams(&_Marketplace.CallOpts, tokenContracts_)
}

// GetDetailedTokenParamsPart is a free data retrieval call binding the contract method 0xce60e378.
//
// Solidity: function getDetailedTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool,bool),string,string)[])
func (_Marketplace *MarketplaceCaller) GetDetailedTokenParamsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenParams, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getDetailedTokenParamsPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceDetailedTokenParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceDetailedTokenParams)).(*[]IMarketplaceDetailedTokenParams)

	return out0, err

}

// GetDetailedTokenParamsPart is a free data retrieval call binding the contract method 0xce60e378.
//
// Solidity: function getDetailedTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool,bool),string,string)[])
func (_Marketplace *MarketplaceSession) GetDetailedTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetDetailedTokenParamsPart is a free data retrieval call binding the contract method 0xce60e378.
//
// Solidity: function getDetailedTokenParamsPart(uint256 offset_, uint256 limit_) view returns((address,(uint256,uint256,uint256,address,address,bool,bool,bool),string,string)[])
func (_Marketplace *MarketplaceCallerSession) GetDetailedTokenParamsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceDetailedTokenParams, error) {
	return _Marketplace.Contract.GetDetailedTokenParamsPart(&_Marketplace.CallOpts, offset_, limit_)
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

// GetNFTRequestsCount is a free data retrieval call binding the contract method 0xaadc9e19.
//
// Solidity: function getNFTRequestsCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetNFTRequestsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getNFTRequestsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNFTRequestsCount is a free data retrieval call binding the contract method 0xaadc9e19.
//
// Solidity: function getNFTRequestsCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetNFTRequestsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetNFTRequestsCount(&_Marketplace.CallOpts)
}

// GetNFTRequestsCount is a free data retrieval call binding the contract method 0xaadc9e19.
//
// Solidity: function getNFTRequestsCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetNFTRequestsCount() (*big.Int, error) {
	return _Marketplace.Contract.GetNFTRequestsCount(&_Marketplace.CallOpts)
}

// GetNFTRequestsPart is a free data retrieval call binding the contract method 0x58fc3a53.
//
// Solidity: function getNFTRequestsPart(uint256 offset_, uint256 limit_) view returns((address,address,uint256,address,uint8)[] nftRequests_)
func (_Marketplace *MarketplaceCaller) GetNFTRequestsPart(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceNFTRequestInfo, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getNFTRequestsPart", offset_, limit_)

	if err != nil {
		return *new([]IMarketplaceNFTRequestInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMarketplaceNFTRequestInfo)).(*[]IMarketplaceNFTRequestInfo)

	return out0, err

}

// GetNFTRequestsPart is a free data retrieval call binding the contract method 0x58fc3a53.
//
// Solidity: function getNFTRequestsPart(uint256 offset_, uint256 limit_) view returns((address,address,uint256,address,uint8)[] nftRequests_)
func (_Marketplace *MarketplaceSession) GetNFTRequestsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceNFTRequestInfo, error) {
	return _Marketplace.Contract.GetNFTRequestsPart(&_Marketplace.CallOpts, offset_, limit_)
}

// GetNFTRequestsPart is a free data retrieval call binding the contract method 0x58fc3a53.
//
// Solidity: function getNFTRequestsPart(uint256 offset_, uint256 limit_) view returns((address,address,uint256,address,uint8)[] nftRequests_)
func (_Marketplace *MarketplaceCallerSession) GetNFTRequestsPart(offset_ *big.Int, limit_ *big.Int) ([]IMarketplaceNFTRequestInfo, error) {
	return _Marketplace.Contract.GetNFTRequestsPart(&_Marketplace.CallOpts, offset_, limit_)
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

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0x842ffe75.
//
// Solidity: function buyTokenWithERC20(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithERC20(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithERC20", buyParams_, sig_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0x842ffe75.
//
// Solidity: function buyTokenWithERC20(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithERC20(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithERC20(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithERC20 is a paid mutator transaction binding the contract method 0x842ffe75.
//
// Solidity: function buyTokenWithERC20(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithERC20(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithERC20(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0x6b7b63a0.
//
// Solidity: function buyTokenWithETH(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithETH(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithETH", buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0x6b7b63a0.
//
// Solidity: function buyTokenWithETH(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithETH(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithETH(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithETH is a paid mutator transaction binding the contract method 0x6b7b63a0.
//
// Solidity: function buyTokenWithETH(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) payable returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithETH(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithETH(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0x9ee0b982.
//
// Solidity: function buyTokenWithNFT(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithNFT(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithNFT", buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0x9ee0b982.
//
// Solidity: function buyTokenWithNFT(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithNFT(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithNFT(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithNFT is a paid mutator transaction binding the contract method 0x9ee0b982.
//
// Solidity: function buyTokenWithNFT(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithNFT(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithNFT(&_Marketplace.TransactOpts, buyParams_, sig_)
}

// BuyTokenWithRequest is a paid mutator transaction binding the contract method 0xb9a5998d.
//
// Solidity: function buyTokenWithRequest((uint256,uint256,string) requestBuyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithRequest(opts *bind.TransactOpts, requestBuyParams_ IMarketplaceRequestBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithRequest", requestBuyParams_, sig_)
}

// BuyTokenWithRequest is a paid mutator transaction binding the contract method 0xb9a5998d.
//
// Solidity: function buyTokenWithRequest((uint256,uint256,string) requestBuyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithRequest(requestBuyParams_ IMarketplaceRequestBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithRequest(&_Marketplace.TransactOpts, requestBuyParams_, sig_)
}

// BuyTokenWithRequest is a paid mutator transaction binding the contract method 0xb9a5998d.
//
// Solidity: function buyTokenWithRequest((uint256,uint256,string) requestBuyParams_, (uint256,bytes32,bytes32,uint8) sig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithRequest(requestBuyParams_ IMarketplaceRequestBuyParams, sig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithRequest(&_Marketplace.TransactOpts, requestBuyParams_, sig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x65a896f3.
//
// Solidity: function buyTokenWithVoucher(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_, (uint256,bytes32,bytes32,uint8) permitSig_) returns()
func (_Marketplace *MarketplaceTransactor) BuyTokenWithVoucher(opts *bind.TransactOpts, buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig, permitSig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buyTokenWithVoucher", buyParams_, sig_, permitSig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x65a896f3.
//
// Solidity: function buyTokenWithVoucher(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_, (uint256,bytes32,bytes32,uint8) permitSig_) returns()
func (_Marketplace *MarketplaceSession) BuyTokenWithVoucher(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig, permitSig_ IMarketplaceSig) (*types.Transaction, error) {
	return _Marketplace.Contract.BuyTokenWithVoucher(&_Marketplace.TransactOpts, buyParams_, sig_, permitSig_)
}

// BuyTokenWithVoucher is a paid mutator transaction binding the contract method 0x65a896f3.
//
// Solidity: function buyTokenWithVoucher(((address,uint256,uint256,uint256),address,address,uint256,string) buyParams_, (uint256,bytes32,bytes32,uint8) sig_, (uint256,bytes32,bytes32,uint8) permitSig_) returns()
func (_Marketplace *MarketplaceTransactorSession) BuyTokenWithVoucher(buyParams_ IMarketplaceBuyParams, sig_ IMarketplaceSig, permitSig_ IMarketplaceSig) (*types.Transaction, error) {
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

// CreateNFTRequest is a paid mutator transaction binding the contract method 0x6d04b6fa.
//
// Solidity: function createNFTRequest(address nftContract_, uint256 nftId_, address tokenContract_) returns(uint256 requestId_)
func (_Marketplace *MarketplaceTransactor) CreateNFTRequest(opts *bind.TransactOpts, nftContract_ common.Address, nftId_ *big.Int, tokenContract_ common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createNFTRequest", nftContract_, nftId_, tokenContract_)
}

// CreateNFTRequest is a paid mutator transaction binding the contract method 0x6d04b6fa.
//
// Solidity: function createNFTRequest(address nftContract_, uint256 nftId_, address tokenContract_) returns(uint256 requestId_)
func (_Marketplace *MarketplaceSession) CreateNFTRequest(nftContract_ common.Address, nftId_ *big.Int, tokenContract_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateNFTRequest(&_Marketplace.TransactOpts, nftContract_, nftId_, tokenContract_)
}

// CreateNFTRequest is a paid mutator transaction binding the contract method 0x6d04b6fa.
//
// Solidity: function createNFTRequest(address nftContract_, uint256 nftId_, address tokenContract_) returns(uint256 requestId_)
func (_Marketplace *MarketplaceTransactorSession) CreateNFTRequest(nftContract_ common.Address, nftId_ *big.Int, tokenContract_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateNFTRequest(&_Marketplace.TransactOpts, nftContract_, nftId_, tokenContract_)
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

// UpdateAllParams is a paid mutator transaction binding the contract method 0x7fd471b5.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactor) UpdateAllParams(opts *bind.TransactOpts, tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateAllParams", tokenContract_, name_, symbol_, newTokenParams_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x7fd471b5.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceSession) UpdateAllParams(tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateAllParams(&_Marketplace.TransactOpts, tokenContract_, name_, symbol_, newTokenParams_)
}

// UpdateAllParams is a paid mutator transaction binding the contract method 0x7fd471b5.
//
// Solidity: function updateAllParams(address tokenContract_, string name_, string symbol_, (uint256,uint256,uint256,address,address,bool,bool,bool) newTokenParams_) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateAllParams(tokenContract_ common.Address, name_ string, symbol_ string, newTokenParams_ IMarketplaceTokenParams) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateAllParams(&_Marketplace.TransactOpts, tokenContract_, name_, symbol_, newTokenParams_)
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
	NftContract   common.Address
	NftId         *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNFTRequestCreated is a free log retrieval operation binding the contract event 0x01ac2115ffca9df821a874e019c7cd01df4a3e49c7f72ead1bf66efef3a7c547.
//
// Solidity: event NFTRequestCreated(uint256 indexed requestId, address indexed requester, address nftContract, uint256 nftId, address indexed tokenContract)
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

// WatchNFTRequestCreated is a free log subscription operation binding the contract event 0x01ac2115ffca9df821a874e019c7cd01df4a3e49c7f72ead1bf66efef3a7c547.
//
// Solidity: event NFTRequestCreated(uint256 indexed requestId, address indexed requester, address nftContract, uint256 nftId, address indexed tokenContract)
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

// ParseNFTRequestCreated is a log parse operation binding the contract event 0x01ac2115ffca9df821a874e019c7cd01df4a3e49c7f72ead1bf66efef3a7c547.
//
// Solidity: event NFTRequestCreated(uint256 indexed requestId, address indexed requester, address nftContract, uint256 nftId, address indexed tokenContract)
func (_Marketplace *MarketplaceFilterer) ParseNFTRequestCreated(log types.Log) (*MarketplaceNFTRequestCreated, error) {
	event := new(MarketplaceNFTRequestCreated)
	if err := _Marketplace.contract.UnpackLog(event, "NFTRequestCreated", log); err != nil {
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

// MarketplaceTokenContractParamsUpdatedIterator is returned from FilterTokenContractParamsUpdated and is used to iterate over the raw logs and unpacked data for TokenContractParamsUpdated events raised by the Marketplace contract.
type MarketplaceTokenContractParamsUpdatedIterator struct {
	Event *MarketplaceTokenContractParamsUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceTokenContractParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceTokenContractParamsUpdated)
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
		it.Event = new(MarketplaceTokenContractParamsUpdated)
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
func (it *MarketplaceTokenContractParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceTokenContractParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceTokenContractParamsUpdated represents a TokenContractParamsUpdated event raised by the Marketplace contract.
type MarketplaceTokenContractParamsUpdated struct {
	TokenContract common.Address
	TokenName     string
	TokenSymbol   string
	TokenParams   IMarketplaceTokenParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenContractParamsUpdated is a free log retrieval operation binding the contract event 0x08c050efe01a2e294a466b4d07435c9a339e5e8e7ee7dc3e1b61e9182cae597d.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) FilterTokenContractParamsUpdated(opts *bind.FilterOpts, tokenContract []common.Address) (*MarketplaceTokenContractParamsUpdatedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenContractParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenContractParamsUpdatedIterator{contract: _Marketplace.contract, event: "TokenContractParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenContractParamsUpdated is a free log subscription operation binding the contract event 0x08c050efe01a2e294a466b4d07435c9a339e5e8e7ee7dc3e1b61e9182cae597d.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) WatchTokenContractParamsUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenContractParamsUpdated, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenContractParamsUpdated", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceTokenContractParamsUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
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

// ParseTokenContractParamsUpdated is a log parse operation binding the contract event 0x08c050efe01a2e294a466b4d07435c9a339e5e8e7ee7dc3e1b61e9182cae597d.
//
// Solidity: event TokenContractParamsUpdated(address indexed tokenContract, string tokenName, string tokenSymbol, (uint256,uint256,uint256,address,address,bool,bool,bool) tokenParams)
func (_Marketplace *MarketplaceFilterer) ParseTokenContractParamsUpdated(log types.Log) (*MarketplaceTokenContractParamsUpdated, error) {
	event := new(MarketplaceTokenContractParamsUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "TokenContractParamsUpdated", log); err != nil {
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
	Recipient      common.Address
	BuyParams      IMarketplaceRequestBuyParams
	NftRequestInfo IMarketplaceNFTRequestInfo
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenSuccessfullyExchanged is a free log retrieval operation binding the contract event 0xbd82d66d3b336458814e5ce7b8c885c577e4ff7c5b8bd58fdc76a31a1ba7f01b.
//
// Solidity: event TokenSuccessfullyExchanged(address indexed recipient, (uint256,uint256,string) buyParams, (address,address,uint256,address,uint8) nftRequestInfo)
func (_Marketplace *MarketplaceFilterer) FilterTokenSuccessfullyExchanged(opts *bind.FilterOpts, recipient []common.Address) (*MarketplaceTokenSuccessfullyExchangedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenSuccessfullyExchanged", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenSuccessfullyExchangedIterator{contract: _Marketplace.contract, event: "TokenSuccessfullyExchanged", logs: logs, sub: sub}, nil
}

// WatchTokenSuccessfullyExchanged is a free log subscription operation binding the contract event 0xbd82d66d3b336458814e5ce7b8c885c577e4ff7c5b8bd58fdc76a31a1ba7f01b.
//
// Solidity: event TokenSuccessfullyExchanged(address indexed recipient, (uint256,uint256,string) buyParams, (address,address,uint256,address,uint8) nftRequestInfo)
func (_Marketplace *MarketplaceFilterer) WatchTokenSuccessfullyExchanged(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenSuccessfullyExchanged, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenSuccessfullyExchanged", recipientRule)
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

// ParseTokenSuccessfullyExchanged is a log parse operation binding the contract event 0xbd82d66d3b336458814e5ce7b8c885c577e4ff7c5b8bd58fdc76a31a1ba7f01b.
//
// Solidity: event TokenSuccessfullyExchanged(address indexed recipient, (uint256,uint256,string) buyParams, (address,address,uint256,address,uint8) nftRequestInfo)
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
	Recipient        common.Address
	MintedTokenPrice *big.Int
	PaidTokensAmount *big.Int
	BuyParams        IMarketplaceBuyParams
	PaymentType      uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenSuccessfullyPurchased is a free log retrieval operation binding the contract event 0x36f4838ae4627a42a86813d34d41a61b3a30dd4bfe08014117f3d32392471c6f.
//
// Solidity: event TokenSuccessfullyPurchased(address indexed recipient, uint256 mintedTokenPrice, uint256 paidTokensAmount, ((address,uint256,uint256,uint256),address,address,uint256,string) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) FilterTokenSuccessfullyPurchased(opts *bind.FilterOpts, recipient []common.Address) (*MarketplaceTokenSuccessfullyPurchasedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "TokenSuccessfullyPurchased", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTokenSuccessfullyPurchasedIterator{contract: _Marketplace.contract, event: "TokenSuccessfullyPurchased", logs: logs, sub: sub}, nil
}

// WatchTokenSuccessfullyPurchased is a free log subscription operation binding the contract event 0x36f4838ae4627a42a86813d34d41a61b3a30dd4bfe08014117f3d32392471c6f.
//
// Solidity: event TokenSuccessfullyPurchased(address indexed recipient, uint256 mintedTokenPrice, uint256 paidTokensAmount, ((address,uint256,uint256,uint256),address,address,uint256,string) buyParams, uint8 paymentType)
func (_Marketplace *MarketplaceFilterer) WatchTokenSuccessfullyPurchased(opts *bind.WatchOpts, sink chan<- *MarketplaceTokenSuccessfullyPurchased, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "TokenSuccessfullyPurchased", recipientRule)
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

// ParseTokenSuccessfullyPurchased is a log parse operation binding the contract event 0x36f4838ae4627a42a86813d34d41a61b3a30dd4bfe08014117f3d32392471c6f.
//
// Solidity: event TokenSuccessfullyPurchased(address indexed recipient, uint256 mintedTokenPrice, uint256 paidTokensAmount, ((address,uint256,uint256,uint256),address,address,uint256,string) buyParams, uint8 paymentType)
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
