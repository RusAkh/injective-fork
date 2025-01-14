package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyRelayerAddr            = sdkerrors.Register(ModuleName, 1, "relayer address is empty")
	ErrBadRatesCount               = sdkerrors.Register(ModuleName, 2, "bad rates count")
	ErrBadResolveTimesCount        = sdkerrors.Register(ModuleName, 3, "bad resolve times")
	ErrBadRequestIDsCount          = sdkerrors.Register(ModuleName, 4, "bad request ID")
	ErrRelayerNotAuthorized        = sdkerrors.Register(ModuleName, 5, "relayer not authorized")
	ErrBadPriceFeedBaseCount       = sdkerrors.Register(ModuleName, 6, "bad price feed base count")
	ErrBadPriceFeedQuoteCount      = sdkerrors.Register(ModuleName, 7, "bad price feed quote count")
	ErrUnsupportedOracleType       = sdkerrors.Register(ModuleName, 8, "unsupported oracle type")
	ErrBadMessagesCount            = sdkerrors.Register(ModuleName, 9, "bad messages count")
	ErrBadCoinbaseMessage          = sdkerrors.Register(ModuleName, 10, "bad Coinbase message")
	ErrInvalidEthereumSignature    = sdkerrors.Register(ModuleName, 11, "bad Ethereum signature")
	ErrBadCoinbaseMessageTimestamp = sdkerrors.Register(ModuleName, 12, "bad Coinbase message timestamp")
	ErrCoinbasePriceNotFound       = sdkerrors.Register(ModuleName, 13, "Coinbase price not found")
	ErrBadPrice                    = sdkerrors.Register(ModuleName, 14, "Prices must be positive")
	ErrPriceTooLarge               = sdkerrors.Register(ModuleName, 15, "Prices must be less than 10 million.")
	ErrInvalidBandIBCRequest       = sdkerrors.Register(ModuleName, 16, "Invalid Band IBC Request")
	ErrSample                      = sdkerrors.Register(ModuleName, 17, "sample error")
	ErrInvalidPacketTimeout        = sdkerrors.Register(ModuleName, 18, "invalid packet timeout")
	ErrBadSymbolsCount             = sdkerrors.Register(ModuleName, 19, "invalid symbols count")
	ErrBadIBCPortBind              = sdkerrors.Register(ModuleName, 20, "could not claim port capability")
	ErrInvalidPortID               = sdkerrors.Register(ModuleName, 21, "invalid IBC Port ID")
	ErrInvalidChannelID            = sdkerrors.Register(ModuleName, 22, "invalid IBC Channel ID")
	ErrBadRequestInterval          = sdkerrors.Register(ModuleName, 23, "invalid Band IBC request interval")
	ErrInvalidBandIBCUpdateRequest = sdkerrors.Register(ModuleName, 24, "Invalid Band IBC Update Request Proposal")
	ErrBandIBCRequestNotFound      = sdkerrors.Register(ModuleName, 25, "Band IBC Oracle Request not found")
	ErrEmptyBaseInfo               = sdkerrors.Register(ModuleName, 26, "Base Info is empty")
	ErrEmptyProvider               = sdkerrors.Register(ModuleName, 27, "provider is empty")
	ErrInvalidProvider             = sdkerrors.Register(ModuleName, 28, "invalid provider name")
	ErrInvalidSymbol               = sdkerrors.Register(ModuleName, 29, "invalid symbol")
	ErrRelayerAlreadyExists        = sdkerrors.Register(ModuleName, 30, "relayer already exists")
	ErrProviderPriceNotFound       = sdkerrors.Register(ModuleName, 31, "provider price not found")
)
