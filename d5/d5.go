package main

import (
	"fmt"
	"regexp"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type input struct {
	lines []line
	count map[pair]int
}

type line struct {
	a, b pair
}

type pair struct {
	x, y int
}

func (l *line) isHz() bool {
	return l.a.y == l.b.y
}

func (l *line) isVr() bool {
	return l.a.x == l.b.x
}

func sign(a, b int) int {
	if a < b {
		return 1
	}
	return -1
}

func (l *line) count(count map[pair]int) {
	dx, dy := 0, 0
	if l.isHz() {
		dx = 1 * sign(l.a.x, l.b.x)
	} else if l.isVr() {
		dy = 1 * sign(l.a.y, l.b.y)
	} else {
		dx = 1 * sign(l.a.x, l.b.x)
		dy = 1 * sign(l.a.y, l.b.y)
	}

	x, y := 0, 0
	for x, y = l.a.x, l.a.y; x != l.b.x || y != l.b.y; x, y = x+dx, y+dy {
		count[pair{x, y}] += 1
	}
	count[pair{x, y}] += 1
}

func (in *input) doCount(part2 bool) int {
	for _, l := range in.lines {
		if part2 || l.isHz() || l.isVr() {
			l.count(in.count)
		}
	}

	count := 0

	for _, c := range in.count {
		if c > 1 {
			count++
		}
	}

	return count
}

func newInput() *input {
	ret := input{count: make(map[pair]int)}
	re := regexp.MustCompile(`^([0-9]*),([0-9]*) -> ([0-9]*),([0-9]*)$`)

	for s := range tools.ReadInput("input.txt") {
		tab := re.FindStringSubmatch(s)
		a := pair{tools.S2i(tab[1]), tools.S2i(tab[2])}
		b := pair{tools.S2i(tab[3]), tools.S2i(tab[4])}
		ret.lines = append(ret.lines, line{a, b})
	}

	return &ret
}

func part1() int {
	return newInput().doCount(false)
}

func part2() int {
	return newInput().doCount(true)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
