package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

// TestStore - функция-хелпер для тестирования запросов к БД
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	store := New(config)

	if err := store.Open(); err != nil {
		t.Fatal(err)
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			_, err := store.connection.Exec(
				context.Background(),
				fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", ")),
			)
			if err != nil {
				t.Fatal(err)
			}
		}
		store.Close()
	}
}
