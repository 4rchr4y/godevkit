package env

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/4rchr4y/godevkit/v2/must"
	"github.com/4rchr4y/godevkit/v2/regex"
)

func MustGetString(key string) string {
	value, ok := os.LookupEnv(key)
	return must.MustBeOk(value, ok, fmt.Errorf("environment variable '%s' was not found", key))
}

func GetStringWithDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}

func MustGetInt(key string) int {
	value, err := strconv.Atoi(MustGetString(key))
	if err != nil {
		err = fmt.Errorf("invalid environment variable '%s' value: %v", key, err)
	}

	return must.Must(value, err)
}

func GetIntWithDefault(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}

	return value
}

func MustGetUint(key string) uint64 {
	value, err := strconv.ParseUint(MustGetString(key), 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid environment variable '%s' value: %v", key, err)
	}

	return must.Must(value, err)
}

func GetUintWithDefault(key string, defaultValue uint) uint {
	value, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return defaultValue
	}

	return uint(value)
}

func MustGetFloat64(key string) float64 {
	value, err := strconv.ParseFloat(MustGetString(key), 64)
	if err != nil {
		err = fmt.Errorf("invalid environment variable '%s' value: %v", key, err)
	}

	return must.Must(value, err)
}

func GetFloat64WithDefault(key string, defaultValue float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return defaultValue
	}

	return value
}

func MustGetDuration(key string) time.Duration {
	value, err := time.ParseDuration(MustGetString(key))
	if err != nil {
		err = fmt.Errorf("invalid environment variable '%s' value: %v", key, err)
	}

	return must.Must(value, err)
}

func GetDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return defaultValue
	}

	return value
}

func MustGetUrl(key string) string {
	value := MustGetString(key)
	matched, err := regexp.MatchString(regex.UrlPatternString, value)
	value = must.Must(value, err)

	if !matched {
		panic("env " + key + " with value '" + value + "' is not matching url pattern")
	}

	return value
}

func GetUrlWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	matched, err := regexp.MatchString(regex.UrlPatternString, value)
	if err != nil || !matched {
		return defaultValue
	}

	return value
}
