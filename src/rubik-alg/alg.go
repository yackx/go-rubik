package rubik_alg

import (
	"fmt"
	"strings"

	"github.com/hultan/go-rubik/src/rubik"
)

// ReverseAlg returns the reverse of an algorithm
func ReverseAlg(alg string) string {
	alg = cleanAlg(alg)
	moves := strings.Split(alg, " ")
	var reverse []string

	for _, move := range moves {
		move = strings.Trim(move, " ")
		if move == "" {
			continue
		}
		reverse = append(reverse, reverseMove(move))
	}

	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	return strings.Join(reverse, " ")
}

// PerformAlg performs the provided algorithm on the provided cube,
// and returns a new cube.
func PerformAlg(cube rubik.Cube, alg string) rubik.Cube {
	alg = cleanAlg(alg)
	moves := strings.Split(alg, " ")

	for _, move := range moves {
		move = strings.Trim(move, " ")
		if move == "" {
			continue
		}
		cube = performMove(cube, move)
	}

	return cube
}

// Returns the reverse of the provided move
func reverseMove(move string) string {
	switch move {
	case "R":
		return "R'"
	case "R2":
		return "R2"
	case "R'":
		return "R"
	case "L":
		return "L1"
	case "L2":
		return "L2"
	case "L'":
		return "L"

	case "U":
		return "U'"
	case "U2":
		return "U2"
	case "U'":
		return "U"
	case "D":
		return "D'"
	case "D2":
		return "D2"
	case "D'":
		return "D"

	case "F":
		return "F'"
	case "F2":
		return "F2"
	case "F'":
		return "F"
	case "B":
		return "B'"
	case "B2":
		return "B2"
	case "B'":
		return "B"

	case "M":
		return "M'"
	case "M2":
		return "M2"
	case "M'":
		return "M"
	case "E":
		return "E'"
	case "E2":
		return "E2"
	case "E'":
		return "E"
	case "S":
		return "S'"
	case "S2":
		return "S2"
	case "S'":
		return "S"

	case "u":
		return "u'"
	case "u'":
		return "u"
	case "u2":
		return "u2"
	case "d":
		return "d'"
	case "d'":
		return "d"
	case "d2":
		return "d2"

	case "f":
		return "f'"
	case "f'":
		return "f"
	case "f2":
		return "f2"
	case "b":
		return "b'"
	case "b'":
		return "b"
	case "b2":
		return "b2"

	case "l":
		return "l'"
	case "l'":
		return "l"
	case "l2":
		return "l2"
	case "r":
		return "r'"
	case "r'":
		return "r"
	case "r2":
		return "r2"

	case "x":
		return "x'"
	case "x2":
		return "x2"
	case "x'":
		return "x"
	case "y":
		return "y'"
	case "y2":
		return "y2"
	case "y'":
		return "y"
	case "z":
		return "z'"
	case "z2":
		return "z2"
	case "z'":
		return "z"

	default:
		panic(fmt.Sprintf("illegal move : %v", move))
	}
}

// Performs the provided move on the provided cube, and returns a new cube
func performMove(cube rubik.Cube, move string) rubik.Cube {
	switch move {
	case "R":
		return cube.R()
	case "R2":
		return cube.R().R()
	case "R'":
		return cube.Rc()
	case "L":
		return cube.L()
	case "L2":
		return cube.L().L()
	case "L'":
		return cube.Lc()

	case "U":
		return cube.U()
	case "U2":
		return cube.U().U()
	case "U'":
		return cube.Uc()
	case "D":
		return cube.D()
	case "D2":
		return cube.D().D()
	case "D'":
		return cube.Dc()

	case "F":
		return cube.F()
	case "F2":
		return cube.F().F()
	case "F'":
		return cube.Fc()
	case "B":
		return cube.B()
	case "B2":
		return cube.B().B()
	case "B'":
		return cube.Bc()

	case "M":
		return cube.M()
	case "M2":
		return cube.M().M()
	case "M'":
		return cube.Mc()
	case "E":
		return cube.E()
	case "E2":
		return cube.E().E()
	case "E'":
		return cube.Ec()
	case "S":
		return cube.S()
	case "S2":
		return cube.S().S()
	case "S'":
		return cube.Sc()

	case "u":
		return cube.Ec().U()
	case "u'":
		return cube.E().Uc()
	case "u2":
		return cube.U().U().Ec().Ec()
	case "d":
		return cube.E().D()
	case "d'":
		return cube.Ec().Dc()
	case "d2":
		return cube.D().D().E().E()

	case "f":
		return cube.S().F()
	case "f'":
		return cube.Sc().Fc()
	case "f2":
		return cube.S().S().F().F()
	case "b":
		return cube.Sc().B()
	case "b'":
		return cube.S().Bc()
	case "b2":
		return cube.Sc().Sc().B().B()

	case "l":
		return cube.L().M()
	case "l'":
		return cube.Lc().Mc()
	case "l2":
		return cube.L().L().M().M()
	case "r":
		return cube.R().Mc()
	case "r'":
		return cube.Rc().M()
	case "r2":
		return cube.R().R().Mc().Mc()

	case "x":
		return cube.L().M().Rc()
	case "x'":
		return cube.Lc().Mc().R()
	case "y":
		return cube.U().Ec().Dc()
	case "y'":
		return cube.Uc().E().D()
	case "z":
		return cube.F().S().Bc()
	case "z'":
		return cube.Fc().Sc().B()

	default:
		panic(fmt.Sprintf("illegal move : %v", move))
	}
}

func cleanAlg(alg string) string {
	alg = strings.Replace(alg, "(", "", -1)
	alg = strings.Replace(alg, ")", "", -1)
	alg = strings.Replace(alg, "[", "", -1)
	alg = strings.Replace(alg, "]", "", -1)
	return alg
}
