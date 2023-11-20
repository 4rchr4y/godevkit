package must

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}

	return obj
}

func BeOk[T any](value T, ok bool) T {
	if !ok {
		panic("must be ok")
	}
	return value
}

func NotBeOk[T any](value T, ok bool) T {
	if ok {
		panic("must not be ok")
	}
	return value
}
