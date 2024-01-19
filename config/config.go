package config

import (
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/4rchr4y/godevkit/must"
	"github.com/4rchr4y/godevkit/regex"
)

func MustGetInt(key string) int                { return must.Must(strconv.Atoi(os.Getenv(key))) }
func MustGetUint(key string) uint64            { return must.Must(strconv.ParseUint(os.Getenv(key), 10, 64)) }
func MustGetFloat64(key string) float64        { return must.Must(strconv.ParseFloat(os.Getenv(key), 64)) }
func MustGetString(key string) string          { return must.MustBeOk(os.LookupEnv(key)) }
func MustGetDuration(key string) time.Duration { return must.Must(time.ParseDuration(os.Getenv(key))) }
func MustGetUrl(key string) string {
	value := must.MustBeOk[string](os.LookupEnv(key))
	matched, err := regexp.MatchString(regex.UrlPatternString, value)
	value = must.Must[string](value, err)

	if !matched {
		panic("env " + key + " with value '" + value + "' is not matching url pattern")
	}

	return value
}

func GetIntWithDefault(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}

	return value
}

func GetUintWithDefault(key string, defaultValue uint) uint {
	value, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return defaultValue
	}

	return uint(value)
}

func GetFloat64WithDefault(key string, defaultValue float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return defaultValue
	}

	return value
}

func GetStringWithDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}
func GetDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return defaultValue
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
