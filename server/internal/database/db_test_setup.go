package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/pressly/goose"
)

func DBTestSetup() (queries *Queries, cleanup func()) {
	var db *sql.DB
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{"POSTGRES_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("host=localhost port=%s user=postgres password=secret dbname=postgres sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	// use goose API to send migrations to test environment
	if err != nil {
		log.Fatalf("Could not resolve absolute path for migrations: %s", err)
	}
	err = goose.Up(db, "../../sql/schema")
	if err != nil {
		log.Fatalf("Could not run migrations: %s", err)
	}
	// configure db to work with sqlc queries
	queries = New(db)
	// as of go1.15 testing.M returns the exit code of m.Run(), so it is safe to use defer here
	cleanup = func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
	return queries, cleanup
}
