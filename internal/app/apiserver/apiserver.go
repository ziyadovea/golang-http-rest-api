package apiserver

// APIServer - структура для сервера
type APIServer struct {
}

// New создает экземпляр сервера
func New() *APIServer {
	return &APIServer{}
}

// Start запускает сервер
func (s *APIServer) Start() error {
	return nil
}
