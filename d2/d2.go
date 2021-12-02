package main

import (
	"fmt"
	"strings"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type pos struct {
	hor   int
	depth int
	aim   int
}

func (p *pos) times(x int) {
	p.hor *= x
	p.depth *= x
	p.aim *= x
}

func (p *pos) add(p2 pos) {
	p.hor += p2.hor
	p.depth += p2.depth
	p.aim += p2.aim
}

func partx(d map[string]pos, aimFactor int) int {
	p := pos{}

	for line := range tools.ReadInput("input.txt") {
		tab := strings.Split(line, " ")
		dir := tab[0]
		amplitude := tools.S2i(tab[1])

		step := d[dir]
		step.times(amplitude)

		p.add(step)

		if dir == "forward" {
			aim := pos{depth: aimFactor}
			aim.times(amplitude * p.aim)
			p.add(aim)
		}
	}

	return p.hor * p.depth
}

func part1() int {
	dirs := map[string]pos{
		"forward": {hor: 1},
		"up":      {depth: -1},
		"down":    {depth: 1},
	}

	return partx(dirs, 0)
}

func part2() int {
	dirs := map[string]pos{
		"forward": {hor: 1},
		"up":      {aim: -1},
		"down":    {aim: 1},
	}

	return partx(dirs, 1)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func init() {

}
