package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()

	assert.Equal(t, &Server{}, s)
}
