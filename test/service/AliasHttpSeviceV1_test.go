package service1_test

import (
	"context"
	"os"
	data1 "pip_toolkit_url_shortener/data/version1"
	logic "pip_toolkit_url_shortener/logic"
	persist "pip_toolkit_url_shortener/persistence"
	service1 "pip_toolkit_url_shortener/service/version1"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
	tclients "github.com/pip-services3-gox/pip-services3-rpc-gox/test"
	"github.com/stretchr/testify/assert"
)

type aliasHttpServiceV1Test struct {
	ALIAS1      *data1.AliasV1
	ALIAS2      *data1.AliasV1
	persistence *persist.AliasSqlitePersistence
	controller  *logic.AliasController
	service     *service1.AliasHttpServiceV1
	client      *tclients.TestCommandableHttpClient
}

func newAliasHttpServiceV1Test() *aliasHttpServiceV1Test {
	ALIAS1 := &data1.AliasV1{
		Id:    "1",
		Alias: "translate",
		Url:   "https://translate.google.com/",
	}

	ALIAS2 := &data1.AliasV1{
		Id:    "2",
		Alias: "youtube",
		Url:   "https://www.youtube.com/",
	}

	restConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3001",
		"connection.host", "localhost",
	)

	sqliteDatabase := os.Getenv("SQLITE_DB")
	if sqliteDatabase == "" {
		sqliteDatabase = "../../storage/storage.db"
	}

	if sqliteDatabase == "" {
		panic("Connection params losse")
	}

	dbConfig := cconf.NewConfigParamsFromTuples(
		"connection.database", sqliteDatabase,
	)

	persistence := persist.NewAliasSqlPersistance()
	persistence.Configure(context.Background(), dbConfig)

	controller := logic.NewAliasController()

	service := service1.NewAliasHttpServiceV1()
	service.Configure(context.Background(), restConfig)

	client := tclients.NewTestCommandableHttpClient("v1/alias")
	client.Configure(context.Background(), restConfig)

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("alias", "persistence", "sqlite", "default", "1.0"), persistence,
		cref.NewDescriptor("alias", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("alias", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("alias", "client", "http", "default", "1.0"), client,
	)

	controller.SetReferences(context.Background(), references)
	service.SetReferences(context.Background(), references)

	return &aliasHttpServiceV1Test{
		ALIAS1:      ALIAS1,
		ALIAS2:      ALIAS2,
		persistence: persistence,
		controller:  controller,
		service:     service,
		client:      client,
	}
}

func (c *aliasHttpServiceV1Test) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.service.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open service", err)
	}

	err = c.client.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open client", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *aliasHttpServiceV1Test) teardown(t *testing.T) {
	err := c.client.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close client", err)
	}

	err = c.service.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close service", err)
	}

	err = c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *aliasHttpServiceV1Test) testCRDOperations(t *testing.T) {

	// Create the first alias
	params := cdata.NewAnyValueMapFromTuples(
		"alias", c.ALIAS1.Clone(),
	)
	response, err := c.client.CallCommand(context.Background(), "create_alias", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	alias, err := cclients.HandleHttpResponse[data1.AliasV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS1.Id, alias.Id)
	assert.Equal(t, c.ALIAS1.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS1.Url, alias.Url)

	// Creating a duplicate alias
	params = cdata.NewAnyValueMapFromTuples(
		"alias", c.ALIAS1.Clone(),
	)
	response, err = c.client.CallCommand(context.Background(), "create_alias", "", params)
	assert.NotNil(t, err)

	alias, err = cclients.HandleHttpResponse[data1.AliasV1](response, "")
	assert.Nil(t, err)
	assert.Equal(t, data1.AliasV1{}, alias)

	// Create the second beacon
	params = cdata.NewAnyValueMapFromTuples(
		"alias", c.ALIAS2.Clone(),
	)
	response, err = c.client.CallCommand(context.Background(), "create_alias", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	alias, err = cclients.HandleHttpResponse[data1.AliasV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS2.Id, alias.Id)
	assert.Equal(t, c.ALIAS2.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS2.Url, alias.Url)

	// Get url by alias1
	params = cdata.NewAnyValueMapFromTuples(
		"alias", c.ALIAS1.Clone().Alias,
	)
	response, err = c.client.CallCommand(context.Background(), "get_one_by_alias", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	alias, err = cclients.HandleHttpResponse[data1.AliasV1](response, "")
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS1.Id, alias.Id)
	assert.Equal(t, c.ALIAS1.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS1.Url, alias.Url)

	// Delete the alias
	params = cdata.NewAnyValueMapFromTuples(
		"alias", c.ALIAS1.Clone().Alias,
	)
	response, err = c.client.CallCommand(context.Background(), "delete_by_alias", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	id, err := cclients.HandleHttpResponse[string](response, "")
	assert.Nil(t, err)
	assert.Equal(t, c.ALIAS1.Id, id)

	// Try to get deleted alias
	params = cdata.NewAnyValueMapFromTuples(
		"alias", c.ALIAS1.Clone().Alias,
	)
	response, err = c.client.CallCommand(context.Background(), "get_one_by_alias", "", params)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	alias, err = cclients.HandleHttpResponse[data1.AliasV1](response, "")
	assert.Nil(t, err)
	assert.Equal(t, data1.AliasV1{}, alias)

}

func TestAliasCommmandableHttpServiceV1(t *testing.T) {
	c := newAliasHttpServiceV1Test()

	c.setup(t)
	t.Run("CRD Operations", c.testCRDOperations)
	c.teardown(t)
}
