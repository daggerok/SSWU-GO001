package wrap_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/09-error-handling-in-go/wrap"
	"github.com/stretchr/testify/assert"
)

func TestWrapAndContextualErrors(t *testing.T) {
	t.Run("should wrap error with context", func(t *testing.T) {
		err := wrap.WrapError()
		t.Log(err)
		assert.NotNil(t, err)
		assert.Equal(t, "Failed to do something: Error 404: Not Found", err.Error())
	})
}
