package helpers

import (
	"context"
	"github.com/dl-nft-books/core-svc/internal/config"
	"github.com/dl-nft-books/core-svc/internal/data"
	networker "github.com/dl-nft-books/network-svc/connector"
	"net/http"

	s3connector "github.com/dl-nft-books/blob-svc/connector/api"
	booker "github.com/dl-nft-books/book-svc/connector"
	"github.com/dl-nft-books/doorman/connector"

	pricer "github.com/dl-nft-books/price-svc/connector"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	// Base configs
	logCtxKey ctxKey = iota
	dbCtxKey

	// Custom configs
	minterCtxKey
	transacterCtxKey
	apiRestrictionsCtxKey
	promocoderCtxKey
	ipfserCtxKey
	mimeTypesCtxKey

	// Connectors
	pricerCtxKey
	bookerCtxKey
	networkerCtxKey
	doormanConnectorCtxKey
	documenterConnectorCtxKey
)

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func DB(r *http.Request) data.DB {
	return r.Context().Value(dbCtxKey).(data.DB).New()
}

func CtxDB(db data.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, db)
	}
}

func Minter(r *http.Request) config.MintConfig {
	return r.Context().Value(minterCtxKey).(config.MintConfig)
}

func CtxMinter(entry config.MintConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, minterCtxKey, entry)
	}
}
func Transacter(r *http.Request) config.TransactionConfig {
	return r.Context().Value(transacterCtxKey).(config.TransactionConfig)
}

func CtxTransacter(entry config.TransactionConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, transacterCtxKey, entry)
	}
}

func Ipfser(r *http.Request) config.IpfserCfg {
	return r.Context().Value(ipfserCtxKey).(config.IpfserCfg)
}

func CtxIpfser(entry config.IpfserCfg) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ipfserCtxKey, entry)
	}
}

func Promocoder(r *http.Request) config.PromocoderCfg {
	return r.Context().Value(promocoderCtxKey).(config.PromocoderCfg)
}

func CtxPromocoder(promocoder config.PromocoderCfg) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, promocoderCtxKey, promocoder)
	}
}

func ApiRestrictions(r *http.Request) config.ApiRestrictions {
	return r.Context().Value(apiRestrictionsCtxKey).(config.ApiRestrictions)
}

func CtxApiRestrictions(restrictions config.ApiRestrictions) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, apiRestrictionsCtxKey, restrictions)
	}
}

func Pricer(r *http.Request) *pricer.Connector {
	return r.Context().Value(pricerCtxKey).(*pricer.Connector)
}

func CtxPricer(entry *pricer.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pricerCtxKey, entry)
	}
}

func CtxDocumenterConnector(entry *s3connector.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, documenterConnectorCtxKey, entry)
	}
}

func DocumenterConnector(r *http.Request) *s3connector.Connector {
	return r.Context().Value(documenterConnectorCtxKey).(*s3connector.Connector)
}

func MimeTypes(r *http.Request) *config.MimeTypes {
	return r.Context().Value(mimeTypesCtxKey).(*config.MimeTypes)
}

func Booker(r *http.Request) *booker.Connector {
	return r.Context().Value(bookerCtxKey).(*booker.Connector)
}

func CtxBooker(entry *booker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, bookerCtxKey, entry)
	}
}

func Networker(r *http.Request) *networker.Connector {
	return r.Context().Value(networkerCtxKey).(*networker.Connector)
}

func CtxNetworker(entry *networker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkerCtxKey, entry)
	}
}

func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}
