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

func testUUUU(t *testing.T) {
    origin := NewSolvedCube()
    circle := origin.U().U().U().U()
    if !origin.Equals(circle) {
        t.Error("o U U U U should return to o", circle)
    }
}

func testUcUcUcUc(t *testing.T) {
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

func testDDDD(t *testing.T) {
    origin := NewSolvedCube()
    circle := origin.D().D().D().D()
    if !origin.Equals(circle) {
        t.Error("o D D D D should return to o", circle)
    }
}

func testDcDcDcDc(t *testing.T) {
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

func testLLLL(t *testing.T) {
    origin := NewSolvedCube()
    circle := origin.L().L().L().L()
    if !origin.Equals(circle) {
        t.Error("o L L L L should return to o", circle)
    }
}

func testLcLcLcLc(t *testing.T) {
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

func testRRRR(t *testing.T) {
    origin := NewSolvedCube()
    circle := origin.R().R().R().R()
    if !origin.Equals(circle) {
        t.Error("o R R R R should return to o", circle)
    }
}

func testRcRcRcRc(t *testing.T) {
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
