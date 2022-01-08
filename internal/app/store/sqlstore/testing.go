package sqlstore

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
	"testing"
)

// TestConnection - функция-хелпер для тестирования запросов к БД
func TestConnection(t *testing.T, databaseURL string) (*pgx.Conn, func(...string)) {
	t.Helper()

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	// Пинг для проверки соединения
	if err := conn.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	return conn, func(tables ...string) {
		if len(tables) > 0 {
			_, err := conn.Exec(
				context.Background(),
				fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", ")),
			)
			if err != nil {
				t.Fatal(err)
			}
		}
		conn.Close(context.Background())
	}
}
