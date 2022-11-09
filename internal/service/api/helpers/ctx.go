package helpers

import (
	"context"
	"net/http"

	pricer "gitlab.com/tokend/nft-books/price-svc/connector"

	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"

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
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxBooksQ(q data.BookQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, booksQCtxKey, q)
	}
}

func CtxPaymentsQ(q data.PaymentsQ) func(context.Context) context.Context {
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

func Pricer(r *http.Request) *pricer.Connector {
	return r.Context().Value(pricerCtxKey).(*pricer.Connector)
}

func Minter(r *http.Request) config.EthMinterConfig {
	return r.Context().Value(minterCtxKey).(config.EthMinterConfig)
}

func BooksQ(r *http.Request) data.BookQ {
	return r.Context().Value(booksQCtxKey).(data.BookQ).New()
}

func PaymentsQ(r *http.Request) data.PaymentsQ {
	return r.Context().Value(paymentsQCtxKey).(data.PaymentsQ).New()
}

func GeneratorDB(r *http.Request) data.GeneratorDB {
	return r.Context().Value(generatorDBCtxKey).(data.GeneratorDB).New()
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
