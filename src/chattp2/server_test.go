package chattp2_test

import (
	"testing"

	"github.com/neomede/chattp2/src/chattp2"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	a := assert.New(t)

	server, err := chattp2.NewServer()
	a.NoError(err)
	a.NotNil(server)
}
