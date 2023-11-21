package must

import "strings"

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}

	return obj
}

func MustBeOk[T any](value T, ok bool, args ...error) T {
	if !ok {
		panic(BuildErrorMessage("must be ok", args))
	}
	return value
}

func MustNotBeOk[T any](value T, ok bool, args ...error) T {
	if ok {
		panic(BuildErrorMessage("must not be ok", args))
	}
	return value
}

func BuildErrorMessage(defaultErrorMessage string, args []error) string {
	if len(args) == 0 {
		return defaultErrorMessage
	}

	var sb strings.Builder
	for i, err := range args {
		sb.WriteString(err.Error())
		if len(args) > 1 && len(args)-1 != i {
			sb.WriteString("; ")
		}
	}

	return sb.String()
}
