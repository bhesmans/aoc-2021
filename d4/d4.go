package main

import (
	"fmt"
	"strings"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type pair struct {
	x, y int
}

type square struct {
	value  int
	marked bool
}

type board struct {
	squares map[pair]*square
	values  map[int]pair
}

type input struct {
	draw   []int
	boards []*board
}

func (b *board) mark(value int) (bool, pair) {
	if p, ok := b.values[value]; ok {
		b.squares[p].marked = true
		return true, p
	}

	return false, pair{}
}

func (b *board) checkX(x int) bool {
	for y := 0; y < 5; y++ {
		if !b.squares[pair{x, y}].marked {
			return false
		}
	}
	return true
}

func (b *board) checkY(y int) bool {
	for x := 0; x < 5; x++ {
		if !b.squares[pair{x, y}].marked {
			return false
		}
	}
	return true
}

func (b *board) check(p pair) bool {
	return b.checkX(p.x) || b.checkY(p.y)
}

func (in *input) winningBoard() (*board, int, int) {
	for _, v := range in.draw {
		for idx, b := range in.boards {
			if ok, p := b.mark(v); ok {
				if b.check(p) {
					return b, v, idx
				}
			}
		}
	}

	panic("haaaaaaaaaaaaaa")
}

func (b *board) sum() int {
	sum := 0
	for _, sq := range b.squares {
		if !sq.marked {
			sum += sq.value
		}
	}

	return sum
}

func newInput() *input {
	ret := input{}
	lines := tools.ReadInput("input.txt")

	d := <-lines
	for _, i := range strings.Split(d, ",") {
		ret.draw = append(ret.draw, tools.S2i(i))
	}

	y := 0
	b := &board{}
	for line := range lines {
		if line == "" {
			b = &board{squares: make(map[pair]*square), values: make(map[int]pair)}
			ret.boards = append(ret.boards, b)
			y = 0
			continue
		}

		for x, v := range strings.Fields(line) {
			b.squares[pair{x, y}] = &square{value: tools.S2i(v)}
			b.values[tools.S2i(v)] = pair{x, y}
		}

		y += 1
	}

	return &ret
}

func part1() int {
	b, v, _ := newInput().winningBoard()
	return b.sum() * v
}

func part2() int {
	in := newInput()

	for len(in.boards) != 1 {
		_, _, idx := in.winningBoard()
		in.boards = append(in.boards[:idx], in.boards[idx+1:]...)
	}

	b, v, _ := in.winningBoard()

	return b.sum() * v
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
