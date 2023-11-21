package must

import "strings"

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}

	return obj
}

func MustBeOk[T any](value T, ok bool, args ...error) T {
	return assertCondition(value, ok, "value must be ok", args...)
}

func MustNotBeOk[T any](value T, ok bool, args ...error) T {
	return assertCondition(value, ok, "value must not be ok", args...)
}

func assertCondition[T any](value T, condition bool, defaultMsg string, args ...error) T {
	if !condition {
		panic(buildErrorMessage(defaultMsg, args...))
	}

	return value
}

func buildErrorMessage(defaultErrorMessage string, args ...error) string {
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
