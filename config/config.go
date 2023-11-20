package config

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/4rchr4y/gdk/must"
)

func MustGetString(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		err := fmt.Errorf("environment variable '%s' not found", key)
		panic(err)
	}

	return value
}

func GetStringWithDefault(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func MustGetInt(key string) int {
	value := must.Must(strconv.Atoi(os.Getenv(key)))

	return value
}

func GetIntWithDefault(key string, fallback int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}

	return value
}

func MustGetUint(key string) uint {
	value := must.Must(strconv.ParseUint(os.Getenv(key), 10, 64))

	return uint(value)
}

func GetUintWithDefault(key string, fallback uint) uint {
	value, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return fallback
	}

	return uint(value)
}

func MustGetFloat32(key string) float32 {
	value := must.Must(strconv.ParseFloat(os.Getenv(key), 32))
	if math.IsNaN(value) {
		err := fmt.Errorf("env value by key '%s' should not be NaN", key)
		panic(err)
	}

	return float32(value)
}

func GetFloat32WithDefault(key string, fallback float32) float32 {
	value, err := strconv.ParseFloat(os.Getenv(key), 32)
	if err != nil || math.IsNaN(value) {
		return fallback
	}

	return float32(value)
}

func MustGetFloat64(key string) float64 {
	value := must.Must(strconv.ParseFloat(os.Getenv(key), 64))
	if math.IsNaN(value) {
		err := fmt.Errorf("env value by key '%s' should not be NaN", key)
		panic(err)
	}

	return value
}

func GetFloat64WithDefault(key string, fallback float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil || math.IsNaN(value) {
		return fallback
	}

	return value
}

func MustGetUrl(key string) string {
	pattern := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	value := os.Getenv(key)
	matched, err := regexp.MatchString(pattern, value)
	if !matched {
		err := fmt.Errorf("environment variable '%s' by key '%s' is not matching URL pattern", value, key)
		panic(err)
	}

	must.Must(value, err)

	return value
}

func GetUrlWithDefault(key string, fallback string) string {
	pattern := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

	value := os.Getenv(key)
	matched, err := regexp.MatchString(pattern, value)
	if err != nil || !matched {
		return fallback
	}

	return value
}

func MustGetDuration(key string) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))

	must.Must(value, err)

	return value
}

func GetDurationWithDefault(key string, fallback time.Duration) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return fallback
	}

	return value
}
