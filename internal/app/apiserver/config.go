package apiserver

type Config struct {
	Port string `json:"port"`
}

// NewConfig создает новый конфиг
func NewConfig() *Config {
	return &Config{}
}
