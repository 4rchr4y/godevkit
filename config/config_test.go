package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustGetString(t *testing.T) {
	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "test string")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := MustGetString("TEST_VALID_STRING")

		assert.Equal(t, "test string", value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetString("TEST_NONEXISTENT_KEY")
		})
	})
}

func TestGetStringWithDefault(t *testing.T) {
	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "test string")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := GetStringWithDefault("TEST_VALID_STRING", "default string")

		assert.Equal(t, "test string", value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetStringWithDefault("TEST_INVALID_STRING", "default string")

		assert.NotEqual(t, value, "test string")
		assert.Equal(t, value, "default string")

	})
}

func TestMustGetInt(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "10")
		defer os.Unsetenv("TEST_VALID_INT")

		assert.Equal(t, 10, MustGetInt("TEST_VALID_INT"))
	})

	t.Run("invalid", func(t *testing.T) {
		assert.Panics(t, func() { MustGetInt("TEST_VALID_INT") })
	})
}

func TestGetIntWithDefault(t *testing.T) {
	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "10")
		defer os.Unsetenv("TEST_VALID_INT")

		assert.Equal(t, 10, GetIntWithDefault("TEST_VALID_INT", 10))
	})

	t.Run("invalid: no env key", func(t *testing.T) {
		assert.Equal(t, 10, GetIntWithDefault("TEST_VALID_INT", 10))
	})
}

func TestMustGetUint(t *testing.T) {
	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", "10")
		defer os.Unsetenv("TEST_VALID_UINT")

		value := MustGetUint("TEST_VALID_UINT")

		assert.Equal(t, uint64(10), value)
	})

	t.Run("invalid: invalid env value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "-10")
		defer os.Unsetenv("TEST_INVALID_UINT")

		assert.Panics(t, func() {
			MustGetUint("TEST_INVALID_UINT")
		})
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", "10")
		defer os.Unsetenv("TEST_VALID_UINT")

		assert.Panics(t, func() {
			MustGetUint("TEST_WRONG_KEY")
		})
	})
}

func TestGetUintWithDefault(t *testing.T) {
	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", "10")
		defer os.Unsetenv("TEST_VALID_UINT")

		value := GetUintWithDefault("TEST_VALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_STRING", "10")
		defer os.Unsetenv("TEST_INVALID_STRING")

		value := GetUintWithDefault("TEST_INVALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})
}

func TestMustGetFloat64(t *testing.T) {
	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := MustGetFloat64("TEST_VALID_FLOAT64")

		assert.Equal(t, float64(10.0), value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetFloat64("TEST_NONEXISTENT_KEY")
		})
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		assert.Panics(t, func() {
			MustGetFloat64("TEST_WRONG_KEY")
		})
	})
}
func TestGetFloat64WithDefault(t *testing.T) {
	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_VALID_FLOAT64", 5.0)

		assert.Equal(t, float64(10.0), value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetFloat64WithDefault("TEST_INVALID_FLOAT64", 5.0)

		assert.Equal(t, float64(5.0), value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_INVALID_ENV_KEY", 5.0)

		assert.Equal(t, float64(5.0), value)
	})
}
func TestMustGetUrl(t *testing.T) {
	t.Run("valid: valid url", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://github.com/4rchr4y")
		defer os.Unsetenv("TEST_VALID_URL")

		value := MustGetUrl("TEST_VALID_URL")
		assert.Equal(t, "https://github.com/4rchr4y", value)
	})

	t.Run("invalid: invalid env value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "test string")
		defer os.Unsetenv("TEST_INVALID_URL")

		assert.Panics(t, func() {
			MustGetUrl("TEST_INVALID_URL")
		})
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetUrl("TEST_NONEXISTENT_KEY")
		})
	})

}

func TestGetUrlWithDefault(t *testing.T) {
	t.Run("valid: valid url", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://github.com/4rchr4y")
		defer os.Unsetenv("TEST_VALID_URL")

		value := GetUrlWithDefault("TEST_VALID_URL", "www.google.com")
		assert.Equal(t, "https://github.com/4rchr4y", value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetUrlWithDefault("TEST_INVALID_ENV", "https://github.com/4rchr4y")

		assert.Equal(t, "https://github.com/4rchr4y", value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "test string")
		defer os.Unsetenv("TEST_INVALID_URL")

		value := GetUrlWithDefault("TEST_INVALID_ENV", "https://github.com/4rchr4y")

		assert.Equal(t, "https://github.com/4rchr4y", value)
	})
}

func TestMustGetDuration(t *testing.T) {
	t.Run("valid: hour input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Hour), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_TIME")

		assert.Panics(t, func() {
			MustGetDuration("TEST_INVALID_TIME")
		})
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetDuration("TEST_NONEXISTENT_KEY")
		})
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "10")
		defer os.Unsetenv("TEST_INVALID_TIME")

		assert.Panics(t, func() {
			MustGetDuration("TEST_WRONG_KEY")
		})
	})
}

func TestGetDurationWithDefault(t *testing.T) {
	t.Run("valid: hour input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Hour), value)
	})
	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "10s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_INVALID_ENV", time.Second)

		assert.Equal(t, time.Duration(time.Second), value)
	})
}
