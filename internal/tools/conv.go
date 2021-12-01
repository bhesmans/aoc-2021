package tools

import "strconv"

func S2i(s string) int {
	i, err := strconv.Atoi(s)
	PanicOnError(err)
	return i
}
