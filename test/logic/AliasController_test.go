package logic_test

import (
	"context"
	"os"
	data1 "pip_toolkit_url_shortener/data/version1"
	logic "pip_toolkit_url_shortener/logic"
	persist "pip_toolkit_url_shortener/persistence"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/stretchr/testify/assert"
)

type aliasControllerTest struct {
	ALIAS1      *data1.AliasV1
	ALIAS2      *data1.AliasV1
	persistence *persist.AliasSqlitePersistence
	controller  *logic.AliasController
}

func newAliasControllerTest() *aliasControllerTest {
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

	references := cref.NewReferencesFromTuples(
		context.Background(),
		cref.NewDescriptor("alias", "persistence", "sqlite", "default", "1.0"), persistence,
		cref.NewDescriptor("alias", "controller", "default", "default", "1.0"), controller,
	)

	controller.SetReferences(context.Background(), references)

	return &aliasControllerTest{
		ALIAS1:      ALIAS1,
		ALIAS2:      ALIAS2,
		persistence: persistence,
		controller:  controller,
	}
}

func (c *aliasControllerTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *aliasControllerTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func (c *aliasControllerTest) testCRDOperations(t *testing.T) {
	var alias1 data1.AliasV1

	// Create the first alias
	alias, err := c.persistence.Create(context.Background(), "", c.ALIAS1.Clone())
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS1.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS1.Url, alias.Url)

	alias1 = alias

	// Creating a duplicate alias
	alias, err = c.persistence.Create(context.Background(), "", c.ALIAS1.Clone())
	assert.NotNil(t, err)
	assert.Equal(t, data1.AliasV1{}, alias)

	// Create the second alias
	alias, err = c.persistence.Create(context.Background(), "", c.ALIAS2.Clone())
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, c.ALIAS2.Alias, alias.Alias)
	assert.Equal(t, c.ALIAS2.Url, alias.Url)

	// Get url by alias
	alias, err = c.persistence.GetOneByKey(context.Background(), "", alias1.Alias)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, alias1.Id, alias.Id)

	// Delete the alias
	id, err := c.persistence.DeleteByKey(context.Background(), "", alias1.Alias)
	assert.Nil(t, err)
	assert.NotEqual(t, data1.AliasV1{}, alias)
	assert.Equal(t, alias1.Id, id)

	// Try to get deleted alias
	alias, err = c.persistence.GetOneByKey(context.Background(), "", alias1.Alias)
	assert.Nil(t, err)
	assert.Equal(t, data1.AliasV1{}, alias)

}

func TestAliasSqlitePersistence(t *testing.T) {
	c := newAliasControllerTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRD Operations", c.testCRDOperations)
	c.teardown(t)
}
