package main

import (
	"fmt"
	"sort"

	. "github.com/bhesmans/aoc-2021/internal/tools"
)

type location struct {
	height int
}

type input struct {
	hmap map[Point]location
}

type basin struct {
	points map[Point]bool
	fringe []Point

	in *input
}

func (b *basin) expandOnce() bool {
	if len(b.fringe) == 0 {
		return false
	}

	current := b.fringe[0]
	b.fringe = b.fringe[1:]

	for n := range current.Neighbors() {
		h, ok := b.in.height(n)
		if !ok {
			continue
		}

		if !b.points[n] && h != 9 {
			b.points[n] = true
			b.fringe = append(b.fringe, n)
		}
	}

	return true
}

func (b *basin) expand() {
	for b.expandOnce() {
	}
}

func newInput() *input {
	ret := input{hmap: make(map[Point]location)}

	y := 0
	for l := range ReadInput("input.txt") {
		for x, v := range l {
			ret.hmap[Point{X: x, Y: y}] = location{height: S2i(string(v))}
		}
		y += 1
	}

	return &ret
}

func (in *input) height(p Point) (int, bool) {
	if l, ok := in.hmap[p]; ok {
		return l.height, true
	}
	return 0, false
}

func (in *input) isLowPoint(p Point) bool {
	hp, _ := in.height(p)
	for n := range p.Neighbors() {
		if hs, ok := in.height(n); ok && hp >= hs {
			return false
		}
	}
	return true
}

func part1() int {
	in := newInput()
	ret := 0
	for p := range in.hmap {
		if in.isLowPoint(p) {
			h, _ := in.height(p)
			ret += (h + 1)
		}
	}
	return ret
}

func (in *input) newBasin(p Point) *basin {
	b := &basin{
		fringe: []Point{p},
		points: map[Point]bool{p: true},
		in:     in,
	}

	b.expand()

	return b
}

func part2() int {
	in := newInput()
	sizes := []int{}

	for p := range in.hmap {
		if in.isLowPoint(p) {
			sizes = append(sizes, len(in.newBasin(p).points))
		}
	}

	sort.Ints(sizes)

	last := len(sizes)
	return sizes[last-1] * sizes[last-2] * sizes[last-3]
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
