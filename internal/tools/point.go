package tools

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
	LeftUp
	RightUp
	LeftDown
	RightDown
)

var (
	Move map[Direction]Point = map[Direction]Point{
		Left:  {-1, 0},
		Right: {1, 0},
		Up:    {0, -1},
		Down:  {0, 1},
	}
	MoveDiag map[Direction]Point = map[Direction]Point{
		LeftUp:    {-1, -1},
		RightUp:   {1, -1},
		LeftDown:  {-1, 1},
		RightDown: {1, 1},
	}
	MoveAll map[Direction]Point = map[Direction]Point{} //check init
)

type Point struct {
	X, Y int
}

func (p1 *Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p *Point) NeighborsDir(dirs map[Direction]Point) <-chan Point {
	c := make(chan Point)
	go func() {
		for _, m := range dirs {
			n := p.Add(m)
			c <- n
		}
		close(c)
	}()
	return c
}

func (p *Point) Neighbors() <-chan Point {
	return p.NeighborsDir(Move)
}

func (p *Point) NeighborsDiag() <-chan Point {
	return p.NeighborsDir(MoveDiag)
}

func (p *Point) NeighborsAll() <-chan Point {
	return p.NeighborsDir(MoveAll)
}

func init() {
	for k, v := range Move {
		MoveAll[k] = v
	}
	for k, v := range MoveDiag {
		MoveAll[k] = v
	}
}
