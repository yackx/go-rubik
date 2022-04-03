package rubik_alg

import (
	"reflect"
	"testing"

	"github.com/hultan/go-rubik/src/rubik"
)

func Test_performAlg(t *testing.T) {
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
			name: "Aa permutation",
			args: args{rubik.NewSolvedCube(), PllPermAa},
			want: rubik.NewCube("wwwwwwwww ggrgggggg brbrrrrrr obgbbbbbb roooooooo yyyyyyyyy"),
		},
		{
			name: "Ab permutation",
			args: args{rubik.NewSolvedCube(), PllPermAb},
			want: rubik.NewCube("wwwwwwwww ggbgggggg orgrrrrrr rbrbbbbbb boooooooo yyyyyyyyy"),
		},
		{
			name: "E permutation",
			args: args{rubik.NewSolvedCube(), PllPermE},
			want: rubik.NewCube("wwwwwwwww rgogggggg grbrrrrrr obrbbbbbb bogoooooo yyyyyyyyy"),
		},
		{
			name: "F permutation",
			args: args{rubik.NewSolvedCube(), PllPermF},
			want: rubik.NewCube("wwwwwwwww gbrgggggg brgrrrrrr rgbbbbbbb ooooooooo yyyyyyyyy"),
		},
		{
			name: "H permutation",
			args: args{rubik.NewSolvedCube(), PllPermH},
			want: rubik.NewCube("wwwwwwwww gbggggggg rorrrrrrr bgbbbbbbb orooooooo yyyyyyyyy"),
		},
		{
			name: "Z permutation",
			args: args{rubik.NewSolvedCube(), PllPermZ},
			want: rubik.NewCube("wwwwwwwww grggggggg rgrrrrrrr bobbbbbbb obooooooo yyyyyyyyy"),
		},
		{
			name: "T permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermT},
			want: rubik.NewCube("wwwwwwwww ggrgggggg bogrrrrrr rbbbbbbbb orooooooo yyyyyyyyy"),
		},
		{
			name: "Ja permutation",
			args: args{rubik.NewSolvedCube(), PllPermJa},
			want: rubik.NewCube("wwwwwwwww ggggggggg rrbrrrrrr oorbbbbbb bbooooooo yyyyyyyyy"),
		},
		{
			name: "Jb permutation",
			args: args{rubik.NewSolvedCube(), PllPermJb},
			want: rubik.NewCube("wwwwwwwww grrgggggg bggrrrrrr rbbbbbbbb ooooooooo yyyyyyyyy"),
		},
		{
			name: "Na permutation",
			args: args{rubik.NewSolvedCube(), PllPermNa},
			want: rubik.NewCube("wwwwwwwww gbbgggggg orrrrrrrr bggbbbbbb roooooooo yyyyyyyyy"),
		},
		{
			name: "Nb permutation",
			args: args{rubik.NewSolvedCube(), PllPermNb},
			want: rubik.NewCube("wwwwwwwww bbggggggg rrorrrrrr ggbbbbbbb ooroooooo yyyyyyyyy"),
		},
		{
			name: "Ra permutation",
			args: args{rubik.NewSolvedCube(), PllPermRa},
			want: rubik.NewCube("wwwwwwggw rrrgggwwg bbbyrryrr ooobbbbbb ggyooooog rroyyyyyy"),
		},
		{
			name: "Rb permutation",
			args: args{rubik.NewSolvedCube(), PllPermRb},
			want: rubik.NewCube("gwywwwwbb grwgggggg rbbrrrrro rgrwbbbbb woooooooo yyyyyyyyw"),
		},
		{
			name: "Ua permutation",
			args: args{rubik.NewSolvedCube(), PllPermUa},
			want: rubik.NewCube("wwwwwwwww goggggggg rgrrrrrrr bbbbbbbbb orooooooo yyyyyyyyy"),
		},
		{
			name: "Ub permutation",
			args: args{rubik.NewSolvedCube(), PllPermUb},
			want: rubik.NewCube("wwwwwwwww grggggggg rorrrrrrr bbbbbbbbb ogooooooo yyyyyyyyy"),
		},
		{
			name: "V permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermV},
			want: rubik.NewCube("wwwwwwwww obrrrrrrr brgbbbbbb roooooooo ggbgggggg yyyyyyyyy"),
		},
		{
			name: "Y permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermY},
			want: rubik.NewCube("wwwwwwwww ggbgggggg orrrrrrrr bogbbbbbb rbooooooo yyyyyyyyy"),
		},
		{
			name: "Ga permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermGa},
			want: rubik.NewCube("wwwwwwwww gbroooooo bgggggggg robrrrrrr orobbbbbb yyyyyyyyy"),
		},
		{
			name: "Gb permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermGb},
			want: rubik.NewCube("wwwwwwwww obgrrrrrr rgrbbbbbb boooooooo grbgggggg yyyyyyyyy"),
		},
		{
			name: "Gc permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermGc},
			want: rubik.NewCube("wwwwwwwww obgrrrrrr rorbbbbbb brooooooo ggbgggggg yyyyyyyyy"),
		},
		{
			name: "Gd permutation",
			args: args{cube: rubik.NewSolvedCube(), alg: PllPermGd},
			want: rubik.NewCube("wwwwwwwww gbroooooo boggggggg rrbrrrrrr ogobbbbbb yyyyyyyyy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteAlg(tt.args.cube, tt.args.alg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecuteAlg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseAlg(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"T permutation", args{PllPermT}, "F R U' R' U R U R2 F' R U R U' R'"},
		{"Aa permutation", args{PllPermAa}, "x R2 D2 R U R' D2 R U' R x'"},
		{"H permutation", args{PllPermH}, "M2 U' M2 U2 M2 U' M2"},
		{"Ua permutation", args{PllPermUa}, "R2 U R U R' U' R' U' R' U R'"},
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
		{
			name: "Ja perm",
			args: args{alg: "     (R' U L' U2) (R U' R' U2 R) L [U'] "},
			want: "R' U L' U2 R U' R' U2 R L U'",
		},
		{
			name: "Rb perm",
			args: args{alg: " (R' U2 R U2') R' F (R U R' U') R' F' R2 [U'] "},
			want: "R' U2 R U2' R' F R U R' U' R' F' R2 U'",
		},
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
			name: "R move",
			args: args{rubik.NewSolvedCube(), "R"},
			want: rubik.NewCube("wwgwwgwwg ggyggyggy rrrrrrrrr wbbwbbwbb ooooooooo yybyybyyb"),
		},
		{
			name: "R2 move",
			args: args{rubik.NewSolvedCube(), "R2"},
			want: rubik.NewCube("wwywwywwy ggbggbggb rrrrrrrrr gbbgbbgbb ooooooooo yywyywyyw"),
		},
		{
			name: "R' move",
			args: args{rubik.NewSolvedCube(), "R'"},
			want: rubik.NewCube("wwbwwbwwb ggwggwggw rrrrrrrrr ybbybbybb ooooooooo yygyygyyg"),
		},
		{
			name: "L move",
			args: args{rubik.NewSolvedCube(), "L"},
			want: rubik.NewCube("bwwbwwbww wggwggwgg rrrrrrrrr bbybbybby ooooooooo gyygyygyy"),
		},
		{
			name: "L2 move",
			args: args{rubik.NewSolvedCube(), "L2"},
			want: rubik.NewCube("ywwywwyww bggbggbgg rrrrrrrrr bbgbbgbbg ooooooooo wyywyywyy"),
		},
		{
			name: "L' move",
			args: args{rubik.NewSolvedCube(), "L'"},
			want: rubik.NewCube("gwwgwwgww yggyggygg rrrrrrrrr bbwbbwbbw ooooooooo byybyybyy"),
		},
		{
			name: "U move",
			args: args{rubik.NewSolvedCube(), "U"},
			want: rubik.NewCube("wwwwwwwww rrrgggggg bbbrrrrrr ooobbbbbb gggoooooo yyyyyyyyy"),
		},
		{
			name: "U2 move",
			args: args{rubik.NewSolvedCube(), "U2"},
			want: rubik.NewCube("wwwwwwwww bbbgggggg ooorrrrrr gggbbbbbb rrroooooo yyyyyyyyy"),
		},
		{
			name: "U' move",
			args: args{rubik.NewSolvedCube(), "U'"},
			want: rubik.NewCube("wwwwwwwww ooogggggg gggrrrrrr rrrbbbbbb bbboooooo yyyyyyyyy"),
		},
		{
			name: "D move",
			args: args{rubik.NewSolvedCube(), "D"},
			want: rubik.NewCube("wwwwwwwww ggggggooo rrrrrrggg bbbbbbrrr oooooobbb yyyyyyyyy"),
		},
		{
			name: "D2 move",
			args: args{rubik.NewSolvedCube(), "D2"},
			want: rubik.NewCube("wwwwwwwww ggggggbbb rrrrrrooo bbbbbbggg oooooorrr yyyyyyyyy"),
		},
		{
			name: "D' move",
			args: args{rubik.NewSolvedCube(), "D'"},
			want: rubik.NewCube("wwwwwwwww ggggggrrr rrrrrrbbb bbbbbbooo ooooooggg yyyyyyyyy"),
		},
		{
			name: "F move",
			args: args{rubik.NewSolvedCube(), "F"},
			want: rubik.NewCube("wwwwwwooo ggggggggg wrrwrrwrr bbbbbbbbb ooyooyooy rrryyyyyy"),
		},
		{
			name: "F2 move",
			args: args{rubik.NewSolvedCube(), "F2"},
			want: rubik.NewCube("wwwwwwyyy ggggggggg orrorrorr bbbbbbbbb oorooroor wwwyyyyyy"),
		},
		{
			name: "F' move",
			args: args{rubik.NewSolvedCube(), "F'"},
			want: rubik.NewCube("wwwwwwrrr ggggggggg yrryrryrr bbbbbbbbb oowoowoow oooyyyyyy"),
		},
		{
			name: "B move",
			args: args{rubik.NewSolvedCube(), "B"},
			want: rubik.NewCube("rrrwwwwww ggggggggg rryrryrry bbbbbbbbb woowoowoo yyyyyyooo"),
		},
		{
			name: "B2 move",
			args: args{rubik.NewSolvedCube(), "B2"},
			want: rubik.NewCube("yyywwwwww ggggggggg rrorrorro bbbbbbbbb roorooroo yyyyyywww"),
		},
		{
			name: "B' move",
			args: args{rubik.NewSolvedCube(), "B'"},
			want: rubik.NewCube("ooowwwwww ggggggggg rrwrrwrrw bbbbbbbbb yooyooyoo yyyyyyrrr"),
		},
		{
			name: "M move",
			args: args{rubik.NewSolvedCube(), "M"},
			want: rubik.NewCube("wbwwbwwbw gwggwggwg rrrrrrrrr bybbybbyb ooooooooo ygyygyygy"),
		},
		{
			name: "M2 move",
			args: args{rubik.NewSolvedCube(), "M2"},
			want: rubik.NewCube("wywwywwyw gbggbggbg rrrrrrrrr bgbbgbbgb ooooooooo ywyywyywy"),
		},
		{
			name: "M' move",
			args: args{rubik.NewSolvedCube(), "M'"},
			want: rubik.NewCube("wgwwgwwgw gyggyggyg rrrrrrrrr bwbbwbbwb ooooooooo ybyybyyby"),
		},
		{
			name: "E move",
			args: args{rubik.NewSolvedCube(), "E"},
			want: rubik.NewCube("wwwwwwwww gggoooggg rrrgggrrr bbbrrrbbb ooobbbooo yyyyyyyyy"),
		},
		{
			name: "E2 move",
			args: args{rubik.NewSolvedCube(), "E2"},
			want: rubik.NewCube("wwwwwwwww gggbbbggg rrrooorrr bbbgggbbb ooorrrooo yyyyyyyyy"),
		},
		{
			name: "E' move",
			args: args{rubik.NewSolvedCube(), "E'"},
			want: rubik.NewCube("wwwwwwwww gggrrrggg rrrbbbrrr bbbooobbb ooogggooo yyyyyyyyy"),
		},
		{
			name: "S move",
			args: args{rubik.NewSolvedCube(), "S"},
			want: rubik.NewCube("wwwooowww ggggggggg rwrrwrrwr bbbbbbbbb oyooyooyo yyyrrryyy"),
		},
		{
			name: "S2 move",
			args: args{rubik.NewSolvedCube(), "S2"},
			want: rubik.NewCube("wwwyyywww ggggggggg rorrorror bbbbbbbbb oroorooro yyywwwyyy"),
		},
		{
			name: "S' move",
			args: args{rubik.NewSolvedCube(), "S'"},
			want: rubik.NewCube("wwwrrrwww ggggggggg ryrryrryr bbbbbbbbb owoowoowo yyyoooyyy"),
		},
		{
			name: "r move",
			args: args{rubik.NewSolvedCube(), "r"},
			want: rubik.NewCube("wggwggwgg gyygyygyy rrrrrrrrr wwbwwbwwb ooooooooo ybbybbybb"),
		},
		{
			name: "r2 move",
			args: args{rubik.NewSolvedCube(), "r2"},
			want: rubik.NewCube("wyywyywyy gbbgbbgbb rrrrrrrrr ggbggbggb ooooooooo ywwywwyww"),
		},
		{
			name: "r' move",
			args: args{rubik.NewSolvedCube(), "r'"},
			want: rubik.NewCube("wbbwbbwbb gwwgwwgww rrrrrrrrr yybyybyyb ooooooooo yggyggygg"),
		},
		{
			name: "l move",
			args: args{rubik.NewSolvedCube(), "l"},
			want: rubik.NewCube("bbwbbwbbw wwgwwgwwg rrrrrrrrr byybyybyy ooooooooo ggyggyggy"),
		},
		{
			name: "l2 move",
			args: args{rubik.NewSolvedCube(), "l2"},
			want: rubik.NewCube("yywyywyyw bbgbbgbbg rrrrrrrrr bggbggbgg ooooooooo wwywwywwy"),
		},
		{
			name: "l' move",
			args: args{rubik.NewSolvedCube(), "l'"},
			want: rubik.NewCube("ggwggwggw yygyygyyg rrrrrrrrr bwwbwwbww ooooooooo bbybbybby"),
		},
		{
			name: "u move",
			args: args{rubik.NewSolvedCube(), "u"},
			want: rubik.NewCube("wwwwwwwww rrrrrrggg bbbbbbrrr oooooobbb ggggggooo yyyyyyyyy"),
		},
		{
			name: "u2 move",
			args: args{rubik.NewSolvedCube(), "u2"},
			want: rubik.NewCube("wwwwwwwww bbbbbbggg oooooorrr ggggggbbb rrrrrrooo yyyyyyyyy"),
		},
		{
			name: "u' move",
			args: args{rubik.NewSolvedCube(), "u'"},
			want: rubik.NewCube("wwwwwwwww ooooooggg ggggggrrr rrrrrrbbb bbbbbbooo yyyyyyyyy"),
		},
		{
			name: "d move",
			args: args{rubik.NewSolvedCube(), "d"},
			want: rubik.NewCube("wwwwwwwww gggoooooo rrrgggggg bbbrrrrrr ooobbbbbb yyyyyyyyy"),
		},
		{
			name: "d2 move",
			args: args{rubik.NewSolvedCube(), "d2"},
			want: rubik.NewCube("wwwwwwwww gggbbbbbb rrroooooo bbbgggggg ooorrrrrr yyyyyyyyy"),
		},
		{
			name: "d' move",
			args: args{rubik.NewSolvedCube(), "d'"},
			want: rubik.NewCube("wwwwwwwww gggrrrrrr rrrbbbbbb bbboooooo ooogggggg yyyyyyyyy"),
		},
		{
			name: "f move",
			args: args{rubik.NewSolvedCube(), "f"},
			want: rubik.NewCube("wwwoooooo ggggggggg wwrwwrwwr bbbbbbbbb oyyoyyoyy rrrrrryyy"),
		},
		{
			name: "f2 move",
			args: args{rubik.NewSolvedCube(), "f2"},
			want: rubik.NewCube("wwwyyyyyy ggggggggg oorooroor bbbbbbbbb orrorrorr wwwwwwyyy"),
		},
		{
			name: "f' move",
			args: args{rubik.NewSolvedCube(), "f'"},
			want: rubik.NewCube("wwwrrrrrr ggggggggg yyryyryyr bbbbbbbbb owwowwoww ooooooyyy"),
		},
		{
			name: "b move",
			args: args{rubik.NewSolvedCube(), "b"},
			want: rubik.NewCube("rrrrrrwww ggggggggg ryyryyryy bbbbbbbbb wwowwowwo yyyoooooo"),
		},
		{
			name: "b2 move",
			args: args{rubik.NewSolvedCube(), "b2"},
			want: rubik.NewCube("yyyyyywww ggggggggg roorooroo bbbbbbbbb rrorrorro yyywwwwww"),
		},
		{
			name: "b' move",
			args: args{rubik.NewSolvedCube(), "b'"},
			want: rubik.NewCube("oooooowww ggggggggg rwwrwwrww bbbbbbbbb yyoyyoyyo yyyrrrrrr"),
		},
		{
			name: "x rotation",
			args: args{rubik.NewSolvedCube(), "x"},
			want: rubik.NewCube("ggggggggg yyyyyyyyy rrrrrrrrr wwwwwwwww ooooooooo bbbbbbbbb"),
		},
		{
			name: "x2 rotation",
			args: args{rubik.NewSolvedCube(), "x2"},
			want: rubik.NewCube("yyyyyyyyy bbbbbbbbb rrrrrrrrr ggggggggg ooooooooo wwwwwwwww"),
		},
		{
			name: "x' rotation",
			args: args{rubik.NewSolvedCube(), "x'"},
			want: rubik.NewCube("bbbbbbbbb wwwwwwwww rrrrrrrrr yyyyyyyyy ooooooooo ggggggggg"),
		},
		{
			name: "y rotation",
			args: args{rubik.NewSolvedCube(), "y"},
			want: rubik.NewCube("wwwwwwwww rrrrrrrrr bbbbbbbbb ooooooooo ggggggggg yyyyyyyyy"),
		},
		{
			name: "y2 rotation",
			args: args{rubik.NewSolvedCube(), "y2"},
			want: rubik.NewCube("wwwwwwwww bbbbbbbbb ooooooooo ggggggggg rrrrrrrrr yyyyyyyyy"),
		},
		{
			name: "y' rotation",
			args: args{rubik.NewSolvedCube(), "y'"},
			want: rubik.NewCube("wwwwwwwww ooooooooo ggggggggg rrrrrrrrr bbbbbbbbb yyyyyyyyy"),
		},
		{
			name: "z rotation",
			args: args{rubik.NewSolvedCube(), "z"},
			want: rubik.NewCube("ooooooooo ggggggggg wwwwwwwww bbbbbbbbb yyyyyyyyy rrrrrrrrr"),
		},
		{
			name: "z2 rotation",
			args: args{rubik.NewSolvedCube(), "z2"},
			want: rubik.NewCube("yyyyyyyyy ggggggggg ooooooooo bbbbbbbbb rrrrrrrrr wwwwwwwww"),
		},
		{
			name: "z' rotation",
			args: args{rubik.NewSolvedCube(), "z'"},
			want: rubik.NewCube("rrrrrrrrr ggggggggg yyyyyyyyy bbbbbbbbb wwwwwwwww ooooooooo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeMove(tt.args.cube, tt.args.move); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeMove() = %v, want %v", got, tt.want)
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
		{name: "R", args: args{"R"}, want: "R'"},
		{name: "R2", args: args{"R2"}, want: "R2"},
		{name: "R'", args: args{"R'"}, want: "R"},
		{name: "L", args: args{"L"}, want: "L'"},
		{name: "L2", args: args{"L2"}, want: "L2"},
		{name: "L'", args: args{"L'"}, want: "L"},
		{name: "U", args: args{"U"}, want: "U'"},
		{name: "U2", args: args{"U2"}, want: "U2"},
		{name: "U'", args: args{"U'"}, want: "U"},
		{name: "D", args: args{"D"}, want: "D'"},
		{name: "D2", args: args{"D2"}, want: "D2"},
		{name: "D'", args: args{"D'"}, want: "D"},
		{name: "F", args: args{"F"}, want: "F'"},
		{name: "F2", args: args{"F2"}, want: "F2"},
		{name: "F'", args: args{"F'"}, want: "F"},
		{name: "B", args: args{"B"}, want: "B'"},
		{name: "B2", args: args{"B2"}, want: "B2"},
		{name: "B'", args: args{"B'"}, want: "B"},
		{name: "M", args: args{"M"}, want: "M'"},
		{name: "M2", args: args{"M2"}, want: "M2"},
		{name: "M'", args: args{"M'"}, want: "M"},
		{name: "E", args: args{"E"}, want: "E'"},
		{name: "E2", args: args{"E2"}, want: "E2"},
		{name: "E'", args: args{"E'"}, want: "E"},
		{name: "S", args: args{"S"}, want: "S'"},
		{name: "S2", args: args{"S2"}, want: "S2"},
		{name: "S'", args: args{"S'"}, want: "S"},
		{name: "u", args: args{"u"}, want: "u'"},
		{name: "u2", args: args{"u2"}, want: "u2"},
		{name: "u'", args: args{"u'"}, want: "u"},
		{name: "d", args: args{"d"}, want: "d'"},
		{name: "d2", args: args{"d2"}, want: "d2"},
		{name: "d'", args: args{"d'"}, want: "d"},
		{name: "f", args: args{"f"}, want: "f'"},
		{name: "f2", args: args{"f2"}, want: "f2"},
		{name: "f'", args: args{"f'"}, want: "f"},
		{name: "b", args: args{"b"}, want: "b'"},
		{name: "b2", args: args{"b2"}, want: "b2"},
		{name: "b'", args: args{"b'"}, want: "b"},
		{name: "r", args: args{"r"}, want: "r'"},
		{name: "r2", args: args{"r2"}, want: "r2"},
		{name: "r'", args: args{"r'"}, want: "r"},
		{name: "l", args: args{"l"}, want: "l'"},
		{name: "l2", args: args{"l2"}, want: "l2"},
		{name: "l'", args: args{"l'"}, want: "l"},
		{name: "x", args: args{"x"}, want: "x'"},
		{name: "x2", args: args{"x2"}, want: "x2"},
		{name: "x'", args: args{"x'"}, want: "x"},
		{name: "y", args: args{"y"}, want: "y'"},
		{name: "y2", args: args{"y2"}, want: "y2"},
		{name: "y'", args: args{"y'"}, want: "y"},
		{name: "z", args: args{"z"}, want: "z'"},
		{name: "z2", args: args{"z2"}, want: "z2"},
		{name: "z'", args: args{"z'"}, want: "z"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseMove(tt.args.move); got != tt.want {
				t.Errorf("reverseMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomScrambles(t *testing.T) {
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
			name: "Random scramble 1",
			args: args{rubik.NewSolvedCube(), "U F' B D R L' B R2 B U2 D2 L B' R L F L' F' L B2 D2 U F2 D R"},
			want: rubik.NewCube("oobgwboyg ggwogrgrr oryyrwybb oywgbrybb boywowrbr wggyyowwr"),
		},
		{
			name: "Random scramble 2",
			args: args{rubik.NewSolvedCube(), "D2 L2 F' B' D2 F2 U' D2 L R' F U2 F' B R' D' U L R' B2 R' L' D2 F2 U"},
			want: rubik.NewCube("oywbwbroy yboggrbrb bygwrywwr rrygbobbg gwggowoyw rgooyowry"),
		},
		{
			name: "Random scramble 3",
			args: args{rubik.NewSolvedCube(), "L U2 L2 F' L' D' F L D2 L' U F2 B L' B2 U' D L R2 U' R L' U F2 R"},
			want: rubik.NewCube("wowgwowwy obrwgybwy gygbroory rgrbbwryg bybgoooby orgrygwrb"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteAlg(tt.args.cube, tt.args.alg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecuteAlg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetScrambledCube(t *testing.T) {
	for i := 0; i < 10; i++ {
		scrambled, alg := GetScrambledCube()
		if scrambled.IsSolved() {
			t.Error("Scrambled cube should not be solved : ", alg)
		}
	}
}
