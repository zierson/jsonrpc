package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	// success
	c, err := NewClient("http://localhost:9000/rpc")

	if assert.NoError(t, err) {
		assert.Equal(t, &Client{
			endpoint: "http://localhost:9000/rpc",
		}, c)
	}

	// failure
	c, err = NewClient("/rpc")

	if assert.Error(t, err) {
		assert.Equal(t, "AddEndpoint: invalid http url", err.Error())
	}
}
