package main

import (
	"fmt"
	"sort"

	. "github.com/bhesmans/aoc-2021/internal/tools"
)

type input struct {
	lines []string
}

var (
	opens = map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
		'(': ')',
	}
	scores = map[rune]int{
		'>': 25137,
		'}': 1197,
		']': 57,
		')': 3,
	}
	cscores = map[rune]int{
		'<': 4,
		'{': 3,
		'[': 2,
		'(': 1,
	}
)

func score(l string) (int, []rune) {
	stack := []rune{}

	for _, c := range l {
		if _, ok := opens[c]; ok {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				panic("meh")
			}

			topi := len(stack) - 1
			top := stack[topi]

			if opens[top] != c {
				return scores[c], nil
			}

			stack = stack[:topi]
		}
	}

	return 0, stack
}

func complete(stack []rune) int {
	score := 0
	slen := len(stack)
	for i := range stack {
		c := stack[slen-1-i]
		score *= 5
		score += cscores[c]
	}
	return score
}

func newInput() *input {
	ret := input{}

	for l := range ReadInput("input.txt") {
		ret.lines = append(ret.lines, l)
	}

	return &ret
}

func part1() int {
	in := newInput()
	sum := 0
	for _, l := range in.lines {
		s, _ := score(l)
		sum += s
	}
	return sum
}

func part2() int {
	in := newInput()
	scores := []int{}
	for _, l := range in.lines {
		_, stack := score(l)
		if len(stack) == 0 {
			continue
		}
		scores = append(scores, complete(stack))
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
