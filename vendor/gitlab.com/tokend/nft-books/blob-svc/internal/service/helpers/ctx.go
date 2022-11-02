package helpers

import (
	"context"

	"gitlab.com/tokend/nft-books/doorman/connector"

	"net/http"

	"gitlab.com/tokend/nft-books/blob-svc/internal/config"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	mimeTypesCtxKey
	awsCfgKey
	doormanConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxMimeTypes(entry *config.MimeTypes) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mimeTypesCtxKey, entry)
	}
}

func CtxAwsConfig(entry *config.AWSConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, awsCfgKey, entry)
	}
}

func AwsConfig(r *http.Request) *config.AWSConfig {
	return r.Context().Value(awsCfgKey).(*config.AWSConfig)
}

func MimeTypes(r *http.Request) *config.MimeTypes {
	return r.Context().Value(mimeTypesCtxKey).(*config.MimeTypes)
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}

func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}
