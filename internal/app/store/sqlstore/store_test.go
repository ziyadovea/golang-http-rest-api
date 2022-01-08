package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// TestMain устанавливает databaseURL перед остальными тестами
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:0000@localhost:5432/test"
	}
	os.Exit(m.Run())
}
