package jsonrpc

import "github.com/pkg/errors"

type (
	Server struct {
		port uint64
	}

	Response struct {
		Version string      `json:"version"`
		Result  interface{} `json:"result"`
		Id      uint64      `json:"id"`
	}

	ResponseError struct {
		Version string                 `json:"version"`
		Error   map[string]interface{} `json:"error"`
		Id      uint64                 `json:"id"`
	}
)

func NewServer(port uint64) (*Server, error) {
	if port < 8000 || port >= 10000 {
		return nil, errors.WithMessage(errors.New("available port out of range"), "NewServer")
	}

	return &Server{
		port: port,
	}, nil
}

func (s *Server) AddMethod(name string, fn func(*Request)) error {
	return nil
}

func (s *Server) Bind() error {
	return nil
}
