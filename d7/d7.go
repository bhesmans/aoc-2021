package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type input struct {
	crabes   []int
	min, max int
}

type fuelFun = func(int, int) int

func newInput() *input {
	ret := input{}

	l := <-tools.ReadInput("input.txt")
	tab := strings.Split(l, ",")
	for _, v := range tab {
		ret.crabes = append(ret.crabes, tools.S2i(v))
	}

	_, min := tools.Min(ret.crabes)
	_, max := tools.Max(ret.crabes)

	ret.min = min
	ret.max = max

	return &ret
}

func fuel(from, to int) int {
	return tools.Abs(from - to)
}

func fuel2(from, to int) int {
	diff := tools.Abs(from - to)
	return diff * (diff + 1) / 2
}

func (in *input) totalFuel(fuel fuelFun, target int) int {
	ret := 0
	for i := range in.crabes {
		ret += fuel(in.crabes[i], target)
	}

	return ret
}

func (in *input) minFuel(fuel fuelFun) int {
	ret := math.MaxInt
	for i := in.min; i <= in.max; i++ {
		current := in.totalFuel(fuel, i)
		if current < ret {
			ret = current
		}
	}

	return ret
}

func part1() int {
	in := newInput()
	return in.minFuel(fuel)
}

func part2() int {
	in := newInput()
	return in.minFuel(fuel2)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
