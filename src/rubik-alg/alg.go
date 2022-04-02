package rubik_alg

import (
	"strings"

	"github.com/hultan/go-rubik/src/rubik"
)

const (
	PLL_Aa_Perm = "x (R' U R') D2 (R U' R') D2 R2 x'"
	PLL_Ab_Perm = "x R2 D2 (R U R') D2 (R U' R) x'"
	PLL_E_Perm  = "z U2 R2 F (R U R' U') (R U R' U') (R U R' U') F' R2 U2 z'"
	PLL_F_Perm  = "R' U' F' (R U R' U') (R' F R2 U') R' U' (R U R' U R)"
	PLL_H_Perm  = "M2 U M2 U2 M2 U M2"
	PLL_Z_Perm  = "M2 U M2 U M' U2 M2 U2 M' U2"
	PLL_Ja_Perm = "(R' U L') U2 (R U' R' U2) L R U'"
	PLL_Jb_Perm = "(R U R' F') (R U R' U') (R' F R2 U') R' U'"
	PLL_T_Perm  = "(R U R' U') (R' F R2 U') R' U' (R U R' F')"
	PLL_Na_Perm = "(L U' R U2 L' U R') (L U' R U2 L' U R') U"
	PLL_Nb_Perm = "(R' U L' U2 R U' L) (R' U L' U2 R U' L) U'"
	PLL_Ra_Perm = "(R U R' F') (R U2' R' U2') (R' F R U) (R U2' R') [U']"
	PLL_Rb_Perm = "(R' U2 R U2') R' F (R U R' U') R' F' R2 [U']"
	PLL_Ua_Perm = "R U' R U R U R U' R' U' R2"
	PLL_Ub_Perm = "R2 U R U R' U' R' U' R' U R'"
	PLL_V_Perm  = "R' U R' U' y R' F' R2 U' R' U R' F R F"
	PLL_Y_Perm  = "R2 U' R2 U' R2 U R' F' R U R2 U' R' F R"

	PLL_Ga_Perm = "R2 u R' U R' U' R u' R2 y' R' U R"
	PLL_Gb_Perm = "L' U' L y L2 u L' U L U' L u' L2"
	PLL_Gc_Perm = "L2 u' L U' L U L' u L2 y L U' L'"
	PLL_Gd_Perm = "R U R' y' R2 u' R U' R' U R' u R2"
)

// ReverseAlg returns the reverse of an algorithm
func ReverseAlg(alg string) string {
	alg = cleanAlg(alg)
	moves := strings.Split(alg, " ")
	var reverse []string

	for _, move := range moves {
		move = strings.Trim(move, " ")
		move = reverseMove(move)
		if move == "" {
			continue
		}
		reverse = append(reverse, move)
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
		cube = performMove(cube, move)
	}

	return cube
}

// Returns the reverse of the provided move, or an empty string if it is an illegal move
func reverseMove(move string) string {
	switch move {
	case "R":
		return "R'"
	case "R2":
		return "R2"
	case "R'":
		return "R"
	case "L":
		return "L'"
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
		return ""
	}
}

// Performs the provided move on the provided cube, and returns a new cube
// Illegal moves are ignored, and just returns the cube as it was
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
		return cube.Ud()
	case "u'":
		return cube.Udc()
	case "u2":
		return cube.Ud().Ud()
	case "d":
		return cube.Dd()
	case "d'":
		return cube.Ddc()
	case "d2":
		return cube.Dd().Dd()

	case "f":
		return cube.Fd()
	case "f'":
		return cube.Fdc()
	case "f2":
		return cube.Fd().Fd()
	case "b":
		return cube.Bd()
	case "b'":
		return cube.Bdc()
	case "b2":
		return cube.Bd().Bd()

	case "l":
		return cube.Ld()
	case "l'":
		return cube.Ldc()
	case "l2":
		return cube.Ld().Ld()
	case "r":
		return cube.Rd()
	case "r'":
		return cube.Rdc()
	case "r2":
		return cube.Rd().Rd()

	case "x":
		return cube.Ldc().R()
	case "x2":
		return cube.Ldc().R().Ldc().R()
	case "x'":
		return cube.Ld().Rc()
	case "y":
		return cube.Ud().Dc()
	case "y2":
		return cube.Ud().Dc().Ud().Dc()
	case "y'":
		return cube.Udc().D()
	case "z":
		return cube.Fd().Bc()
	case "z2":
		return cube.Fd().Bc().Fd().Bc()
	case "z'":
		return cube.Fdc().B()

	default:
		return cube
	}
}

func cleanAlg(alg string) string {
	alg = strings.Trim(alg, " \t\n\r")
	alg = strings.Replace(alg, "(", "", -1)
	alg = strings.Replace(alg, ")", "", -1)
	alg = strings.Replace(alg, "[", "", -1)
	alg = strings.Replace(alg, "]", "", -1)
	return alg
}
