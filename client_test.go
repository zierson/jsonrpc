package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	assert.Equal(t, &Client{}, c)
}
