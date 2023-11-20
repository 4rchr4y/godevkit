package must

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("valid: valid string and no error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Must("valid string", nil)
		})
	})

	t.Run("valid: valid int and no error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Must(10, nil)
		})
	})

	t.Run("invalid: errors.New()", func(t *testing.T) {
		assert.Panics(t, func() {
			Must("valid string", errors.New("asd"))
		})
	})

	t.Run("invalid: fmt.Errorf", func(t *testing.T) {
		assert.Panics(t, func() {
			Must("valid string", fmt.Errorf("asd"))
		})
	})
}

func TestBeOk(t *testing.T) {
	t.Run("valid: true boolean input", func(t *testing.T) {
		assert.NotPanics(t, func() {
			MustBeOk("valid string", true)
		})
	})

	t.Run("invalid: false boolean input", func(t *testing.T) {
		assert.Panics(t, func() {
			MustBeOk("valid string", false)
		})
	})
}

func TestNotBeOk(t *testing.T) {
	t.Run("valid: false boolean input", func(t *testing.T) {
		assert.NotPanics(t, func() {
			MustNotBeOk("valid string", false)
		})
	})

	t.Run("invalid: true boolean input", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNotBeOk("valid string", true)
		})
	})
}
