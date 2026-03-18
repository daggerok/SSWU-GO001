package custom_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/09-error-handling-in-go/custom"
	"github.com/stretchr/testify/assert"
)

func TestCustomError(t *testing.T) {
	t.Run("should handle custom error", func(t *testing.T) {
		_, err := custom.HttpGet()
		t.Log(err)
		assert.NotNil(t, err)
		assert.Equal(t, "Error 404: Not Found", err.Error())
	})
}
