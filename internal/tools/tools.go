package tools

import "fmt"

var (
	EmptySlice = fmt.Errorf("Empty Slice")
)

func Min(s []int) (error, int) {
	if len(s) == 0 {
		return EmptySlice, 0
	}

	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}

	return nil, min
}

func Max(s []int) (error, int) {
	if len(s) == 0 {
		return EmptySlice, 0
	}

	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}

	return nil, max
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
