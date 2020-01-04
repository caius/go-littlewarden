package littlewarden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("secret")

	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "secret", c.ApiKey)
	assert.Equal(t, defaultApiURL, c.BaseURL.String())
}
