package chattp2_test

import (
	"testing"

	"github.com/neomede/chattp2/src/chattp2"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	a := assert.New(t)

	client, err := chattp2.NewClient("sender", "receiver")
	a.NoError(err)
	a.NotNil(client)
}
