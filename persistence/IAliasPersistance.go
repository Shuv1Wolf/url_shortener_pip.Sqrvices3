package persistence

import (
	"context"
	data1 "pip_toolkit_url_shortener/data/version1"
)

type IAliasPersistance interface {
	Create(ctx context.Context, correlationId string, item data1.AliasV1) (data1.AliasV1, error)

	GetOneByKey(ctx context.Context, correlationId string, alias string) (item data1.AliasV1, err error)

	DeleteByKey(ctx context.Context, correlationId string, alias string) (id string, err error)
}
