package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	// success
	s, err := NewServer(9000)

	if assert.NoError(t, err) {
		assert.Equal(t, &Server{
			port: 9000,
		}, s)
	}

	// failure
	s, err = NewServer(15000)

	if assert.Error(t, err) {
		assert.Equal(t, "NewServer: available port out of range", err.Error())
	}
}
