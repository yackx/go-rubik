package rubik

import (
    "testing"
)

type testCase struct {
    Cube Cube
    Expected []Move
}

func TestSolver(t *testing.T) {
    testCases := []*testCase {
        &testCase{NewSolvedCube().F(), []Move{FRONT_COUNTER}},
        &testCase{NewSolvedCube().Fc(), []Move{FRONT}},
        &testCase{NewSolvedCube().U(), []Move{UP_COUNTER}},
        &testCase{NewSolvedCube().F().U(), []Move{UP_COUNTER, FRONT_COUNTER}},
        &testCase{NewSolvedCube().U().U(), []Move{UP, UP}},
        &testCase{NewSolvedCube().F().U().R(), []Move{RIGHT_COUNTER, UP_COUNTER, FRONT_COUNTER}},

        // Too complex for current naive solver
        // &testCase{NewSolvedCube().F().R().R().U().L(), []Move{RIGHT_COUNTER, UP_COUNTER, FRONT_COUNTER}},
        // &testCase{NewCube("bwwbwwyyr orwogbygb gbbrrwrrw ooygbygbr oogoogyyb rrwgywgyo"), []Move{RIGHT_COUNTER, UP_COUNTER, FRONT_COUNTER}},
    }

    for _, tc := range testCases {
        solved := Solve(tc.Cube)

        if len(tc.Expected) != len(solved) {
            t.Errorf("Error while solving. Incompatible solution sizes.\nGot  :%s\nWant :%s\n", solved, tc.Expected)
            return
        }

        for i, move := range solved {
            if (move != tc.Expected[i]) {
                t.Errorf("Error while solving at move %d.\nGot  %s\nWant %s\n", i,  solved, tc.Expected)
            }
        }
    }
}
