package rubik

import (
	"strings"
)

// Representation of the Rubik's Cube.

// Faces colors.
const (
	BLUE   = 'b'
	WHITE  = 'w'
	YELLOW = 'y'
	ORANGE = 'o'
	GREEN  = 'g'
	RED    = 'r'
)

// Moves defined using Singmaster's notation.
// https://en.wikipedia.org/wiki/Rubik%27s_Cube#Move_notation
type Move string

var Moves = []Move{
	UP, UP_COUNTER, DOWN, DOWN_COUNTER, LEFT, LEFT_COUNTER,
	RIGHT, RIGHT_COUNTER, FRONT, FRONT_COUNTER, BACK, BACK_COUNTER,
}

const (
	UP            = "U"
	UP_COUNTER    = "U'"
	DOWN          = "D"
	DOWN_COUNTER  = "D'"
	LEFT          = "L"
	LEFT_COUNTER  = "L'"
	RIGHT         = "R"
	RIGHT_COUNTER = "R'"
	FRONT         = "F"
	FRONT_COUNTER = "F'"
	BACK          = "B"
	BACK_COUNTER  = "B'"
)

// See Cube.pdf for an illustrated version.
// The cube is a []rune where each face is stored as follows:
//
//              #3 BLU
//             35 34 33
//             32 31 30
//             29 28 27
//
//  #4 ORG      #0 WHT     #2 RED
// 42 39 36     0  1  2    20 23 26
// 43 40 37     3  4  5    19 22 25
// 44 41 38     6  7  8    18 21 24
//
//              #1 GRN
//              9 10 11
//             12 13 14
//             15 16 17
//
//              #5 YLW
//             45 46 47
//             48 49 50
//             51 52 53
type Cube []rune

// During a move, a facet is moved `From` `To`.
type Pair struct {
	From, To int
}

// Build a new Cube from the given string representation.
//
// For instance:
// "bwwbwwyyr orwogbygb gbbrrwrrw ooygbygbr oogoogyyb rrwgywgyo"
// with white, green, red, blue, orange and yellow faces as described earlier.
// "wwwwwwwww ggggggggg rrrrrrrrr bbbbbbbbb ooooooooo yyyyyyyyy"
// would correspond to a solved cube.
// No validity check is performed.
func NewCube(s string) Cube {
	stripped := strings.Replace(s, " ", "", -1)
	if len(stripped) != 9*6 {
		panic("Not a valid cube: " + s)
	}
	cube := make([]rune, 9*6)
	for i, c := range stripped {
		cube[i] = c
	}
	return cube
}

// Build a new solved Cube.
func NewSolvedCube() Cube {
	cube := make([]rune, 9*6)
	colors := []rune{WHITE, GREEN, RED, BLUE, ORANGE, YELLOW}
	for i, color := range colors {
		for j := 0; j < 9; j++ {
			cube[i*9+j] = color
		}
	}
	return cube
}

// String representation.
func (cube Cube) String() string {
	var s string
	for i, c := range cube {
		if i != 0 && i%9 == 0 {
			s = s + " "
		}
		s = s + string(c)
	}
	return s
}

// Return `true` if the given face is solved,
// aka made of a unique color.
func faceIsSolved(face []rune) bool {
	ref := face[0]
	for _, color := range face {
		if color != ref {
			return false
		}
	}
	return true
}

// Return `true` is this cube is solved.
func (cube Cube) IsSolved() bool {
	for f := 0; f < 6*9; f = f + 9 {
		face := cube[f : f+9]
		if !faceIsSolved(face) {
			return false
		}
	}
	return true
}

// Check if two cubes are equal.
func (cube Cube) Equals(other Cube) bool {
	for i, c := range cube {
		if c != other[i] {
			return false
		}
	}
	return true
}

// Clone a cube.
func (cube Cube) Copy() Cube {
	clone := make([]rune, len(cube))
	copy(clone, cube)
	return clone
}

// Turn the cube using the given move, and return a new Cube.
func (cube Cube) Turn(move Move) Cube {
	switch move {
	case UP:
		return cube.U()
	case UP_COUNTER:
		return cube.Uc()
	case DOWN:
		return cube.D()
	case DOWN_COUNTER:
		return cube.Dc()
	case LEFT:
		return cube.L()
	case LEFT_COUNTER:
		return cube.Lc()
	case RIGHT:
		return cube.R()
	case RIGHT_COUNTER:
		return cube.Rc()
	case FRONT:
		return cube.F()
	case FRONT_COUNTER:
		return cube.Fc()
	case BACK:
		return cube.B()
	case BACK_COUNTER:
		return cube.Bc()
	default:
		panic("Unknown move: " + move)
	}
}

// Return a new Cube after Pair.To becomes Pair.From.
func (cube Cube) exchange(pairs []*Pair) Cube {
	clone := cube.Copy()
	for _, pair := range pairs {
		clone[pair.To] = cube[pair.From]
	}
	return clone
}

// Rotate a face right, starting from `index`.
func (cube Cube) rotateRight(index int) Cube {
	clone := cube.Copy()
	displacements := []int{6, 3, 0, 7, 4, 1, 8, 5, 2}
	for i, d := range displacements {
		clone[i+index] = cube[d+index]
	}
	return clone
}

// Rotate a face left, starting from `index`.
func (cube Cube) rotateLeft(index int) Cube {
	clone := cube.Copy()
	displacements := []int{2, 5, 8, 1, 4, 7, 0, 3, 6}
	for i, d := range displacements {
		clone[i+index] = cube[d+index]
	}
	return clone
}

// Move Front face clockwise, and return a new Cube.
func (cube Cube) F() Cube {
	pairs := []*Pair{
		&Pair{6, 18}, &Pair{7, 21}, &Pair{8, 24},
		&Pair{18, 47}, &Pair{21, 46}, &Pair{24, 45},
		&Pair{45, 38}, &Pair{46, 41}, &Pair{47, 44},
		&Pair{38, 8}, &Pair{41, 7}, &Pair{44, 6},
	}

	return cube.exchange(pairs).rotateRight(9)
}

// Move Front face counter clockwise, and return a new Cube.
func (cube Cube) Fc() Cube {
	pairs := []*Pair{
		&Pair{6, 44}, &Pair{7, 41}, &Pair{8, 38},
		&Pair{38, 45}, &Pair{41, 46}, &Pair{44, 47},
		&Pair{45, 24}, &Pair{46, 21}, &Pair{47, 18},
		&Pair{18, 6}, &Pair{21, 7}, &Pair{24, 8},
	}
	return cube.exchange(pairs).rotateLeft(9)
}

// Move Back face clockwise, and return a new Cube.
func (cube Cube) B() Cube {
	pairs := []*Pair{
		&Pair{0, 42}, &Pair{1, 39}, &Pair{2, 36},
		&Pair{36, 51}, &Pair{39, 52}, &Pair{42, 53},
		&Pair{53, 20}, &Pair{52, 23}, &Pair{51, 26},
		&Pair{20, 0}, &Pair{23, 1}, &Pair{26, 2},
	}
	return cube.exchange(pairs).rotateRight(27)
}

// Move Back face counter clockwise, and return a new Cube.
func (cube Cube) Bc() Cube {
	pairs := []*Pair{
		&Pair{0, 20}, &Pair{1, 23}, &Pair{2, 26},
		&Pair{20, 53}, &Pair{23, 52}, &Pair{26, 51},
		&Pair{51, 36}, &Pair{52, 39}, &Pair{53, 42},
		&Pair{36, 2}, &Pair{39, 1}, &Pair{42, 0},
	}
	return cube.exchange(pairs).rotateLeft(27)
}

// Move Upper face clockwise, and return a new Cube.
func (cube Cube) U() Cube {
	pairs := []*Pair{
		&Pair{9, 36}, &Pair{10, 37}, &Pair{11, 38},
		&Pair{36, 27}, &Pair{37, 28}, &Pair{38, 29},
		&Pair{27, 18}, &Pair{28, 19}, &Pair{29, 20},
		&Pair{18, 9}, &Pair{19, 10}, &Pair{20, 11},
	}
	return cube.exchange(pairs).rotateRight(0)
}

// Move Upper face counter clockwise, and return a new Cube.
func (cube Cube) Uc() Cube {
	pairs := []*Pair{
		&Pair{9, 18}, &Pair{10, 19}, &Pair{11, 20},
		&Pair{18, 27}, &Pair{19, 28}, &Pair{20, 29},
		&Pair{27, 36}, &Pair{28, 37}, &Pair{29, 38},
		&Pair{36, 9}, &Pair{37, 10}, &Pair{38, 11},
	}
	return cube.exchange(pairs).rotateLeft(0)
}

// Move Down face clockwise, and return a new Cube.
func (cube Cube) D() Cube {
	pairs := []*Pair{
		&Pair{15, 24}, &Pair{16, 25}, &Pair{17, 26},
		&Pair{24, 33}, &Pair{25, 34}, &Pair{26, 35},
		&Pair{33, 42}, &Pair{34, 43}, &Pair{35, 44},
		&Pair{42, 15}, &Pair{43, 16}, &Pair{44, 17},
	}
	return cube.exchange(pairs).rotateRight(45)
}

// Move Down face counter clockwise, and return a new Cube.
func (cube Cube) Dc() Cube {
	pairs := []*Pair{
		&Pair{15, 42}, &Pair{16, 43}, &Pair{17, 44},
		&Pair{42, 33}, &Pair{43, 34}, &Pair{44, 35},
		&Pair{33, 24}, &Pair{34, 25}, &Pair{35, 26},
		&Pair{24, 15}, &Pair{25, 16}, &Pair{26, 17},
	}
	return cube.exchange(pairs).rotateLeft(45)
}

// Move Left face clockwise, and return a new Cube.
func (cube Cube) L() Cube {
	pairs := []*Pair{
		&Pair{0, 9}, &Pair{3, 12}, &Pair{6, 15},
		&Pair{9, 45}, &Pair{12, 48}, &Pair{15, 51},
		&Pair{45, 35}, &Pair{48, 32}, &Pair{51, 29},
		&Pair{29, 6}, &Pair{32, 3}, &Pair{35, 0},
	}
	return cube.exchange(pairs).rotateRight(36)
}

// Move Left face counter clockwise, and return a new Cube.
func (cube Cube) Lc() Cube {
	pairs := []*Pair{
		&Pair{0, 35}, &Pair{3, 32}, &Pair{6, 29},
		&Pair{29, 51}, &Pair{32, 48}, &Pair{35, 45},
		&Pair{45, 9}, &Pair{48, 12}, &Pair{51, 15},
		&Pair{9, 0}, &Pair{12, 3}, &Pair{15, 6},
	}
	return cube.exchange(pairs).rotateLeft(36)
}

// Move Right face clockwise, and return a new Cube.
func (cube Cube) R() Cube {
	pairs := []*Pair{
		&Pair{2, 33}, &Pair{5, 30}, &Pair{8, 27},
		&Pair{27, 53}, &Pair{30, 50}, &Pair{33, 47},
		&Pair{47, 11}, &Pair{50, 14}, &Pair{53, 17},
		&Pair{11, 2}, &Pair{14, 5}, &Pair{17, 8},
	}
	return cube.exchange(pairs).rotateRight(18)
}

// Move Right face counter clockwise, and return a new Cube.
func (cube Cube) Rc() Cube {
	pairs := []*Pair{
		&Pair{2, 11}, &Pair{5, 14}, &Pair{8, 17},
		&Pair{11, 47}, &Pair{14, 50}, &Pair{17, 53},
		&Pair{47, 33}, &Pair{50, 30}, &Pair{53, 27},
		&Pair{27, 8}, &Pair{30, 5}, &Pair{33, 2},
	}
	return cube.exchange(pairs).rotateLeft(18)
}
