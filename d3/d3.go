package main

import (
	"fmt"
	"strconv"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type input struct {
	lines []string
}

func newInput() *input {
	ret := input{}
	for line := range tools.ReadInput("input.txt") {
		ret.lines = append(ret.lines, line)
	}
	return &ret
}

func (in *input) mostCommon(pos int) byte {
	count := 0
	for _, line := range in.lines {
		if line[pos] == '1' {
			count += 1
		} else {
			count -= 1
		}
	}

	if count >= 0 {
		return '1'
	} else {
		return '0'
	}
}

func (in *input) keep(b byte, pos int) {
	keep := []string{}

	for _, line := range in.lines {
		if line[pos] == b {
			keep = append(keep, line)
		}
	}

	in.lines = keep
}

func (in *input) reduce(pos int, rev bool) {
	mc := in.mostCommon(pos)

	if rev {
		if mc == '1' {
			mc = '0'
		} else {
			mc = '1'
		}
	}

	in.keep(mc, pos)
}

func (in *input) getVal(rev bool) int {
	for i := 0; len(in.lines) != 1; i++ {
		in.reduce(i, rev)
	}

	val, _ := strconv.ParseInt(in.lines[0], 2, 64)
	return int(val)
}

func part1() int {
	ones := map[int]int{}

	for line := range tools.ReadInput("input.txt") {
		for i, b := range line {
			if b == '1' {
				ones[i] += 1
			} else {
				ones[i] -= 1
			}
		}
	}

	gam, eps := 0, 0
	exp := 1

	for i := len(ones) - 1; i >= 0; i-- {
		if ones[i] > 0 {
			gam += exp
		} else {
			eps += exp
		}
		exp *= 2
	}

	return gam * eps
}

func part2() int {
	oxy := newInput().getVal(false)
	co2 := newInput().getVal(true)

	return oxy * co2
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
