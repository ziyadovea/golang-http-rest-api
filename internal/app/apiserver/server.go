package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/model"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/store"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	sessionName = "session-name"
)

var (
	errorIncorrectEmailOrPassword = errors.New("неправильный email или пароль")
)

// server - структура для сервера (внутренние методы)
type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

// newServer создает экземпляр сервера
func newServer(store store.Store, sessionsStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionsStore,
	}
	return s
}

// ServerHTTP обрабатывает запросы, нужно для реализации интферфейса Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter задает конфигурацию для роутера
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handlerUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handlerSessionsCreate()).Methods("POST")
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

// handlerUsersCreate - регистрация. обработчик "/users"
func (s *server) handlerUsersCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		s.logger.Info("/users")

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error("некорректный формат запроса: " + err.Error())
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		// хешируем пароль
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			s.logger.Error("ошибка при хешировании пароля: " + err.Error())
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u := &model.User{
			Email:             req.Email,
			EncryptedPassword: string(hashBytes),
		}
		if err := s.store.User().Create(u); err != nil {
			s.logger.Error("ошибка при создании пользователя: " + err.Error())
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.logger.Info("новый пользователь успешно создан, email = " + u.Email)
		s.respond(w, r, http.StatusCreated, u)
	}

}

// handlerSessionsCreate - авторизация. обработчик "/sessions"
func (s *server) handlerSessionsCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		s.logger.Info("/sessions")

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error("некорректный формат запроса: " + err.Error())
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.logger.Error("неправильный логин или пароль")
			s.error(w, r, http.StatusUnauthorized, errorIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.logger.Error("проблема при создании сессии: " + err.Error())
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.ID
		if err = s.sessionStore.Save(r, w, session); err != nil {
			s.logger.Error("проблема при сохранении сессии: " + err.Error())
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.logger.Info("пользователь успешно прошел аутентификацию, email = " + u.Email)
		s.respond(w, r, http.StatusOK, nil)
	}

}

// error - функция-хелпер для обработки ошибок
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{
		"error": err.Error(),
	})
}

// respond - более абстрактная функция-хелпер, которого будет отдавать результат работы сервера клиенту
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
