package main

import (
	"fmt"
	"strings"

	"github.com/bhesmans/aoc-2021/internal/tools"
)

type signal = string

type display struct {
	patterns []signal
	values   []signal

	mapping map[byte]byte
}

type input struct {
	displays []*display
}

var sigToDec map[int]int

func sigToInt(s signal) int {
	ret := 0
	for _, c := range s {
		ret += 1 << (c - 'a')
	}

	return ret
}

func minus(s1, s2 signal) signal {
	ret := ""
	for _, c := range s1 {
		if !strings.ContainsAny(s2, string(c)) {
			ret += string(c)
		}
	}
	return ret
}

func contains(s1, s2 signal) bool {
	for _, c := range s2 {
		if !strings.ContainsAny(s1, string(c)) {
			return false
		}
	}

	return true
}

func initSigToDec() {
	sigToDec = make(map[int]int)
	sigToDec[sigToInt("abcefg")] = 0
	sigToDec[sigToInt("cf")] = 1
	sigToDec[sigToInt("acdeg")] = 2
	sigToDec[sigToInt("acdfg")] = 3
	sigToDec[sigToInt("bcdf")] = 4
	sigToDec[sigToInt("abdfg")] = 5
	sigToDec[sigToInt("abdefg")] = 6
	sigToDec[sigToInt("acf")] = 7
	sigToDec[sigToInt("abcdefg")] = 8
	sigToDec[sigToInt("abcdfg")] = 9

}

func (d *display) doMapping() {
	// a
	a := minus(d.find7(), d.find1())[0]
	d.mapping[a] = 'a'
	// g
	g := minus(minus(d.find3(), d.find4()), signal(a))[0]
	d.mapping[g] = 'g'
	// e
	e := minus(minus(minus(d.find8(), d.find4()), signal(a)), signal(g))[0]
	d.mapping[e] = 'e'
	// d
	dd := minus(minus(minus(d.find2(), d.find7()), signal(e)), signal(g))[0]
	d.mapping[dd] = 'd'
	// b
	b := minus(minus(d.find4(), d.find1()), signal(dd))[0]
	d.mapping[b] = 'b'
	// c
	c := minus(minus(minus(minus(d.find2(), signal(dd)), signal(g)), signal(e)), signal(a))[0]
	d.mapping[c] = 'c'
	// f
	f := minus(minus(minus(minus(minus(minus(d.find8(), signal(a)), signal(b)), signal(c)), signal(dd)), signal(e)), signal(g))[0]
	d.mapping[f] = 'f'

	return
}

func newInput() *input {
	ret := input{}

	for l := range tools.ReadInput("input.txt") {
		tab := strings.Split(l, " | ")
		d := &display{
			patterns: strings.Split(tab[0], " "),
			values:   strings.Split(tab[1], " "),
			mapping:  make(map[byte]byte),
		}

		d.doMapping()

		ret.displays = append(ret.displays, d)
	}

	return &ret
}

func is1478(s signal) bool {
	l := len(s)
	return l == 2 || l == 4 || l == 3 || l == 7
}

func (d *display) findSignalLen(l int) signal {
	for i := range d.patterns {
		if len(d.patterns[i]) == l {
			return d.patterns[i]
		}
	}
	panic("haaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}

func (d *display) find7() signal {
	return d.findSignalLen(3)
}

func (d *display) find1() signal {
	return d.findSignalLen(2)
}

func (d *display) find4() signal {
	return d.findSignalLen(4)
}

func (d *display) find8() signal {
	return d.findSignalLen(7)
}

func (d *display) find3() signal {
	one := d.find1()
	for _, p := range d.patterns {
		if len(p) == 5 && contains(p, one) {
			return p
		}
	}
	panic("meh")
}

func (d *display) find2() signal {
	e := ""
	for k, v := range d.mapping {
		if v == 'e' {
			e = string(k)
		}
	}

	if e == "" {
		panic("fichtre")
	}

	for _, p := range d.patterns {
		if len(p) == 5 && contains(p, e) {
			return p
		}
	}
	panic("meh")
}

func (in *input) count1478() int {
	count := 0

	for _, d := range in.displays {
		for i := range d.values {
			if is1478(d.values[i]) {
				count += 1
			}
		}
	}

	return count
}

func part1() int {
	in := newInput()
	return in.count1478()
}

func (d *display) decodeSignal(s signal) signal {
	decodedSignal := ""

	for _, c := range s {
		decodedSignal += string(d.mapping[byte(c)])
	}

	return decodedSignal
}

func (d *display) decode() int {
	ret := 0
	for _, s := range d.values {
		ret *= 10
		ret += sigToDec[sigToInt(d.decodeSignal(s))]
	}
	return ret
}

func part2() int {
	in := newInput()

	sum := 0
	for _, d := range in.displays {
		sum += d.decode()
	}
	return sum
}

func main() {
	initSigToDec()

	fmt.Println(part1())
	fmt.Println(part2())
}
