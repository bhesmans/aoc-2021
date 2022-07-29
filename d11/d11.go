package main

import (
	"fmt"

	. "github.com/bhesmans/aoc-2021/internal/tools"
)

type octopus struct {
	energy int
	point  Point
}

type input struct {
	octopuses  map[Point]*octopus
	mustFlash  []Point
	flashCount int
}

func newInput() *input {
	ret := input{
		octopuses: make(map[Point]*octopus),
	}

	y := 0
	// for l := range ReadInput("small_input.txt") {
	for l := range ReadInput("input.txt") {
		for x, v := range l {
			ret.octopuses[Point{X: x, Y: y}] = &octopus{
				energy: S2i(string(v)),
				point:  Point{X: x, Y: y},
			}
		}
		y += 1
	}

	return &ret
}

func (o *octopus) mustFlash() bool {
	// Strictly 10. if higher it has already flashed and must not flash twice
	return o.energy == 10
}

func (o *octopus) hasFlashed() bool {
	return o.energy > 9
}

func (in *input) enqueueFlash(o *octopus) {
	if o.mustFlash() {
		in.mustFlash = append(in.mustFlash, o.point)
	}
}

func (in *input) increaseEnergyOctopus(o *octopus) {
	o.energy += 1
	in.enqueueFlash(o)
}

func (in *input) increaseEnergy() {
	for _, o := range in.octopuses {
		in.increaseEnergyOctopus(o)
	}
}

func (in *input) resetEnergy() bool {
	flashed := 0
	for _, o := range in.octopuses {
		if o.hasFlashed() {
			flashed += 1
			o.energy = 0
		}
	}

	return flashed == len(in.octopuses)
}

func (in *input) flashThemAll() {
	for len(in.mustFlash) != 0 {
		for _, p := range in.mustFlash {
			in.flashCount += 1
			in.flash(p)
			in.mustFlash = in.mustFlash[1:]
		}
	}
}

func (in *input) flash(p Point) {
	for n := range p.NeighborsAll() {
		if o, ok := in.octopuses[n]; ok {
			in.increaseEnergyOctopus(o)
		}
	}
}

// return true if all the octopuses flash together in this step
func (in *input) step() bool {
	in.increaseEnergy()
	in.flashThemAll()
	return in.resetEnergy()
}

func (in *input) steps(n int) {
	for i := 0; i < n; i++ {
		_ = in.step()
	}
}

func (in *input) String() string {
	return fmt.Sprintf("%v", in.octopuses)
}

func (o *octopus) String() string {
	return fmt.Sprintf("%v", o.energy)
}

func part1() int {
	in := newInput()
	in.steps(100)
	return in.flashCount
}

func part2() int {
	in := newInput()
	iter := 1
	for !in.step() {
		iter += 1
	}
	return iter
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
