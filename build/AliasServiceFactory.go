package build

import (
	"pip_toolkit_url_shortener/logic"
	"pip_toolkit_url_shortener/persistence"
	service1 "pip_toolkit_url_shortener/service/version1"

	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type AliasServiceFactory struct {
	cbuild.Factory
}

func NewAliasServiceFactory() *AliasServiceFactory {
	c := &AliasServiceFactory{
		Factory: *cbuild.NewFactory(),
	}

	sqlitePersistanceDescriptor := cref.NewDescriptor("alias", "persistence", "sqlite", "*", "1.0")
	controllerDescriptor := cref.NewDescriptor("alias", "controller", "default", "*", "1.0")
	httpServiceV1Descriptor := cref.NewDescriptor("alias", "service", "http", "*", "1.0")

	c.RegisterType(sqlitePersistanceDescriptor, persistence.NewAliasSqlPersistance)
	c.RegisterType(controllerDescriptor, logic.NewAliasController)
	c.RegisterType(httpServiceV1Descriptor, service1.NewAliasHttpServiceV1)

	return c
}
