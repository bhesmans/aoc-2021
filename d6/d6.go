package main

import (
	"fmt"
	"strings"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type input struct {
	fishes map[int]map[int]int
}

func (in *input) addFish(qty, day, left int) {
	if _, ok := in.fishes[day]; !ok {
		in.fishes[day] = make(map[int]int)
	}

	in.fishes[day][left] += qty
}

func newInput() *input {
	ret := input{fishes: make(map[int]map[int]int)}

	l := <-tools.ReadInput("input.txt")
	tab := strings.Split(l, ",")
	for _, v := range tab {
		ret.addFish(1, 0, tools.S2i(v))
	}

	return &ret
}

func (in *input) rmDay(day, max int) int {
	ret := 0

	for left, qty := range in.fishes[day] {
		ret += qty
		for i := day + left + 1; i < max; i += 7 {
			in.addFish(qty, i, 8)
		}
	}

	delete(in.fishes, day)

	return ret
}

func rmDays(max int) int {
	count := 0
	in := newInput()

	for i := 0; i < max+1; i++ {
		count += in.rmDay(i, max+1)
	}

	return count
}

func part1() int {
	return rmDays(80)
}

func part2() int {
	return rmDays(256)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
