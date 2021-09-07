package main

import (
	"github.com/zierson/jsonrpc"
	"time"
)

func main() {
	s, err := jsonrpc.NewServer(9000)
	if err != nil {
		panic(err)
	}

	if err := s.AddMethod("test", func(request *jsonrpc.Request) {
		// some code
	}); err != nil {
		panic(err)
	}

	if err := s.Bind(); err != nil {
		panic(err)
	}

	// TODO: refactor
	time.Sleep(60 * time.Second)
}
