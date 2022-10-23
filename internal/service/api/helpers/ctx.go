package helpers

import (
	"context"
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	booksQCtxKey
	generatorDBCtxKey
	minterCtxKey
	coingeckoCtxKey
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

func CtxCoingecko(entry config.CoingeckoConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, coingeckoCtxKey, entry)
	}
}

func Coingecko(r *http.Request) config.CoingeckoConfig {
	return r.Context().Value(coingeckoCtxKey).(config.CoingeckoConfig)
}

func Minter(r *http.Request) config.EthMinterConfig {
	return r.Context().Value(minterCtxKey).(config.EthMinterConfig)
}

func BooksQ(r *http.Request) data.BookQ {
	return r.Context().Value(booksQCtxKey).(data.BookQ)
}

func GeneratorDB(r *http.Request) data.GeneratorDB {
	return r.Context().Value(generatorDBCtxKey).(data.GeneratorDB)
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
