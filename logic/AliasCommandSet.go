package logic

import (
	"context"
	data1 "pip_toolkit_url_shortener/data/version1"

	ccmd "github.com/pip-services3-gox/pip-services3-commons-gox/commands"
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
)

type AliasCommandSet struct {
	ccmd.CommandSet
	controller     IAliasController
	aliasConvertor cconv.IJSONEngine[data1.AliasV1]
}

func NewAliasCommandSet(controller IAliasController) *AliasCommandSet {
	c := &AliasCommandSet{
		CommandSet:     *ccmd.NewCommandSet(),
		controller:     controller,
		aliasConvertor: cconv.NewDefaultCustomTypeJsonConvertor[data1.AliasV1](),
	}

	c.AddCommand(c.makeGetCreateAliasCommand())
	c.AddCommand(c.makeGetOneByAliasCommand())
	c.AddCommand(c.makeDeleteByAliasCommand())

	return c
}

func (c *AliasCommandSet) makeGetCreateAliasCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"create_alias",
		cvalid.NewObjectSchema().
			WithRequiredProperty("alias", data1.NewAliasV1Schema()),
		func(ctx context.Context, correlationId string, args *run.Parameters) (any, error) {

			var alias data1.AliasV1
			if _alias, ok := args.GetAsObject("alias"); ok {
				buf, err := cconv.JsonConverter.ToJson(_alias)
				if err != nil {
					return nil, err
				}
				alias, err = c.aliasConvertor.FromJson(buf)
				if err != nil {
					return nil, err
				}
			}
			return c.controller.CreateAlias(ctx, correlationId, alias)
		})
}

func (c *AliasCommandSet) makeGetOneByAliasCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_one_by_alias",
		cvalid.NewObjectSchema().
			WithRequiredProperty("alias", cconv.String),
		func(ctx context.Context, correlationId string, args *run.Parameters) (any, error) {
			return c.controller.GetOneByAlias(ctx, correlationId, args.GetAsString("alias"))
		})
}

func (c *AliasCommandSet) makeDeleteByAliasCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"delete_by_alias",
		cvalid.NewObjectSchema().
			WithRequiredProperty("alias", cconv.String),
		func(ctx context.Context, correlationId string, args *run.Parameters) (any, error) {
			return c.controller.DeleteByAlias(ctx, correlationId, args.GetAsString("alias"))
		})
}
