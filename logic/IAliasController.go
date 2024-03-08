package logic

import (
	"context"
	data1 "pip_toolkit_url_shortener/data/version1"
)

type IAliasController interface {
	CreateAlias(ctx context.Context, correlationId string, item data1.AliasV1) (data1.AliasV1, error)

	GetOneByAlias(ctx context.Context, correlationId string, alias string) (item data1.AliasV1, err error)

	DeleteByAlias(ctx context.Context, correlationId string, alias string) (id string, err error)
}
