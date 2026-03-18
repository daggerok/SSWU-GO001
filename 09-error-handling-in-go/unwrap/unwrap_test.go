package unwrap_test

import (
	"errors"
	"testing"

	"github.com/daggerok/SSWU-GO001/09-error-handling-in-go/unwrap"
	"github.com/stretchr/testify/assert"
)

func TestUnwrap(t *testing.T) {
	t.Run("should unwrap error", func(t *testing.T) {
		_, err := unwrap.HttpGet()
		t.Log(err)

		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, unwrap.ErrorNotFound))
		assert.Equal(t, "Operation failed: Not Found", err.Error())
	})
}
