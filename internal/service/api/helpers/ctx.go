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
	dbCtxKey
	minterCtxKey
	coingeckoCtxKey
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

func DB(r *http.Request) data.DB {
	return r.Context().Value(dbCtxKey).(data.DB)
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
