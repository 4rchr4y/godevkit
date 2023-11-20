package config

import (
	"math"
	"os"
	"strconv"
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

	t.Run("valid: empty string as a valid value", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := MustGetString("TEST_VALID_STRING")

		assert.Equal(t, "", value)
	})

	t.Run("valid: special characters in value", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "value@123!")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := MustGetString("TEST_VALID_STRING")

		assert.Equal(t, "value@123!", value)
	})

	t.Run("valid: numeric value as string", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "12345")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := MustGetString("TEST_VALID_STRING")

		assert.Equal(t, "12345", value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetString("TEST_NONEXISTENT_KEY")
		})
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_STRING", "test string")
		defer os.Unsetenv("TEST_INVALID_STRING")

		assert.Panics(t, func() {
			MustGetString("TEST_WRONG_KEY")
		})
	})

	t.Run("invalid: case sensitivity of a variable", func(t *testing.T) {
		os.Setenv("TEST_INVALID_STRING", "Test String")
		defer os.Unsetenv("TEST_INVALID_STRING")

		value := MustGetString("TEST_INVALID_STRING")

		assert.NotEqual(t, value, "test string")
	})
}

func TestGetStringWithDefault(t *testing.T) {
	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "test string")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := GetStringWithDefault("TEST_VALID_STRING", "fallback string")

		assert.Equal(t, "test string", value)
	})

	t.Run("valid: empty string as a valid value", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := GetStringWithDefault("TEST_VALID_STRING", "fallback string")

		assert.Equal(t, "", value)
	})

	t.Run("valid: special characters in value", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "value@123!")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := GetStringWithDefault("TEST_VALID_STRING", "fallback string")

		assert.Equal(t, "value@123!", value)
	})

	t.Run("valid: numeric value as string", func(t *testing.T) {
		os.Setenv("TEST_VALID_STRING", "12345")
		defer os.Unsetenv("TEST_VALID_STRING")

		value := GetStringWithDefault("TEST_VALID_STRING", "fallback string")

		assert.Equal(t, "12345", value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetStringWithDefault("TEST_NONEXISTENT_KEY", "fallback string")

		assert.Equal(t, "fallback string", value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_STRING", "test string")
		defer os.Unsetenv("TEST_INVALID_STRING")

		value := GetStringWithDefault("TEST_INVALID_KEY", "fallback string")

		assert.Equal(t, "fallback string", value)
	})

	t.Run("invalid: case sensitivity of a variable", func(t *testing.T) {
		os.Setenv("TEST_INVALID_STRING", "Test String")
		defer os.Unsetenv("TEST_INVALID_STRING")

		value := GetStringWithDefault("TEST_INVALID_STRING", "fallback string")

		assert.NotEqual(t, value, "test string")
		assert.NotEqual(t, value, "fallback string")
		assert.Equal(t, "Test String", value)

	})
}

func TestMustGetInt(t *testing.T) {

	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "10")
		defer os.Unsetenv("TEST_VALID_INT")

		value := MustGetInt("TEST_VALID_INT")

		assert.Equal(t, 10, value)
	})

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "007")
		defer os.Unsetenv("TEST_VALID_INT")

		value := MustGetInt("TEST_VALID_INT")

		assert.Equal(t, 7, value)
	})

	t.Run("valid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "-7")
		defer os.Unsetenv("TEST_VALID_INT")

		value := MustGetInt("TEST_VALID_INT")

		assert.Equal(t, -7, value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", strconv.Itoa(math.MaxInt))
		defer os.Unsetenv("TEST_VALID_INT")

		value := MustGetInt("TEST_VALID_INT")

		assert.Equal(t, math.MaxInt, value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_INT", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_INT")

		assert.Panics(t, func() {
			MustGetInt("TEST_INVALID_INT")
		})
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_INT", "10!@")
		defer os.Unsetenv("TEST_INVALID_INT")

		assert.Panics(t, func() {
			MustGetInt("TEST_INVALID_INT")
		})
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetInt("TEST_NONEXISTENT_KEY")
		})
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_INT", "10")
		defer os.Unsetenv("TEST_INVALID_INT")

		assert.Panics(t, func() {
			MustGetInt("TEST_WRONG_KEY")
		})
	})
}

func TestGetIntWithDefault(t *testing.T) {
	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "10")
		defer os.Unsetenv("TEST_VALID_INT")

		value := GetIntWithDefault("TEST_VALID_INT", 5)

		assert.Equal(t, 10, value)
	})

	t.Run("valid: valid fallback", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "")
		defer os.Unsetenv("TEST_VALID_INT")

		value := GetIntWithDefault("TEST_VALID_INT", 5)

		assert.Equal(t, 5, value)
	})

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "007")
		defer os.Unsetenv("TEST_VALID_INT")

		value := GetIntWithDefault("TEST_VALID_INT", 5)

		assert.Equal(t, 7, value)
	})

	t.Run("valid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "-7")
		defer os.Unsetenv("TEST_VALID_INT")

		value := GetIntWithDefault("TEST_VALID_INT", 5)

		assert.Equal(t, -7, value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", strconv.Itoa(math.MaxInt))
		defer os.Unsetenv("TEST_VALID_INT")

		value := GetIntWithDefault("TEST_VALID_INT", 5)

		assert.Equal(t, math.MaxInt, value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_INT", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_INT")

		value := GetIntWithDefault("TEST_INVALID_INT", 5)

		assert.Equal(t, 5, value)
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_INT", "10!@")
		defer os.Unsetenv("TEST_INVALID_INT")

		value := GetIntWithDefault("TEST_INVALID_INT", 5)

		assert.Equal(t, 5, value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetIntWithDefault("TEST_INVALID_KEY", 5)

		assert.Equal(t, 5, value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_INT", "10")
		defer os.Unsetenv("TEST_VALID_INT")

		value := GetIntWithDefault("TEST_INVALID_KEY", 5)

		assert.Equal(t, 5, value)
	})
}

func TestMustGetUint(t *testing.T) {

	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", "10")
		defer os.Unsetenv("TEST_VALID_UINT")

		value := MustGetUint("TEST_VALID_UINT")

		assert.Equal(t, uint(10), value)
	})

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", "007")
		defer os.Unsetenv("TEST_VALID_UINT")

		value := MustGetUint("TEST_VALID_UINT")

		assert.Equal(t, uint(7), value)
	})

	t.Run("valid: large uint value", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", strconv.FormatUint(math.MaxUint64, 10))
		defer os.Unsetenv("TEST_VALID_UINT")

		value := MustGetUint("TEST_VALID_UINT")

		assert.Equal(t, uint(math.MaxUint64), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_UINT")

		assert.Panics(t, func() {
			MustGetUint("TEST_INVALID_UINT")
		})
	})

	t.Run("invalid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "-7")
		defer os.Unsetenv("TEST_INVALID_UINT")

		assert.Panics(t, func() {
			MustGetUint("TEST_INVALID_UINT")
		})
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "10!@")
		defer os.Unsetenv("TEST_INVALID_UINT")

		assert.Panics(t, func() {
			MustGetUint("TEST_INVALID_UINT")
		})
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetUint("TEST_NONEXISTENT_KEY")
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

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", "007")
		defer os.Unsetenv("TEST_VALID_UINT")

		value := GetUintWithDefault("TEST_VALID_UINT", 10)

		assert.Equal(t, uint(7), value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_UINT", strconv.Itoa(math.MaxInt))
		defer os.Unsetenv("TEST_VALID_UINT")

		value := GetUintWithDefault("TEST_VALID_UINT", 10)

		assert.Equal(t, uint(math.MaxInt), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_UINT")

		value := GetUintWithDefault("TEST_INVALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})

	t.Run("invalid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "-7")
		defer os.Unsetenv("TEST_INVALID_UINT")

		value := GetUintWithDefault("TEST_INVALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_UINT", "10!@")
		defer os.Unsetenv("TEST_INVALID_UINT")

		value := GetUintWithDefault("TEST_INVALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetUintWithDefault("TEST_INVALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_STRING", "10")
		defer os.Unsetenv("TEST_INVALID_STRING")

		value := GetUintWithDefault("TEST_INVALID_UINT", 10)

		assert.Equal(t, uint(10), value)
	})
}

func TestMustGetFloat32(t *testing.T) {

	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := MustGetFloat32("TEST_VALID_FLOAT32")

		assert.Equal(t, float32(10.0), value)
	})

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "007")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := MustGetFloat32("TEST_VALID_FLOAT32")

		assert.Equal(t, float32(7.0), value)
	})

	t.Run("valid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "-7")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := MustGetFloat32("TEST_VALID_FLOAT32")

		assert.Equal(t, float32(-7.0), value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", strconv.FormatFloat(math.MaxFloat32, 'f', -1, 64))
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := MustGetFloat32("TEST_VALID_FLOAT32")

		assert.Equal(t, float32(math.MaxFloat32), value)
	})

	t.Run("invalid: value is nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		assert.Panics(t, func() {
			MustGetFloat32("TEST_INVALID_FLOAT32")
		})
	})

	t.Run("invalid: value is math.nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "math.NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		assert.Panics(t, func() {
			MustGetFloat32("TEST_INVALID_FLOAT32")
		})
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "not_a_float32")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		assert.Panics(t, func() {
			MustGetFloat32("TEST_INVALID_FLOAT32")
		})
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "10!@")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		assert.Panics(t, func() {
			MustGetFloat32("TEST_INVALID_FLOAT32")
		})
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		assert.Panics(t, func() {
			MustGetFloat32("TEST_NONEXISTENT_KEY")
		})
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		assert.Panics(t, func() {
			MustGetFloat32("TEST_WRONG_KEY")
		})
	})
}

func TestGetFloat32WithDefault(t *testing.T) {

	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_VALID_FLOAT32", 5.0)

		assert.Equal(t, float32(10.0), value)
	})

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "007")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_VALID_FLOAT32", 5.0)

		assert.Equal(t, float32(7.0), value)
	})

	t.Run("valid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "-7")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_VALID_FLOAT32", 5.0)

		assert.Equal(t, float32(-7.0), value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", strconv.FormatFloat(math.MaxFloat32, 'f', -1, 64))
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_VALID_FLOAT32", 5.0)

		assert.Equal(t, float32(math.MaxFloat32), value)
	})

	t.Run("invalid: value is nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_INVALID_FLOAT32", 5.0)

		assert.Equal(t, float32(5.0), value)
	})

	t.Run("invalid: value is math.nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "math.NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_INVALID_FLOAT32", 5.0)

		assert.Equal(t, float32(5.0), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "not_a_float32")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_INVALID_FLOAT32", 5.0)

		assert.Equal(t, float32(5.0), value)
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT32", "10!@")
		defer os.Unsetenv("TEST_INVALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_INVALID_FLOAT32", 5.0)

		assert.Equal(t, float32(5.0), value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetFloat32WithDefault("TEST_INVALID_FLOAT32", 5.0)

		assert.Equal(t, float32(5.0), value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT32", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT32")

		value := GetFloat32WithDefault("TEST_INVALID_ENV_KEY", 5.0)

		assert.Equal(t, float32(5.0), value)
	})
}

func TestMustGetFloat64(t *testing.T) {

	t.Run("valid: valid input", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "10")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := MustGetFloat64("TEST_VALID_FLOAT64")

		assert.Equal(t, float64(10.0), value)
	})

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "007")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := MustGetFloat64("TEST_VALID_FLOAT64")

		assert.Equal(t, float64(7.0), value)
	})

	t.Run("valid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "-7")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := MustGetFloat64("TEST_VALID_FLOAT64")

		assert.Equal(t, float64(-7.0), value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64))
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := MustGetFloat64("TEST_VALID_FLOAT64")

		assert.Equal(t, float64(math.MaxFloat64), value)
	})

	t.Run("invalid: value is nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		assert.Panics(t, func() {
			MustGetFloat64("TEST_INVALID_FLOAT64")
		})
	})

	t.Run("invalid: value is math.nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "math.NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		assert.Panics(t, func() {
			MustGetFloat64("TEST_INVALID_FLOAT64")
		})
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "not_a_float64")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		assert.Panics(t, func() {
			MustGetFloat64("TEST_INVALID_FLOAT64")
		})
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "10!@")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		assert.Panics(t, func() {
			MustGetFloat64("TEST_INVALID_FLOAT64")
		})
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

	t.Run("valid: integer value with leading zeros", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "007")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_VALID_FLOAT64", 5.0)

		assert.Equal(t, float64(7.0), value)
	})

	t.Run("valid: negative integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", "-7")
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_VALID_FLOAT64", 5.0)

		assert.Equal(t, float64(-7.0), value)
	})

	t.Run("valid: large integer value", func(t *testing.T) {
		os.Setenv("TEST_VALID_FLOAT64", strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64))
		defer os.Unsetenv("TEST_VALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_VALID_FLOAT64", 5.0)

		assert.Equal(t, float64(math.MaxFloat64), value)
	})

	t.Run("invalid: value is nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_INVALID_FLOAT64", 5.0)

		assert.Equal(t, float64(5.0), value)
	})

	t.Run("invalid: value is math.nan", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "math.NaN")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_INVALID_FLOAT64", 5.0)

		assert.Equal(t, float64(5.0), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "not_a_float64")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_INVALID_FLOAT64", 5.0)

		assert.Equal(t, float64(5.0), value)
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_FLOAT64", "10!@")
		defer os.Unsetenv("TEST_INVALID_FLOAT64")

		value := GetFloat64WithDefault("TEST_INVALID_FLOAT64", 5.0)

		assert.Equal(t, float64(5.0), value)
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
	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://www.google.com/")
		defer os.Unsetenv("TEST_VALID_URL")

		value := MustGetUrl("TEST_VALID_URL")

		assert.Equal(t, "https://www.google.com/", value)
	})

	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://xkcd.com/2293/")
		defer os.Unsetenv("TEST_VALID_URL")

		value := MustGetUrl("TEST_VALID_URL")

		assert.Equal(t, "https://xkcd.com/2293/", value)
	})

	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://github.com/orgs/4rchr4y/repositories?type=source")
		defer os.Unsetenv("TEST_VALID_URL")

		value := MustGetUrl("TEST_VALID_URL")

		assert.Equal(t, "https://github.com/orgs/4rchr4y/repositories?type=source", value)
	})

	t.Run("invalid: empty string as a valid value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "")
		defer os.Unsetenv("TEST_INVALID_URL")

		assert.Panics(t, func() {
			MustGetUrl("TEST_INVALID_URL")
		})
	})

	t.Run("invalid: special characters in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "value@123!")
		defer os.Unsetenv("TEST_INVALID_URL")

		assert.Panics(t, func() {
			MustGetUrl("TEST_INVALID_URL")
		})
	})

	t.Run("invalid: numeric value as string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "12345")
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

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "test string")
		defer os.Unsetenv("TEST_INVALID_URL")

		assert.Panics(t, func() {
			MustGetUrl("TEST_WRONG_KEY")
		})
	})
}

func TestGetUrlWithDefault(t *testing.T) {
	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://www.google.com/")
		defer os.Unsetenv("TEST_VALID_URL")

		value := GetUrlWithDefault("TEST_VALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://www.google.com/", value)
	})

	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://xkcd.com/2293/")
		defer os.Unsetenv("TEST_VALID_URL")

		value := GetUrlWithDefault("TEST_VALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://xkcd.com/2293/", value)
	})

	t.Run("valid: valid env variable", func(t *testing.T) {
		os.Setenv("TEST_VALID_URL", "https://github.com/orgs/4rchr4y/repositories?type=source")
		defer os.Unsetenv("TEST_VALID_URL")

		value := GetUrlWithDefault("TEST_VALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://github.com/orgs/4rchr4y/repositories?type=source", value)
	})

	t.Run("invalid: empty string as a valid value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "")
		defer os.Unsetenv("TEST_INVALID_URL")

		value := GetUrlWithDefault("TEST_INVALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://www.google.com/", value)
	})

	t.Run("invalid: special characters in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "value@123!")
		defer os.Unsetenv("TEST_INVALID_URL")

		value := GetUrlWithDefault("TEST_INVALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://www.google.com/", value)
	})

	t.Run("invalid: numeric value as string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "12345")
		defer os.Unsetenv("TEST_INVALID_URL")

		value := GetUrlWithDefault("TEST_INVALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://www.google.com/", value)
	})

	t.Run("invalid: env key is not set", func(t *testing.T) {
		value := GetUrlWithDefault("TEST_INVALID_URL", "https://www.google.com/")

		assert.Equal(t, "https://www.google.com/", value)
	})

	t.Run("invalid: invalid env key", func(t *testing.T) {
		os.Setenv("TEST_INVALID_URL", "test string")
		defer os.Unsetenv("TEST_INVALID_URL")

		value := GetUrlWithDefault("TEST_INVALID_ENV", "https://www.google.com/")

		assert.Equal(t, "https://www.google.com/", value)
	})
}

func TestMustGetDuration(t *testing.T) {
	t.Run("valid: second input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("valid: minute input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1m")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Minute), value)
	})

	t.Run("valid: hour input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Hour), value)
	})

	t.Run("valid: millisecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1ms")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Millisecond), value)
	})
	t.Run("valid: microsecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1µs")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Microsecond), value)
	})

	t.Run("valid: nanosecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Nanosecond), value)
	})

	t.Run("valid: 4.000000001s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "4.000000001s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(4*time.Second+time.Nanosecond), value)
	})

	t.Run("valid: 1h0m4.000000001s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h0m4.000000001s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Hour+4*time.Second+time.Nanosecond), value)
	})

	t.Run("valid: 1h1m0.01s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h1m0.01s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(61*time.Minute+10*time.Millisecond), value)
	})

	t.Run("valid: 1h1m0.123456789s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h1m0.123456789s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(61*time.Minute+123456789*time.Nanosecond), value)
	})

	t.Run("valid: 1.00002ms input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1.00002ms")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Millisecond+20*time.Nanosecond), value)
	})

	t.Run("valid: 1.00000002s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1.00000002s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Second+20*time.Nanosecond), value)
	})
	t.Run("valid: 693 nanoseconds input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "693ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(693*time.Nanosecond), value)
	})
	t.Run("valid: 10s1us693ns input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "10s1us693ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(10*time.Second+time.Microsecond+693*time.Nanosecond), value)
	})

	t.Run("valid: second input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("valid: negative second input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(-time.Second), value)
	})

	t.Run("valid: negative minute input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1m")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(-time.Minute), value)
	})

	t.Run("valid: negative hour input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1h")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(-time.Hour), value)
	})

	t.Run("valid: negative millisecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1ms")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(-time.Millisecond), value)
	})
	t.Run("valid: negative microsecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1µs")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(-time.Microsecond), value)
	})

	t.Run("valid: negative nanosecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := MustGetDuration("TEST_VALID_TIME")

		assert.Equal(t, time.Duration(-time.Nanosecond), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_TIME")

		assert.Panics(t, func() {
			MustGetDuration("TEST_INVALID_TIME")
		})
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "10s!@")
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
	t.Run("valid: second input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("valid: minute input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1m")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Minute), value)
	})

	t.Run("valid: hour input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Hour), value)
	})

	t.Run("valid: millisecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1ms")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Millisecond), value)
	})
	t.Run("valid: microsecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1µs")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Microsecond), value)
	})

	t.Run("valid: nanosecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Nanosecond), value)
	})

	t.Run("valid: 4.000000001s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "4.000000001s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(4*time.Second+time.Nanosecond), value)
	})

	t.Run("valid: 1h0m4.000000001s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h0m4.000000001s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Hour+4*time.Second+time.Nanosecond), value)
	})

	t.Run("valid: 1h1m0.01s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h1m0.01s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(61*time.Minute+10*time.Millisecond), value)
	})

	t.Run("valid: 1h1m0.123456789s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1h1m0.123456789s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(61*time.Minute+123456789*time.Nanosecond), value)
	})

	t.Run("valid: 1.00002ms input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1.00002ms")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Millisecond+20*time.Nanosecond), value)
	})

	t.Run("valid: 1.00000002s input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1.00000002s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Second+20*time.Nanosecond), value)
	})
	t.Run("valid: 693 nanoseconds input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "693ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(693*time.Nanosecond), value)
	})
	t.Run("valid: 10s1us693ns input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "10s1us693ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(10*time.Second+time.Microsecond+693*time.Nanosecond), value)
	})

	t.Run("valid: second input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "1s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("valid: negative second input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1s")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(-time.Second), value)
	})

	t.Run("valid: negative minute input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1m")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(-time.Minute), value)
	})

	t.Run("valid: negative hour input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1h")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(-time.Hour), value)
	})

	t.Run("valid: negative millisecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1ms")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(-time.Millisecond), value)
	})
	t.Run("valid: negative microsecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1µs")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(-time.Microsecond), value)
	})

	t.Run("valid: negative nanosecond input", func(t *testing.T) {
		os.Setenv("TEST_VALID_TIME", "-1ns")
		defer os.Unsetenv("TEST_VALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(-time.Nanosecond), value)
	})

	t.Run("invalid: value is string", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "not_an_int")
		defer os.Unsetenv("TEST_INVALID_TIME")

		value := GetDurationWithDefault("TEST_VALID_TIME", time.Second)

		assert.Equal(t, time.Duration(time.Second), value)
	})

	t.Run("invalid: special char in value", func(t *testing.T) {
		os.Setenv("TEST_INVALID_TIME", "10s!@")
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
