package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/bagashiz/sea-salon/internal/config"
	"github.com/bagashiz/sea-salon/internal/postgres"
)

// testDB is a db connection instance for testing
var testDB *postgres.DB

// TestMain sets up the test database connection and runs the tests
func TestMain(m *testing.M) {
	ctx := context.Background()
	config, err := config.New(os.Getenv, "../../../")
	if err != nil {
		panic(err)
	}

	testDB, err = postgres.NewDB(ctx, config.DB)
	if err != nil {
		panic(err)
	}

	if err := testDB.Migrate(config.DB.Type); err != nil {
		panic(err)
	}

	exitCode := m.Run()

	testDB.Close()

	os.Exit(exitCode)
}
