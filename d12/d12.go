package main

import (
	"fmt"
	"strings"

	. "github.com/bhesmans/aoc-2021/internal/tools"
)

type cave struct {
	id string
}

type input struct {
	paths map[cave][]cave
}

type path struct {
	visited     map[cave]bool
	caves       []cave
	doubleVisit bool
}

func newPath() path {
	return path{visited: make(map[cave]bool)}
}

func (c *cave) isSmall() bool {
	return c.id[0] > 96
}

func (c cave) isEnd() bool {
	return c.id == "end"
}

func (c cave) isStart() bool {
	return c.id == "start"
}

func (p *path) clone() path {
	ret := path{visited: make(map[cave]bool), caves: make([]cave, len(p.caves))}
	copy(ret.caves, p.caves)

	for k, v := range p.visited {
		ret.visited[k] = v
	}

	ret.doubleVisit = p.doubleVisit

	return ret
}

func (p path) expand(c cave) path {
	ret := p.clone()

	if c.isSmall() && p.visited[c] {
		ret.doubleVisit = true
	}

	ret.visited[c] = true
	ret.caves = append(ret.caves, c)

	return ret
}

func (p *path) current() cave {
	return p.caves[len(p.caves)-1]
}

func newInput() *input {
	ret := input{
		paths: make(map[cave][]cave),
	}

	// for l := range ReadInput("small_input.txt") {
	for l := range ReadInput("input.txt") {
		tab := strings.Split(l, "-")
		a := cave{id: tab[0]}
		b := cave{id: tab[1]}

		ret.paths[a] = append(ret.paths[a], b)
		ret.paths[b] = append(ret.paths[b], a)
	}

	return &ret
}

func (in *input) expand(p path, part2 bool) []path {
	ret := []path{}
	current := p.current()

	for _, c := range in.paths[current] {
		if c.isStart() {
			continue
		}

		if !part2 && c.isSmall() && p.visited[c] {
			continue
		}

		if part2 && c.isSmall() && p.visited[c] && p.doubleVisit {
			continue
		}

		newPath := p.expand(c)

		ret = append(ret, newPath)

	}

	return ret
}

func (in *input) countPaths(part2 bool) int {
	pathCount := 0
	paths := []path{newPath().expand(cave{id: "start"})}
	for len(paths) != 0 {
		current := paths[0]
		paths = paths[1:]

		if current.current().isEnd() {
			pathCount += 1
			continue
		}

		paths = append(paths, in.expand(current, part2)...)
	}

	return pathCount

}

func (in *input) String() string {
	return fmt.Sprintf("%v", in.paths)
}

func part1() int {
	in := newInput()
	return in.countPaths(false)
}

func part2() int {
	in := newInput()
	return in.countPaths(true)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
