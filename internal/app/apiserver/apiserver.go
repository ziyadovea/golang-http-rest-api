package apiserver

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

// Start запускает сервер
func Start(config *Config) error {

	conn, err := connectDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	store := sqlstore.New(conn)

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	server := newServer(store, sessionStore)
	server.configureLogger(config)
	server.configureRouter()

	server.logger.Info("starting server on port " + config.Port)
	return http.ListenAndServe(config.Port, server)

}

// connectDB коннектиться с БД
func connectDB(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	return conn, nil
}
