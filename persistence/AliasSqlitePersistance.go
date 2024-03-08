package persistence

import (
	"context"
	data1 "pip_toolkit_url_shortener/data/version1"

	csql "github.com/pip-services3-gox/pip-services3-sqlite-gox/persistence"
)

type AliasSqlitePersistence struct {
	*csql.IdentifiableSqlitePersistence[data1.AliasV1, string]
}

func NewAliasSqlPersistance() *AliasSqlitePersistence {
	c := &AliasSqlitePersistence{}
	c.IdentifiableSqlitePersistence = csql.InheritIdentifiableSqlitePersistence[data1.AliasV1, string](c, "alias")
	return c
}

func (c *AliasSqlitePersistence) DefineSchema() {
	c.ClearSchema()
	c.IdentifiableSqlitePersistence.DefineSchema()
	c.EnsureSchema("CREATE TABLE " + c.QuotedTableName() + " (\"id\" TEXT PRIMARY KEY, \"alias\" TEXT, \"url\" TEXT)")
	c.EnsureIndex(c.TableName+"_key", map[string]string{"alias": "1"}, map[string]string{"unique": "true"})
}

func (c *AliasSqlitePersistence) GetOneByKey(ctx context.Context, correlationId string, key string) (item data1.AliasV1, err error) {
	query := "SELECT * FROM " + c.QuotedTableName() + " WHERE \"alias\"=$1"

	qResult, err := c.Client.QueryContext(ctx, query, key)
	if err != nil {
		return item, err
	}
	defer qResult.Close()

	if !qResult.Next() {
		return item, qResult.Err()
	}

	result, err := c.Overrides.ConvertToPublic(qResult)

	if err == nil {
		c.Logger.Trace(ctx, correlationId, "Retrieved from %s with key = %s", c.TableName, key)
		return result, err
	}
	c.Logger.Trace(ctx, correlationId, "Nothing found from %s with key = %s", c.TableName, key)
	return item, err
}

func (c *AliasSqlitePersistence) DeleteByKey(ctx context.Context, correlationId string, key string) (id string, err error) {
	alias, err := c.GetOneByKey(ctx, correlationId, key)
	if err != nil {
		return "", err
	}

	result, err := c.DeleteById(ctx, correlationId, alias.Id)
	if err != nil {
		c.Logger.Trace(ctx, correlationId, "Nothing found from %s with key = %s", c.TableName, key)
		return "", err
	}

	c.Logger.Trace(ctx, correlationId, "Removed from %s with key = %s", c.TableName, key)
	return result.Id, nil
}
