package setting

import "os"

type Server struct {
	Port string
}

func (s *Server) Setup() *Server {
	s.Port = os.Getenv("HTTP_PORT")

	return s
}
