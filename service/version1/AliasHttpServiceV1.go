package service1

import (
	"context"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cservices "github.com/pip-services3-gox/pip-services3-rpc-gox/services"
)

type AliasHttpServiceV1 struct {
	cservices.CommandableHttpService
}

func NewAliasHttpServiceV1() *AliasHttpServiceV1 {
	c := &AliasHttpServiceV1{}
	c.CommandableHttpService = *cservices.InheritCommandableHttpService(c, "v1/alias")
	c.DependencyResolver.Put(context.Background(), "controller", cref.NewDescriptor("alias", "controller", "*", "*", "1.0"))
	return c
}
