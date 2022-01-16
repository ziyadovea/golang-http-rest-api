# golang-http-rest-api

Приложение представляет собой REST API на языке Golang. <br>
На текущий момент реализована архитектура приложения, коннект с БД PostgreSQL, часть кода покрыта тестами.

Для работы с HTTP используется стандартный пакет "net/http", а также ["gorilla/mux"](https://github.com/gorilla/mux). <br>
Для работы с БД - пакет ["pgx/v4"](https://pkg.go.dev/github.com/jackc/pgx/v4). <br>
Для логирования используется пакет ["logrus"](https://github.com/sirupsen/logrus). <br>
Для валидации данных используется пакет ["go-playground/validator/v10"](https://github.com/go-playground/validator).

todolist:
* Регистрация, авторизация и аутентификация на основе сессий.
