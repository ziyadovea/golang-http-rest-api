package apiserver

import "github.com/ziyadovea/golang-http-rest-api/internal/app/store"

// Config - структура для конфига приложения
type Config struct {
	Port        string        `json:"port"`
	LogLevel    string        `json:"log_level"`
	StoreConfig *store.Config `json:"store_config"`
}

// NewConfig создает новый конфиг
func NewConfig() *Config {
	return &Config{
		Port:        ":8080",
		LogLevel:    "debug",
		StoreConfig: store.NewConfig(),
	}
}
