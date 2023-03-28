package helpers

import (
	"context"
	"github.com/dl-nft-books/core-svc/internal/config"
	"github.com/dl-nft-books/core-svc/internal/data"
	"net/http"

	s3connector "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector"
	"gitlab.com/tokend/nft-books/doorman/connector"

	pricer "gitlab.com/tokend/nft-books/price-svc/connector"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	// Base configs
	logCtxKey ctxKey = iota
	dbCtxKey

	// Custom configs
	minterCtxKey
	apiRestrictionsCtxKey
	promocoderCtxKey
	ipfserCtxKey
	mimeTypesCtxKey

	// Connectors
	pricerCtxKey
	bookerCtxKey
	trackerCtxKey
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

func CtxDocumenterConnector(entry s3connector.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, documenterConnectorCtxKey, entry)
	}
}

func DocumenterConnector(r *http.Request) s3connector.Connector {
	return r.Context().Value(documenterConnectorCtxKey).(s3connector.Connector)
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

func Tracker(r *http.Request) *tracker.Connector {
	return r.Context().Value(trackerCtxKey).(*tracker.Connector)
}

func CtxTracker(entry *tracker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, trackerCtxKey, entry)
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
