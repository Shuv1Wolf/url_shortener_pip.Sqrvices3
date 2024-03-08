package test_persistance

import (
	"context"
	"os"
	persist "pip_toolkit_url_shortener/persistence"
	"testing"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type AliasSqlitePersistenceTest struct {
	persistence *persist.AliasSqlitePersistence
	fixture     *AliasPersisteceFixture
}

func newAliasSqlitePersistenceTest() *AliasSqlitePersistenceTest {
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
	fixture := NewAliasPersistenceFixture(*persistence)

	return &AliasSqlitePersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *AliasSqlitePersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open(context.Background(), "")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear(context.Background(), "")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *AliasSqlitePersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close(context.Background(), "")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestAliasSqlitePersistence(t *testing.T) {
	c := newAliasSqlitePersistenceTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRD Operations", c.fixture.TestCRDOperations)
	c.teardown(t)
}
