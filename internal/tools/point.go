package tools

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

var (
	Move map[Direction]Point = map[Direction]Point{
		Left:  {-1, 0},
		Right: {1, 0},
		Up:    {0, -1},
		Down:  {0, 1},
	}
)

type Point struct {
	X, Y int
}

func (p1 *Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p *Point) Neighbors() <-chan Point {
	c := make(chan Point)
	go func() {
		for _, m := range Move {
			n := p.Add(m)
			c <- n
		}
		close(c)
	}()
	return c
}
