package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost user=postgres password=1 dbname=restapi sslmode=disable"
	}

	os.Exit(m.Run())
}
