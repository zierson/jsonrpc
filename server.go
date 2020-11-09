package jsonrpc

type (
	Server struct {
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

func NewServer() *Server {
	return &Server{}
}
