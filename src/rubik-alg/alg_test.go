package rubik_alg

import (
	"reflect"
	"testing"

	"github.com/hultan/go-rubik/src/rubik"
)

func TestPerformAlg(t *testing.T) {
	type args struct {
		cube rubik.Cube
		alg  string
	}
	tests := []struct {
		name string
		args args
		want rubik.Cube
	}{
		{
			name: "T permutation",
			args: args{rubik.NewSolvedCube(), "(R U R' U') (R' F R2 U') R' U' (R U R' F')"},
			want: rubik.NewCube("wwwwwwwww ggrgggggg bogrrrrrr rbbbbbbbb orooooooo yyyyyyyyy"),
		},
		{
			name: "H permutation",
			args: args{rubik.NewSolvedCube(), "M2 U M2 U2 M2 U M2"},
			want: rubik.NewCube("wwwwwwwww gbggggggg rorrrrrrr bgbbbbbbb orooooooo yyyyyyyyy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PerformAlg(tt.args.cube, tt.args.alg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PerformAlg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseAlg(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"T permutation", args{"(R U R' U') (R' F R2 U') R' U' (R U R' F')"}, "F R U' R' U R U R2 F' R U R U' R'"},
		{"Aa permutation", args{"x R' U R' D2 R U' R' D2 R2"}, "R2 D2 R U R' D2 R U' R x'"},
		{"H permutation", args{"M2 U M2 U2 M2 U M2"}, "M2 U' M2 U2 M2 U' M2"},
		{"Ua permutation", args{"y2 [R U' R U] R U R U' R' U' R2"}, "R2 U R U R' U' R' U' R' U R' y2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseAlg(tt.args.alg); got != tt.want {
				t.Errorf("ReverseAlg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanAlg(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanAlg(tt.args.alg); got != tt.want {
				t.Errorf("cleanAlg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_performMove(t *testing.T) {
	type args struct {
		cube rubik.Cube
		move string
	}
	tests := []struct {
		name string
		args args
		want rubik.Cube
	}{
		{
			name: "F move",
			args: args{rubik.NewSolvedCube(), "F"},
			want: rubik.NewCube("wwwwwwooo ggggggggg wrrwrrwrr bbbbbbbbb ooyooyooy rrryyyyyy"),
		},
		{
			name: "M move",
			args: args{rubik.NewSolvedCube(), "M"},
			want: rubik.NewCube("wbwwbwwbw gwggwggwg rrrrrrrrr bybbybbyb ooooooooo ygyygyygy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := performMove(tt.args.cube, tt.args.move); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("performMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseMove(t *testing.T) {
	type args struct {
		move string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseMove(tt.args.move); got != tt.want {
				t.Errorf("reverseMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
