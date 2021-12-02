package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type pos struct {
	Hor   int
	Depth int
	Aim   int
}

func (p *pos) times(x int) {
	// p.Hor *= x
	// p.Depth *= x
	// p.Aim *= x

	v := reflect.Indirect(reflect.ValueOf(p))
	for i := 0; i < v.NumField(); i++ {
		v.Field(i).SetInt(v.Field(i).Int() * int64(x))
	}
}

func (p *pos) add(p2 pos) {
	// p.Hor += p2.Hor
	// p.Depth += p2.Depth
	// p.Aim += p2.Aim

	v := reflect.Indirect(reflect.ValueOf(p))
	v2 := reflect.ValueOf(p2)
	for i := 0; i < v.NumField(); i++ {
		v.Field(i).SetInt(v.Field(i).Int() + v2.Field(i).Int())
	}
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
			aim := pos{Depth: aimFactor}
			aim.times(amplitude * p.Aim)
			p.add(aim)
		}
	}

	return p.Hor * p.Depth
}

func part1() int {
	dirs := map[string]pos{
		"forward": {Hor: 1},
		"up":      {Depth: -1},
		"down":    {Depth: 1},
	}

	return partx(dirs, 0)
}

func part2() int {
	dirs := map[string]pos{
		"forward": {Hor: 1},
		"up":      {Aim: -1},
		"down":    {Aim: 1},
	}

	return partx(dirs, 1)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func init() {

}
