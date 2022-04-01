// ------------------------------------------------------------ //
// THIS CODE WAS WRITTEN BY YACKX. SLICE MOVES ADDED BY HULTAN. //
// See yackx/go-rubik : https://github.com/yackx/go-rubik	    //
// ------------------------------------------------------------ //

// Package rubik : Implements a rubik's cube
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
		{6, 18}, {7, 21}, {8, 24},
		{18, 47}, {21, 46}, {24, 45},
		{45, 38}, {46, 41}, {47, 44},
		{38, 8}, {41, 7}, {44, 6},
	}

	return cube.exchange(pairs).rotateRight(9)
}

// Move Front face counter clockwise, and return a new Cube.
func (cube Cube) Fc() Cube {
	pairs := []*Pair{
		{6, 44}, {7, 41}, {8, 38},
		{38, 45}, {41, 46}, {44, 47},
		{45, 24}, {46, 21}, {47, 18},
		{18, 6}, {21, 7}, {24, 8},
	}
	return cube.exchange(pairs).rotateLeft(9)
}

// Move Back face clockwise, and return a new Cube.
func (cube Cube) B() Cube {
	pairs := []*Pair{
		{0, 42}, {1, 39}, {2, 36},
		{36, 51}, {39, 52}, {42, 53},
		{53, 20}, {52, 23}, {51, 26},
		{20, 0}, {23, 1}, {26, 2},
	}
	return cube.exchange(pairs).rotateRight(27)
}

// Move Back face counter clockwise, and return a new Cube.
func (cube Cube) Bc() Cube {
	pairs := []*Pair{
		{0, 20}, {1, 23}, {2, 26},
		{20, 53}, {23, 52}, {26, 51},
		{51, 36}, {52, 39}, {53, 42},
		{36, 2}, {39, 1}, {42, 0},
	}
	return cube.exchange(pairs).rotateLeft(27)
}

// Move Upper face clockwise, and return a new Cube.
func (cube Cube) U() Cube {
	pairs := []*Pair{
		{9, 36}, {10, 37}, {11, 38},
		{36, 27}, {37, 28}, {38, 29},
		{27, 18}, {28, 19}, {29, 20},
		{18, 9}, {19, 10}, {20, 11},
	}
	return cube.exchange(pairs).rotateRight(0)
}

// Move Upper face counter clockwise, and return a new Cube.
func (cube Cube) Uc() Cube {
	pairs := []*Pair{
		{9, 18}, {10, 19}, {11, 20},
		{18, 27}, {19, 28}, {20, 29},
		{27, 36}, {28, 37}, {29, 38},
		{36, 9}, {37, 10}, {38, 11},
	}
	return cube.exchange(pairs).rotateLeft(0)
}

// Move Down face clockwise, and return a new Cube.
func (cube Cube) D() Cube {
	pairs := []*Pair{
		{15, 24}, {16, 25}, {17, 26},
		{24, 33}, {25, 34}, {26, 35},
		{33, 42}, {34, 43}, {35, 44},
		{42, 15}, {43, 16}, {44, 17},
	}
	return cube.exchange(pairs).rotateRight(45)
}

// Move Down face counter clockwise, and return a new Cube.
func (cube Cube) Dc() Cube {
	pairs := []*Pair{
		{15, 42}, {16, 43}, {17, 44},
		{42, 33}, {43, 34}, {44, 35},
		{33, 24}, {34, 25}, {35, 26},
		{24, 15}, {25, 16}, {26, 17},
	}
	return cube.exchange(pairs).rotateLeft(45)
}

// Move Left face clockwise, and return a new Cube.
func (cube Cube) L() Cube {
	pairs := []*Pair{
		{0, 9}, {3, 12}, {6, 15},
		{9, 45}, {12, 48}, {15, 51},
		{45, 35}, {48, 32}, {51, 29},
		{29, 6}, {32, 3}, {35, 0},
	}
	return cube.exchange(pairs).rotateRight(36)
}

// Move Left face counter clockwise, and return a new Cube.
func (cube Cube) Lc() Cube {
	pairs := []*Pair{
		{0, 35}, {3, 32}, {6, 29},
		{29, 51}, {32, 48}, {35, 45},
		{45, 9}, {48, 12}, {51, 15},
		{9, 0}, {12, 3}, {15, 6},
	}
	return cube.exchange(pairs).rotateLeft(36)
}

// Move Right face clockwise, and return a new Cube.
func (cube Cube) R() Cube {
	pairs := []*Pair{
		{2, 33}, {5, 30}, {8, 27},
		{27, 53}, {30, 50}, {33, 47},
		{47, 11}, {50, 14}, {53, 17},
		{11, 2}, {14, 5}, {17, 8},
	}
	return cube.exchange(pairs).rotateRight(18)
}

// Move Right face counter clockwise, and return a new Cube.
func (cube Cube) Rc() Cube {
	pairs := []*Pair{
		{2, 11}, {5, 14}, {8, 17},
		{11, 47}, {14, 50}, {17, 53},
		{47, 33}, {50, 30}, {53, 27},
		{27, 8}, {30, 5}, {33, 2},
	}
	return cube.exchange(pairs).rotateLeft(18)
}

// Move middle layer clockwise, and return a new Cube.
func (cube Cube) M() Cube {
	pairs := []*Pair{
		{52, 34}, {49, 31}, {46, 28},
		{28, 7}, {31, 4}, {34, 1},
		{7, 16}, {4, 13}, {1, 10},
		{16, 52}, {13, 49}, {10, 46},
	}
	return cube.exchange(pairs)
}

// Move middle layer counter clockwise, and return a new Cube.
func (cube Cube) Mc() Cube {
	pairs := []*Pair{
		{46, 10}, {49, 13}, {52, 16},
		{10, 1}, {13, 4}, {16, 7},
		{1, 34}, {4, 31}, {7, 28},
		{28, 46}, {31, 49}, {34, 52},
	}
	return cube.exchange(pairs)
}

// Move slide layer clockwise, and return a new Cube.
func (cube Cube) S() Cube {
	pairs := []*Pair{
		{43, 3}, {40, 4}, {37, 5},
		{3, 19}, {4, 22}, {5, 25},
		{19, 50}, {22, 49}, {25, 48},
		{50, 43}, {49, 40}, {48, 37},
	}
	return cube.exchange(pairs)
}

// Move slide layer counter clockwise, and return a new Cube.
func (cube Cube) Sc() Cube {
	pairs := []*Pair{
		{37, 48}, {40, 49}, {43, 50},
		{48, 25}, {49, 22}, {50, 19},
		{25, 5}, {22, 4}, {19, 3},
		{5, 37}, {4, 40}, {3, 43},
	}
	return cube.exchange(pairs)
}

// Move equator layer clockwise, and return a new Cube.
func (cube Cube) E() Cube {
	pairs := []*Pair{
		{12, 21}, {13, 22}, {14, 23},
		{21, 30}, {22, 31}, {23, 32},
		{30, 39}, {31, 40}, {32, 41},
		{39, 12}, {40, 13}, {41, 14},
	}
	return cube.exchange(pairs)
}

// Move equator layer counter clockwise, and return a new Cube.
func (cube Cube) Ec() Cube {
	pairs := []*Pair{
		{14, 41}, {13, 40}, {12, 39},
		{41, 32}, {40, 31}, {39, 30},
		{32, 23}, {31, 22}, {30, 21},
		{23, 14}, {22, 13}, {21, 12},
	}
	return cube.exchange(pairs)
}

//
// DOUBLE MOVES (u u' d d' f f' b b' r r' l l')
//

// Move upper two layers clockwise, and return a new cube
func (cube Cube) Ud() Cube {
	return cube.Ec().U()
}

// Move upper two layers counter clockwise, and return a new cube
func (cube Cube) Udc() Cube {
	return cube.E().Uc()
}

// Move lower two layers clockwise, and return a new cube
func (cube Cube) Dd() Cube {
	return cube.E().D()
}

// Move lower two layers counter clockwise, and return a new cube
func (cube Cube) Ddc() Cube {
	return cube.Ec().Dc()
}

// Move front two layers clockwise, and return a new cube
func (cube Cube) Fd() Cube {
	return cube.S().F()
}

// Move front two layers counter clockwise, and return a new cube
func (cube Cube) Fdc() Cube {
	return cube.Sc().Fc()
}

// Move back two layers clockwise, and return a new cube
func (cube Cube) Bd() Cube {
	return cube.Sc().B()
}

// Move back two layers counter clockwise, and return a new cube
func (cube Cube) Bdc() Cube {
	return cube.S().Bc()
}

// Move left two layers clockwise, and return a new cube
func (cube Cube) Ld() Cube {
	return cube.L().M()
}

// Move left two layers counter clockwise, and return a new cube
func (cube Cube) Ldc() Cube {
	return cube.Lc().Mc()
}

// Move right two layers clockwise, and return a new cube
func (cube Cube) Rd() Cube {
	return cube.R().Mc()
}

// Move right two layers counter clockwise, and return a new cube
func (cube Cube) Rdc() Cube {
	return cube.Rc().M()
}

//
// CUBE ROTATIONS (ADDED BY HULTAN)
//

// Rotate the cube around the x axis clockwise, and return a new cube
func (cube Cube) X() Cube {
	return cube.L().M().Rc()
}

// Rotate the cube around the x axis counter clockwise, and return a new cube
func (cube Cube) Xc() Cube {
	return cube.Lc().Mc().R()
}

// Rotate the cube around the y axis clockwise, and return a new cube
func (cube Cube) Y() Cube {
	return cube.U().Ec().Dc()
}

// Rotate the cube around the y axis counter clockwise, and return a new cube
func (cube Cube) Yc() Cube {
	return cube.Uc().E().D()
}

// Rotate the cube around the z axis clockwise, and return a new cube
func (cube Cube) Z() Cube {
	return cube.F().S().Bc()
}

// Rotate the cube around the z axis counter clockwise, and return a new cube
func (cube Cube) Zc() Cube {
	return cube.Fc().Sc().B()
}
