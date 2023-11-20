package must

import "strings"

const (
	errMustBeOkString = "must be ok"
)

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}

	return obj
}

func MustBeOk[T any](value T, ok bool, args ...string) T {
	if !ok {
		panic(errMustBeOkString + " " + strings.Join(args, " "))
	}
	return value
}

func MustNotBeOk[T any](value T, ok bool) T {
	if ok {
		panic("must not be ok")
	}
	return value
}
