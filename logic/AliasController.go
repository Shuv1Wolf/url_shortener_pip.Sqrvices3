package logic

import (
	"context"
	data1 "pip_toolkit_url_shortener/data/version1"
	persist "pip_toolkit_url_shortener/persistence"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
)

type AliasController struct {
	persistence persist.IAliasPersistance
	commandSet  *AliasCommandSet
}

func NewAliasController() *AliasController {
	c := &AliasController{}
	return c
}

func (c *AliasController) SetReferences(ctx context.Context, references cref.IReferences) {
	locator := cref.NewDescriptor("alias", "persistence", "*", "*", "1.0")
	p, err := references.GetOneRequired(locator)
	if p != nil && err == nil {
		if _pers, ok := p.(persist.IAliasPersistance); ok {
			c.persistence = _pers
			return
		}
	}
	panic(cref.NewReferenceError("alias.controller.SetReferences", locator))
}

func (c *AliasController) GetCommandSet() *ccmd.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewAliasCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

func (c *AliasController) CreateAlias(ctx context.Context, correlationId string,
	alais data1.AliasV1) (data1.AliasV1, error) {
	if alais.Id == "" {
		alais.Id = cdata.IdGenerator.NextLong()
	}

	return c.persistence.Create(ctx, correlationId, alais)
}

func (c *AliasController) DeleteByAlias(ctx context.Context, correlationId string,
	alais string) (string, error) {

	return c.persistence.DeleteByKey(ctx, correlationId, alais)
}

func (c *AliasController) GetOneByAlias(ctx context.Context, correlationId string,
	alais string) (data1.AliasV1, error) {

	return c.persistence.GetOneByKey(ctx, correlationId, alais)
}
