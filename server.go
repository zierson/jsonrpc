package jsonrpc

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type (
	Server struct {
		port    uint64
		httpSrv *fasthttp.Server
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

	s := &Server{
		port: port,
		httpSrv: &fasthttp.Server{
			Name:               "jsonrpc",
			Concurrency:        1024,
			DisableKeepalive:   true,
			ReadTimeout:        10 * time.Second,
			WriteTimeout:       10 * time.Second,
			MaxConnsPerIP:      256,
			MaxRequestsPerConn: 1,
			MaxRequestBodySize: 16 * 1024 * 1024,
			LogAllErrors:       true,
			//Logger: ,
		},
	}

	return s, nil
}

func (s *Server) AddMethod(name string, fn func(*Request)) error {
	return nil
}

func (s *Server) Bind() error {
	h := func(ctx *fasthttp.RequestCtx) {
		ctx.SetBodyString("qweqwe")
	}

	s.httpSrv.Handler = fasthttp.TimeoutWithCodeHandler(h, 10*time.Second, fasthttp.StatusMessage(fasthttp.StatusRequestTimeout), fasthttp.StatusRequestTimeout)

	go func(s *Server) {
		p := strconv.Itoa(int(s.port))
		if err := s.httpSrv.ListenAndServe("0.0.0.0:" + p); err != nil {
			panic(errors.Wrap(err, "Server_Bind"))
		}
	}(s)

	return nil
}

func (s *Server) Stop() error {
	if err := s.httpSrv.Shutdown(); err != nil {
		return errors.Wrap(err, "Server_Stop")
	}

	return nil
}
