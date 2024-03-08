package test_persistance

import (
	"context"
	data1 "pip_toolkit_url_shortener/data/version1"
	persist "pip_toolkit_url_shortener/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AliasPersisteceFixture struct {
	ALIAS1      *data1.AliasV1
	ALIAS2      *data1.AliasV1
	ALIAS3      *data1.AliasV1
	persistence persist.IAliasPersistance
}

func NewAliasPersistenceFixture(persistence persist.AliasSqlitePersistence) *AliasPersisteceFixture {
	c := AliasPersisteceFixture{}

	c.ALIAS1 = &data1.AliasV1{
		Id:    "1",
		Alias: "translate",
		Url:   "https://translate.google.com/",
	}

	c.ALIAS2 = &data1.AliasV1{
		Id:    "2",
		Alias: "youtube",
		Url:   "https://www.youtube.com/",
	}

	c.ALIAS3 = &data1.AliasV1{
		Id:    "3",
		Alias: "hub",
		Url:   "https://github.com/",
	}

	c.persistence = &persistence
	return &c
}

func (c *AliasPersisteceFixture) testCreateAlias(t *testing.T) data1.AliasV1 {
	// Create the first alias
	alias, err := c.persistence.Create(context.Background(), "", *c.ALIAS1)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS1.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS1.Url, alias.Url)

	// Creating a duplicate alias
	alias, err = c.persistence.Create(context.Background(), "", *c.ALIAS1)
	assert.NotNil(t, err)
	assert.Equal(t, data1.AliasV1{}, alias)

	// Create the second alias
	alias, err = c.persistence.Create(context.Background(), "", *c.ALIAS2)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS2.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS2.Url, alias.Url)

	// Create the third alias
	alias, err = c.persistence.Create(context.Background(), "", *c.ALIAS3)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS3.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS3.Url, alias.Url)

	return alias
}

func (c *AliasPersisteceFixture) TestCRDOperations(t *testing.T) {
	// Create items
	alias3 := c.testCreateAlias(t)

	// Get url by alias
	alias, err := c.persistence.GetOneByKey(context.Background(), "", alias3.Alias)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, alias3.Id, alias.Id)

	// Delete the alias
	id, err := c.persistence.DeleteByKey(context.Background(), "", alias3.Alias)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, alias3.Id, id)

	// Try to get deleted alias
	alias, err = c.persistence.GetOneByKey(context.Background(), "", alias3.Alias)
	assert.Nil(t, err)
	assert.Equal(t, data1.AliasV1{}, alias)
}
