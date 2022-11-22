package helpers

import (
	"context"
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data/external"

	pricer "gitlab.com/tokend/nft-books/price-svc/connector"

	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	networkerConnector "gitlab.com/tokend/nft-books/network-svc/connector/api"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	booksQCtxKey
	paymentsQCtxKey
	generatorDBCtxKey
	minterCtxKey
	pricerCtxKey
	networkerConnectorCtxKey
	apiRestrictionsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxBooksQ(q external.BookQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, booksQCtxKey, q)
	}
}

func CtxPaymentsQ(q external.PaymentsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, paymentsQCtxKey, q)
	}
}

func CtxGeneratorDB(db data.GeneratorDB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, generatorDBCtxKey, db)
	}
}

func CtxMinter(entry config.EthMinterConfig) func(context.Context) context.Context {
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

func Minter(r *http.Request) config.EthMinterConfig {
	return r.Context().Value(minterCtxKey).(config.EthMinterConfig)
}

func BooksQ(r *http.Request) external.BookQ {
	return r.Context().Value(booksQCtxKey).(external.BookQ).New()
}

func PaymentsQ(r *http.Request) external.PaymentsQ {
	return r.Context().Value(paymentsQCtxKey).(external.PaymentsQ).New()
}

func GeneratorDB(r *http.Request) data.GeneratorDB {
	return r.Context().Value(generatorDBCtxKey).(data.GeneratorDB).New()
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func ApiRestrictions(r *http.Request) config.ApiRestrictions {
	return r.Context().Value(apiRestrictionsCtxKey).(config.ApiRestrictions)
}

func CtxNetworkerConnector(entry networkerConnector.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkerConnectorCtxKey, entry)
	}
}

func NetworkerConnector(r *http.Request) networkerConnector.Connector {
	return r.Context().Value(networkerConnectorCtxKey).(networkerConnector.Connector)
}
