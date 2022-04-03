package rubik_alg

import (
	"math/rand"
	"strings"
	"time"

	"github.com/hultan/go-rubik/src/rubik"
)

// ------------------------------- //
// SAMPLE PLL ALGS                 //
// http://algdb.net/puzzle/333/pll //
// ------------------------------- //
const (
	PllPermAa = "x (R' U R') D2 (R U' R') D2 R2 x'"
	PllPermAb = "x R2 D2 (R U R') D2 (R U' R) x'"
	PllPermE  = "z U2 R2 F (R U R' U') (R U R' U') (R U R' U') F' R2 U2 z'"
	PllPermF  = "R' U' F' (R U R' U') (R' F R2 U') R' U' (R U R' U R)"
	PllPermH  = "M2 U M2 U2 M2 U M2"
	PllPermZ  = "M2 U M2 U M' U2 M2 U2 M' U2"
	PllPermJa = "(R' U L') U2 (R U' R' U2) L R U'"
	PllPermJb = "(R U R' F') (R U R' U') (R' F R2 U') R' U'"
	PllPermT  = "(R U R' U') (R' F R2 U') R' U' (R U R' F')"
	PllPermNa = "(L U' R U2 L' U R') (L U' R U2 L' U R') U"
	PllPermNb = "(R' U L' U2 R U' L) (R' U L' U2 R U' L) U'"
	PllPermRa = "(R U R' F') (R U2' R' U2') (R' F R U) (R U2' R') [U']"
	PllPermRb = "(R' U2 R U2') R' F (R U R' U') R' F' R2 [U']"
	PllPermUa = "R U' R U R U R U' R' U' R2"
	PllPermUb = "R2 U R U R' U' R' U' R' U R'"
	PllPermV  = "R' U R' U' y R' F' R2 U' R' U R' F R F"
	PllPermY  = "R2 U' R2 U' R2 U R' F' R U R2 U' R' F R"
	PllPermGa = "R2 u R' U R' U' R u' R2 y' R' U R"
	PllPermGb = "L' U' L y L2 u L' U L U' L u' L2"
	PllPermGc = "L2 u' L U' L U L' u L2 y L U' L'"
	PllPermGd = "R U R' y' R2 u' R U' R' U R' u R2"
)

var validMoves = []string{
	"R ", "R' ", "R2 ", "L ", "L' ", "L2 ",
	"U ", "U' ", "U2 ", "D ", "D' ", "D2 ",
	"F ", "F' ", "F2 ", "B ", "B' ", "B2 ",
	"r ", "r' ", "r2 ", "l ", "l' ", "l2 ",
	"u ", "u' ", "u2 ", "d ", "d' ", "d2 ",
	"f ", "f' ", "f2 ", "b ", "b' ", "b2 ",
	"M ", "M' ", "M2 ", "E ", "E' ", "E2 ", "S ", "S' ", "S2 ",
	"x ", "x' ", "x2 ", "y ", "y' ", "y2 ", "z ", "z' ", "z2 ",
}

// GetScrambledCube : Get a scrambled sube
func GetScrambledCube() (rubik.Cube, string) {
	alg := ""
	rand.Seed(time.Now().UnixNano())

	// Create a valid scramble
	for i := 0; i < 30; i++ {
		a := rand.Intn(len(validMoves))
		alg += validMoves[a]
	}

	return ExecuteAlg(rubik.NewSolvedCube(), alg), alg
}

// ReverseAlg returns the reverse of an algorithm
func ReverseAlg(alg string) string {
	var reverse []string

	// Remove ():s and []:s
	alg = cleanAlg(alg)

	// Split up alg, reverse move, and copy to a slice
	moves := strings.Split(alg, " ")
	for _, move := range moves {
		move = strings.Trim(move, " ")
		move = reverseMove(move)
		if move == "" {
			continue
		}
		reverse = append(reverse, move)
	}

	// Reverse the slice
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}

	// Return the reverse alg
	return strings.Join(reverse, " ")
}

// ExecuteAlg executes the provided algorithm on the provided cube,
// and returns a new cube.
func ExecuteAlg(cube rubik.Cube, alg string) rubik.Cube {
	// Remove ():s and []:s
	alg = cleanAlg(alg)

	// Split up alg, and execute moves
	moves := strings.Split(alg, " ")
	for _, move := range moves {
		move = strings.Trim(move, " ")
		cube = executeMove(cube, move)
	}

	// Return the new cube
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
		// Invalid moves are ignored
		return ""
	}
}

// Performs the provided move on the provided cube, and returns a new cube
// Illegal moves are ignored, and just returns the cube as it was
func executeMove(cube rubik.Cube, move string) rubik.Cube {
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
		// Invalid moves are ignored
		return cube
	}
}

// cleanAlg removes spaces, enter, tabs, () and []
func cleanAlg(alg string) string {
	alg = strings.Trim(alg, " \t\n\r")
	alg = strings.Replace(alg, "(", "", -1)
	alg = strings.Replace(alg, ")", "", -1)
	alg = strings.Replace(alg, "[", "", -1)
	alg = strings.Replace(alg, "]", "", -1)
	return alg
}
