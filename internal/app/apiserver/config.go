package apiserver

// Config - структура для конфига приложения
type Config struct {
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
}

// NewConfig создает новый конфиг
func NewConfig() *Config {
	return &Config{
		Port:     ":8080",
		LogLevel: "debug",
	}
}
