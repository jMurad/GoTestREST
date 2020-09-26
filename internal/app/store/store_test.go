package store_test

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
		databaseURL = "host=localhost dbname=test_db_EXAMPLE user=test_user password=M353694 sslmode=disable"
	}

	os.Exit(m.Run())
}
