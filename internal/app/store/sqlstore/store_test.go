package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://levchenko.an:K1~hubnt@127.0.0.1:5432/willknow_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
