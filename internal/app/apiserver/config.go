package apiserver

// Config - структура для конфига приложения
type Config struct {
	Port        string `json:"port"`
	LogLevel    string `json:"log_level"`
	DatabaseURL string `json:"database_url"`
	SessionKey  string `json:"session_key"`
}

// NewConfig создает новый конфиг
func NewConfig() *Config {
	return &Config{
		Port:        ":8080",
		LogLevel:    "debug",
		DatabaseURL: "postgres://postgres:0000@localhost:5432/go_restapi",
	}
}
