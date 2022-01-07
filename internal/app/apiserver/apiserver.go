package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
	"io"
	"net/http"
)

// APIServer - структура для сервера
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	Store  *store.Store
}

// New создает экземпляр сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start запускает сервер
func (s *APIServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.Port, s.router)

}

// configureLogger задает конфигурацию для логгера
func (s *APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lvl)
	return nil
}

// configureRouter задает конфигурацию для роутера
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// configureStore задает конфигурация для БД
func (s *APIServer) configureStore() error {
	st := store.New(s.config.StoreConfig)
	if err := st.Open(); err != nil {
		return err
	}
	s.Store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
