package store

// Config - конфиг для БД
type Config struct {
	DatabaseURL string `json:"database_url"`
}

// NewConfig оздает новый конфиг для БД
func NewConfig() *Config {
	return &Config{
		DatabaseURL: "postgres://postgres:0000@localhost:5432/users_go_restapi",
	}
}
