package main

import (
	"fmt"
	"math"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

func part1() int {
	prec := math.MaxInt32
	inc := 0

	for line := range tools.ReadInput("input.txt") {
		cur := tools.S2i(line)
		if cur > prec {
			inc += 1
		}
		prec = cur
	}

	return inc
}

func part2() int {
	inc := 0

	lines := tools.ReadInput("input.txt")
	a := tools.S2i(<-lines)
	b := tools.S2i(<-lines)
	c := tools.S2i(<-lines)

	for line := range lines {
		d := tools.S2i(line)
		if a+b+c < b+c+d {
			inc += 1
		}
		a = b
		b = c
		c = d
	}

	return inc
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
