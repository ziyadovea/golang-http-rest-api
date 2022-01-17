# golang-http-rest-api

Приложение представляет собой REST API на языке Golang. <br>
Архитектура приложения строится согласно ["Standard Go Project Layout"](https://github.com/golang-standards/project-layout). <br>
В приложении реализована регистрация, аутентификация на основе сессий и авторизация пользователей, а также их выход из аккаунта. <br>
В качестве БД используется PostgreSQL.

Для работы с HTTP используется стандартный пакет "net/http", а также ["gorilla/mux"](https://github.com/gorilla/mux). <br>
Для работы с БД - пакет ["pgx/v4"](https://pkg.go.dev/github.com/jackc/pgx/v4). <br>
Для логирования используется пакет ["logrus"](https://github.com/sirupsen/logrus). <br>
Для валидации данных используется пакет ["go-playground/validator/v10"](https://github.com/go-playground/validator).

