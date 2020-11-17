package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	// success
	s, err := NewServer(9000)

	if assert.NoError(t, err) {
		assert.Equal(t, 9000, int(s.port))
		assert.NotEmpty(t, s.httpSrv)
	}

	// failure
	s, err = NewServer(15000)

	if assert.Error(t, err) {
		assert.Equal(t, "NewServer: available port out of range", err.Error())
	}
}

func TestServer_AddMethod(t *testing.T) {

}

func TestServer_Bind(t *testing.T) {
	// success
	s, err := NewServer(9000)

	if assert.NoError(t, err) {
		if assert.NoError(t, s.Bind()) {
			time.Sleep(time.Second)

			assert.NoError(t, s.Stop())
		}
	}
}
