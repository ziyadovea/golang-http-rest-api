package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
	"net/http"
)

// server - структура для сервера (внутренние методы)
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// newServer создает экземпляр сервера
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	return s
}

// ServerHTTP обрабатывает запросы, нужно для реализации интферфейса Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter задает конфигурацию для роутера
func (s *server) configureRouter() {

}

// configureLogger задает конфигурацию для логгера
func (s *server) configureLogger(config *Config) error {
	lvl, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lvl)
	return nil
}
