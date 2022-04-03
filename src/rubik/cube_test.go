package rubik

import (
	"testing"
)

func TestNotSolved(t *testing.T) {
	cube := NewCube("yrrrrrrrrbbbbbbbbbooooooooogggggggggwwwwwwwwwyyyyyyyyr")
	if cube.IsSolved() {
		t.Error("Cube should not be solved", cube)
	}
}

func TestNewSolved(t *testing.T) {
	solved := NewSolvedCube()
	if !solved.IsSolved() {
		t.Error("New solved cube should be solved", solved)
	}
}

func TestSolved(t *testing.T) {
	cube := NewCube("wwwwwwwww ggggggggg rrrrrrrrr bbbbbbbbb ooooooooo yyyyyyyyy")
	if !cube.IsSolved() {
		t.Error("Cube should be solved", cube)
	}
}

func TestEquals(t *testing.T) {
	cube1 := NewCube("wwwwwwwww ggggggggg rrrrrrrrr bbbbbbbbb ooooooooo yyyyyyyyy")
	cube2 := NewSolvedCube()
	if !cube1.Equals(cube2) {
		t.Error("Cubes should be equal", cube1, cube2)
	}

	cube3 := cube2.F()
	if cube3.Equals(cube2) {
		t.Error("Cubes should not be equal", cube1, cube3)
	}
}

//
// FRONT
//

func TestF(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.F()
	target := NewCube("wwwwwwooo ggggggggg wrrwrrwrr bbbbbbbbb ooyooyooy rrryyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o F incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestFc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Fc()
	target := NewCube("wwwwwwrrr ggggggggg yrryrryrr bbbbbbbbb oowoowoow oooyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o F' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestFFc(t *testing.T) {
	origin := NewSolvedCube()
	back := origin.F().Fc()
	if !back.Equals(origin) {
		t.Errorf("o F F' incorrect:\nGot  %s\nWant %s", back, origin)
	}
}

func TestFFFF(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.F().F().F().F()
	if !origin.Equals(circle) {
		t.Error("o F F F F should return to o", circle)
	}
}

func TestFcFcFcFc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Fc().Fc().Fc().Fc()
	if !origin.Equals(circle) {
		t.Error("o F' F' F' F' should return to o", circle)
	}
}

//
// BACK
//

func TestB(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.B()
	target := NewCube("rrrwwwwww ggggggggg rryrryrry bbbbbbbbb woowoowoo yyyyyyooo")
	if !target.Equals(turned) {
		t.Errorf("o B incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestBc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Bc()
	target := NewCube("ooowwwwww ggggggggg rrwrrwrrw bbbbbbbbb yooyooyoo yyyyyyrrr")
	if !target.Equals(turned) {
		t.Errorf("o B' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestBBc(t *testing.T) {
	origin := NewSolvedCube()
	back := origin.B().Bc()
	if !back.Equals(origin) {
		t.Errorf("o B B' incorrect:\nGot  %s\nWant %s", back, origin)
	}
}

func TestBBBB(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.B().B().B().B()
	if !origin.Equals(circle) {
		t.Error("o B B B B should return to o", circle)
	}
}

func TestBcBcBcBc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Bc().Bc().Bc().Bc()
	if !origin.Equals(circle) {
		t.Error("o B' B' B' B' should return to o", circle)
	}
}

//
// UP
//

func TestU(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.U()
	target := NewCube("wwwwwwwww rrrgggggg bbbrrrrrr ooobbbbbb gggoooooo yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o U incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestUc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Uc()
	target := NewCube("wwwwwwwww ooogggggg gggrrrrrr rrrbbbbbb bbboooooo yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o LU' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestUUUU(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.U().U().U().U()
	if !origin.Equals(circle) {
		t.Error("o U U U U should return to o", circle)
	}
}

func TestUcUcUcUc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Uc().Uc().Uc().Uc()
	if !origin.Equals(circle) {
		t.Error("o U' U' U' U' should return to o", circle)
	}
}

func TestUUc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.U().Uc()
	if !origin.Equals(doAndUndo) {
		t.Error("o U U' should return to o", doAndUndo)
	}
}

//
// DOWN
//

func TestD(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.D()
	target := NewCube("wwwwwwwww ggggggooo rrrrrrggg bbbbbbrrr oooooobbb yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o D incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestDc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Dc()
	target := NewCube("wwwwwwwww ggggggrrr rrrrrrbbb bbbbbbooo ooooooggg yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o D' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestDDDD(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.D().D().D().D()
	if !origin.Equals(circle) {
		t.Error("o D D D D should return to o", circle)
	}
}

func TestDcDcDcDc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Dc().Dc().Dc().Dc()
	if !origin.Equals(circle) {
		t.Error("o D' D' D' D' should return to o", circle)
	}
}

func TestDDc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.D().Dc()
	if !origin.Equals(doAndUndo) {
		t.Error("o D D' should return to o", doAndUndo)
	}
}

//
// LEFT
//

func TestL(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.L()
	target := NewCube("bwwbwwbww wggwggwgg rrrrrrrrr bbybbybby ooooooooo gyygyygyy")
	if !target.Equals(turned) {
		t.Errorf("o L incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestLc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Lc()
	target := NewCube("gwwgwwgww yggyggygg rrrrrrrrr bbwbbwbbw ooooooooo byybyybyy")
	if !target.Equals(turned) {
		t.Errorf("o L' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestLLLL(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.L().L().L().L()
	if !origin.Equals(circle) {
		t.Error("o L L L L should return to o", circle)
	}
}

func TestLcLcLcLc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Lc().Lc().Lc().Lc()
	if !origin.Equals(circle) {
		t.Error("o L' L' L' L' should return to o", circle)
	}
}

func TestLLc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.L().Lc()
	if !origin.Equals(doAndUndo) {
		t.Error("o L L' should return to o", doAndUndo)
	}
}

//
// RIGHT
//

func TestR(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.R()
	target := NewCube("wwgwwgwwg ggyggyggy rrrrrrrrr wbbwbbwbb ooooooooo yybyybyyb")
	if !target.Equals(turned) {
		t.Errorf("o R incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestRc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Rc()
	target := NewCube("wwbwwbwwb ggwggwggw rrrrrrrrr ybbybbybb ooooooooo yygyygyyg")
	if !target.Equals(turned) {
		t.Errorf("o R' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestRRRR(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.R().R().R().R()
	if !origin.Equals(circle) {
		t.Error("o R R R R should return to o", circle)
	}
}

func TestRcRcRcRc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Rc().Rc().Rc().Rc()
	if !origin.Equals(circle) {
		t.Error("o R' R' R' R' should return to o", circle)
	}
}

func TestRRc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.R().Rc()
	if !origin.Equals(doAndUndo) {
		t.Error("o R R' should return to o", doAndUndo)
	}
}

//
// MIXING
//

func TestUFFcUc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.U().F().Fc().Uc()
	if !turned.Equals(origin) {
		t.Errorf("o U F F' U' incorrect:\nGot  %s\nWant %s", turned, origin)
	}
}

func TestFU(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.F().U()
	target := NewCube("owwowwoww wrrgggggg bbbwrrwrr ooybbbbbb gggooyooy rrryyyyyy")
	if !turned.Equals(target) {
		t.Errorf("o F U incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestFUU(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.F().U().U()
	target := NewCube("ooowwwwww bbbgggggg ooywrrwrr gggbbbbbb wrrooyooy rrryyyyyy")
	if !turned.Equals(target) {
		t.Errorf("o F U U incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func Test01(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.F().U().U().Uc()
	target := NewCube("owwowwoww wrrgggggg bbbwrrwrr ooybbbbbb gggooyooy rrryyyyyy")
	if !turned.Equals(target) {
		t.Errorf("o F U U U' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func Test02(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.F().R().R().U().L()
	target := NewCube("bwwbwwyyr orwogbygb gbbrrwrrw ooygbygbr oogoogyyb rrwgywgyo")
	if !turned.Equals(target) {
		t.Errorf("Test 02 incorrect:\nGot  %s\nWant %s", turned, target)
	}
	andBack := turned.Lc().Uc().Rc().Rc().Fc()
	if !andBack.Equals(origin) {
		t.Errorf("Test 02 not returned to origin:\nGot  %s\nWant %s", turned, target)
	}
}

//
// SLICE MOVES (ADDED BY HULTAN)
//

//
// MIDDLE
//

func TestM(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.M()
	target := NewCube("wbwwbwwbw gwggwggwg rrrrrrrrr bybbybbyb ooooooooo ygyygyygy")
	if !target.Equals(turned) {
		t.Errorf("o M incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestMc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Mc()
	target := NewCube("wgwwgwwgw gyggyggyg rrrrrrrrr bwbbwbbwb ooooooooo ybyybyyby")
	if !target.Equals(turned) {
		t.Errorf("o M' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestMMMM(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.M().M().M().M()
	if !origin.Equals(circle) {
		t.Error("o M M M M should return to o", circle)
	}
}

func TestMcMcMcMc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Mc().Mc().Mc().Mc()
	if !origin.Equals(circle) {
		t.Error("o M' M' M' M' should return to o", circle)
	}
}

func TestMMc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.M().Mc()
	if !origin.Equals(doAndUndo) {
		t.Error("o M M' should return to o", doAndUndo)
	}
}

//
// EQUATOR
//

func TestE(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.E()
	target := NewCube("wwwwwwwww gggoooggg rrrgggrrr bbbrrrbbb ooobbbooo yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o E incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestEc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Ec()
	target := NewCube("wwwwwwwww gggrrrggg rrrbbbrrr bbbooobbb ooogggooo yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o E' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestEEEE(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.E().E().E().E()
	if !origin.Equals(circle) {
		t.Error("o E E E E should return to o", circle)
	}
}

func TestEcEcEcEc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Ec().Ec().Ec().Ec()
	if !origin.Equals(circle) {
		t.Error("o E' E' E' E' should return to o", circle)
	}
}

func TestEEc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.E().Ec()
	if !origin.Equals(doAndUndo) {
		t.Error("o E E' should return to o", doAndUndo)
	}
}

//
// SIDE
//

func TestS(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.S()
	target := NewCube("wwwooowww ggggggggg rwrrwrrwr bbbbbbbbb oyooyooyo yyyrrryyy")
	if !target.Equals(turned) {
		t.Errorf("o S incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestSc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Sc()
	target := NewCube("wwwrrrwww ggggggggg ryrryrryr bbbbbbbbb owoowoowo yyyoooyyy")
	if !target.Equals(turned) {
		t.Errorf("o S' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestSSSS(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.S().S().S().S()
	if !origin.Equals(circle) {
		t.Error("o S S S S should return to o", circle)
	}
}

func TestScScScSc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Sc().Sc().Sc().Sc()
	if !origin.Equals(circle) {
		t.Error("o S' S' S' S' should return to o", circle)
	}
}

func TestSSc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.S().Sc()
	if !origin.Equals(doAndUndo) {
		t.Error("o S S' should return to o", doAndUndo)
	}
}

//
// DOUBLE MOVES (ADDED BY HULTAN)
//

//
// Ud AND Udc
//

func TestUd(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Ud()
	target := NewCube("wwwwwwwww rrrrrrggg bbbbbbrrr oooooobbb ggggggooo yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o u incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestUdc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Udc()
	target := NewCube("wwwwwwwww ooooooggg ggggggrrr rrrrrrbbb bbbbbbooo yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o u' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestUdUdUdUd(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Ud().Ud().Ud().Ud()
	if !origin.Equals(circle) {
		t.Error("o u u u u should return to o", circle)
	}
}

func TestUdcUdcUdcUdc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Udc().Udc().Udc().Udc()
	if !origin.Equals(circle) {
		t.Error("o u' u' u' u' should return to o", circle)
	}
}

func TestUdUdc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Ud().Udc()
	if !origin.Equals(doAndUndo) {
		t.Error("o u u' should return to o", doAndUndo)
	}
}

//
// Dd AND Ddc
//

func TestDd(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Dd()
	target := NewCube("wwwwwwwww gggoooooo rrrgggggg bbbrrrrrr ooobbbbbb yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o d incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestDdc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Ddc()
	target := NewCube("wwwwwwwww gggrrrrrr rrrbbbbbb bbboooooo ooogggggg yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o d' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestDdDdDdDd(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Dd().Dd().Dd().Dd()
	if !origin.Equals(circle) {
		t.Error("o d d d d should return to o", circle)
	}
}

func TestDdcDdcDdcDdc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Ddc().Ddc().Ddc().Ddc()
	if !origin.Equals(circle) {
		t.Error("o d' d' d' d' should return to o", circle)
	}
}

func TestDdDdc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Dd().Ddc()
	if !origin.Equals(doAndUndo) {
		t.Error("o d d' should return to o", doAndUndo)
	}
}

//
// Fd AND Fdc
//

func TestFd(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Fd()
	target := NewCube("wwwoooooo ggggggggg wwrwwrwwr bbbbbbbbb oyyoyyoyy rrrrrryyy")
	if !target.Equals(turned) {
		t.Errorf("o f incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestFdc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Fdc()
	target := NewCube("wwwrrrrrr ggggggggg yyryyryyr bbbbbbbbb owwowwoww ooooooyyy")
	if !target.Equals(turned) {
		t.Errorf("o f' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestFdFdFdFd(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Fd().Fd().Fd().Fd()
	if !origin.Equals(circle) {
		t.Error("o f f f f should return to o", circle)
	}
}

func TestFdcFdcFdcFdc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Fdc().Fdc().Fdc().Fdc()
	if !origin.Equals(circle) {
		t.Error("o f' f' f' f' should return to o", circle)
	}
}

func TestFdFdc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Fd().Fdc()
	if !origin.Equals(doAndUndo) {
		t.Error("o f f' should return to o", doAndUndo)
	}
}

//
// Bd AND Bdc
//

func TestBd(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Bd()
	target := NewCube("rrrrrrwww ggggggggg ryyryyryy bbbbbbbbb wwowwowwo yyyoooooo")
	if !target.Equals(turned) {
		t.Errorf("o b incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestBdc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Bdc()
	target := NewCube("oooooowww ggggggggg rwwrwwrww bbbbbbbbb yyoyyoyyo yyyrrrrrr")
	if !target.Equals(turned) {
		t.Errorf("o b' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestBdBdBdBd(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Bd().Bd().Bd().Bd()
	if !origin.Equals(circle) {
		t.Error("o b b b b should return to o", circle)
	}
}

func TestBdcBdcBdcBdc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Bdc().Bdc().Bdc().Bdc()
	if !origin.Equals(circle) {
		t.Error("o b' b' b' b' should return to o", circle)
	}
}

func TestBdBdc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Bd().Bdc()
	if !origin.Equals(doAndUndo) {
		t.Error("o b b' should return to o", doAndUndo)
	}
}

//
// Ld AND Ldc
//

func TestLd(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Ld()
	target := NewCube("bbwbbwbbw wwgwwgwwg rrrrrrrrr byybyybyy ooooooooo ggyggyggy")
	if !target.Equals(turned) {
		t.Errorf("o l incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestLdc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Ldc()
	target := NewCube("ggwggwggw yygyygyyg rrrrrrrrr bwwbwwbww ooooooooo bbybbybby")
	if !target.Equals(turned) {
		t.Errorf("o l' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestLdLdLdLd(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Ld().Ld().Ld().Ld()
	if !origin.Equals(circle) {
		t.Error("o l l l l should return to o", circle)
	}
}

func TestLdcLdcLdcLdc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Ldc().Ldc().Ldc().Ldc()
	if !origin.Equals(circle) {
		t.Error("o l' l' l' l' should return to o", circle)
	}
}

func TestLdLdc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Ld().Ldc()
	if !origin.Equals(doAndUndo) {
		t.Error("o l l' should return to o", doAndUndo)
	}
}

//
// Rd AND Rdc
//

func TestRd(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Rd()
	target := NewCube("wggwggwgg gyygyygyy rrrrrrrrr wwbwwbwwb ooooooooo ybbybbybb")
	if !target.Equals(turned) {
		t.Errorf("o r incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestRdc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Rdc()
	target := NewCube("wbbwbbwbb gwwgwwgww rrrrrrrrr yybyybyyb ooooooooo yggyggygg")
	if !target.Equals(turned) {
		t.Errorf("o r' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestRdRdRdRd(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Rd().Rd().Rd().Rd()
	if !origin.Equals(circle) {
		t.Error("o r r r r should return to o", circle)
	}
}

func TestRdcRdcRdcRdc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Rdc().Rdc().Rdc().Rdc()
	if !origin.Equals(circle) {
		t.Error("o r' r' r' r' should return to o", circle)
	}
}

func TestRdRdc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Rd().Rdc()
	if !origin.Equals(doAndUndo) {
		t.Error("o r r' should return to o", doAndUndo)
	}
}

//
// ROTATIONS (ADDED BY HULTAN)
//

//
// X AND X'
//

func TestX(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.X()
	target := NewCube("bbbbbbbbb wwwwwwwww rrrrrrrrr yyyyyyyyy ooooooooo ggggggggg")
	if !target.Equals(turned) {
		t.Errorf("o x incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestXc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Xc()
	target := NewCube("ggggggggg yyyyyyyyy rrrrrrrrr wwwwwwwww ooooooooo bbbbbbbbb")
	if !target.Equals(turned) {
		t.Errorf("o x' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestXXXX(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.X().X().X().X()
	if !origin.Equals(circle) {
		t.Error("o x x x x should return to o", circle)
	}
}

func TestXcXcXcXc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Xc().Xc().Xc().Xc()
	if !origin.Equals(circle) {
		t.Error("o x' x' x' x' should return to o", circle)
	}
}

func TestXXc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.X().Xc()
	if !origin.Equals(doAndUndo) {
		t.Error("o x x' should return to o", doAndUndo)
	}
}

//
// Y AND Y'
//

func TestY(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Y()
	target := NewCube("wwwwwwwww rrrrrrrrr bbbbbbbbb ooooooooo ggggggggg yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o y incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestYc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Yc()
	target := NewCube("wwwwwwwww ooooooooo ggggggggg rrrrrrrrr bbbbbbbbb yyyyyyyyy")
	if !target.Equals(turned) {
		t.Errorf("o y' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestYYYY(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Y().Y().Y().Y()
	if !origin.Equals(circle) {
		t.Error("o y y y y should return to o", circle)
	}
}

func TestYcYcYcYc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Yc().Yc().Yc().Yc()
	if !origin.Equals(circle) {
		t.Error("o y' y' y' y' should return to o", circle)
	}
}

func TestYYc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Y().Yc()
	if !origin.Equals(doAndUndo) {
		t.Error("o y y' should return to o", doAndUndo)
	}
}

//
// Z AND Z'
//

func TestZ(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Z()
	target := NewCube("ooooooooo ggggggggg wwwwwwwww bbbbbbbbb yyyyyyyyy rrrrrrrrr")
	if !target.Equals(turned) {
		t.Errorf("o z incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestZc(t *testing.T) {
	origin := NewSolvedCube()
	turned := origin.Zc()
	target := NewCube("rrrrrrrrr ggggggggg yyyyyyyyy bbbbbbbbb wwwwwwwww ooooooooo")
	if !target.Equals(turned) {
		t.Errorf("o z' incorrect:\nGot  %s\nWant %s", turned, target)
	}
}

func TestZZZZ(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Z().Z().Z().Z()
	if !origin.Equals(circle) {
		t.Error("o z z z z should return to o", circle)
	}
}

func TestZcZcZcZc(t *testing.T) {
	origin := NewSolvedCube()
	circle := origin.Zc().Zc().Zc().Zc()
	if !origin.Equals(circle) {
		t.Error("o z' z' z' z' should return to o", circle)
	}
}

func TestZZc(t *testing.T) {
	origin := NewSolvedCube()
	doAndUndo := origin.Z().Zc()
	if !origin.Equals(doAndUndo) {
		t.Error("o z z' should return to o", doAndUndo)
	}
}
