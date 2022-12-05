package helpers

import (
	"context"
	"net/http"

	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector"
	"gitlab.com/tokend/nft-books/doorman/connector"

	pricer "gitlab.com/tokend/nft-books/price-svc/connector"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
	minterCtxKey
	apiRestrictionsCtxKey

	pricerCtxKey
	bookerCtxKey
	trackerCtxKey
	doormanConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxDB(db data.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, db)
	}
}

func DB(r *http.Request) data.DB {
	return r.Context().Value(dbCtxKey).(data.DB).New()
}

func CtxMinter(entry config.MintConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, minterCtxKey, entry)
	}
}

func CtxPricer(entry *pricer.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pricerCtxKey, entry)
	}
}

func CtxApiRestrictions(restrictions config.ApiRestrictions) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, apiRestrictionsCtxKey, restrictions)
	}
}

func Pricer(r *http.Request) *pricer.Connector {
	return r.Context().Value(pricerCtxKey).(*pricer.Connector)
}

func Minter(r *http.Request) config.MintConfig {
	return r.Context().Value(minterCtxKey).(config.MintConfig)
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func ApiRestrictions(r *http.Request) config.ApiRestrictions {
	return r.Context().Value(apiRestrictionsCtxKey).(config.ApiRestrictions)
}

func CtxBooker(entry *booker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, bookerCtxKey, entry)
	}
}

func Booker(r *http.Request) *booker.Connector {
	return r.Context().Value(bookerCtxKey).(*booker.Connector)
}

func CtxTracker(entry *tracker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, trackerCtxKey, entry)
	}
}

func Tracker(r *http.Request) *tracker.Connector {
	return r.Context().Value(trackerCtxKey).(*tracker.Connector)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}

func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}
