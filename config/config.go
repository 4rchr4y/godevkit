package config

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/4rchr4y/gdk/must"
	"github.com/4rchr4y/gdk/regex"
)

func requireEnvVar(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New("") // TODO: set err message
	}

	return value, nil
}

func MustGetInt(key string) int         { return must.Must(strconv.Atoi(os.Getenv(key))) }
func MustGetUint(key string) uint64     { return must.Must(strconv.ParseUint(os.Getenv(key), 10, 64)) }
func MustGetFloat64(key string) float64 { return must.Must(strconv.ParseFloat(os.Getenv(key), 64)) }
func MustGetString(key string) string   { return must.Must(requireEnvVar(key)) }

// func MustGetString(key string) string {
// 	value, ok := os.LookupEnv(key)
// 	if !ok {
// 		err := fmt.Errorf(""environment variable '%s' not found", key)
// 		panic(err)
// 	}

// 	return value
// }

func GetStringWithDefault(key string, defaultValue string) string {
	// TODO: use MustBeOk
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
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

func GetFloat32WithDefault(key string, defaultValue float32) float32 {
	value, err := strconv.ParseFloat(os.Getenv(key), 32)
	if err != nil {
		return defaultValue
	}

	return float32(value)
}

func GetFloat64WithDefault(key string, defaultValue float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return defaultValue
	}

	return value
}

func MustGetUrl(key string) string {
	value := os.Getenv(key)
	matched, err := regexp.MatchString(regex.UrlPatternString, value)
	if !matched {
		err := fmt.Errorf("environment variable '%s' by key '%s' is not matching URL pattern", value, key)
		panic(err)
	}

	must.Must(value, err)

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

func MustGetDuration(key string) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	must.Must(value, err)

	return value
}

func GetDurationWithDefault(key string, defaultValue time.Duration) time.Duration {
	value, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return defaultValue
	}

	return value
}
