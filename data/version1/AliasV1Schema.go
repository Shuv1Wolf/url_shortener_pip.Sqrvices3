package data1

import (
	cconv "github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	cvalid "github.com/pip-services3-gox/pip-services3-commons-gox/validate"
)

type AliasV1Schema struct {
	cvalid.ObjectSchema
}

func NewAliasV1Schema() *AliasV1Schema {
	c := AliasV1Schema{}
	c.ObjectSchema = *cvalid.NewObjectSchema()

	c.WithOptionalProperty("id", cconv.String)
	c.WithRequiredProperty("alias", cconv.String)
	c.WithRequiredProperty("url", cconv.String)
	return &c
}
