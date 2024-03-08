package containers

import (
	factory "pip_toolkit_url_shortener/build"

	cproc "github.com/pip-services3-gox/pip-services3-container-gox/container"
	rbuild "github.com/pip-services3-gox/pip-services3-rpc-gox/build"
	sqlite "github.com/pip-services3-gox/pip-services3-sqlite-gox/build"
	cswagger "github.com/pip-services3-gox/pip-services3-swagger-gox/build"
)

type AliasProcess struct {
	cproc.ProcessContainer
}

func NewAliasProcess() *AliasProcess {
	c := &AliasProcess{
		ProcessContainer: *cproc.NewProcessContainer("alias", "Alias microservice"),
	}

	c.AddFactory(factory.NewAliasServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	c.AddFactory(sqlite.NewDefaultSqliteFactory())
	c.AddFactory(cswagger.NewDefaultSwaggerFactory())

	return c
}
