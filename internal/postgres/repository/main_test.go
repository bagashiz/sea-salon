package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/bagashiz/sea-salon/internal/postgres"
	"github.com/testcontainers/testcontainers-go"
	pgtestcontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// testDB is a db connection instance for testing
var testDB *postgres.DB

// TestMain sets up the test database connection and runs the tests
func TestMain(m *testing.M) {
	ctx := context.Background()

	postgresContainer, err := setupPostgresContainer(ctx)
	if err != nil {
		panic(err)
	}
	defer terminatePostgresContainer(ctx, postgresContainer)

	connURI, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		panic(err)
	}

	testDB, err = postgres.Connect(ctx, connURI)
	if err != nil {
		panic(err)
	}

	if err := testDB.Migrate(); err != nil {
		panic(err)
	}

	exitCode := m.Run()

	testDB.Close()

	os.Exit(exitCode)
}

// setupPostgresContainer sets up a postgres test container for testing
func setupPostgresContainer(ctx context.Context) (*pgtestcontainer.PostgresContainer, error) {
	dbImage := "docker.io/postgres:16-alpine"
	dbName := "sea_salon"
	dbUser := "postgres"
	dbPassword := "password"

	postgresContainer, err := pgtestcontainer.Run(ctx,
		dbImage,
		pgtestcontainer.WithDatabase(dbName),
		pgtestcontainer.WithUsername(dbUser),
		pgtestcontainer.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	return postgresContainer, nil
}

// terminatePostgresContainer deletes the postgres test container after testing
func terminatePostgresContainer(ctx context.Context, container *pgtestcontainer.PostgresContainer) {
	if err := container.Terminate(ctx); err != nil {
		panic(err)
	}
}
