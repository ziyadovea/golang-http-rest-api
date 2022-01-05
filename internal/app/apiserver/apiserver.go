package apiserver

// APIServer - структура для сервера
type APIServer struct {
	config *Config
}

// New создает экземпляр сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start запускает сервер
func (s *APIServer) Start() error {
	return nil
}
