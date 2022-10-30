package rubikscube

import (
	"testing"
)

func Test2By2(t *testing.T) {
	got := NewCube(2)

	expected := [6][][]int{
		{
			{0, 0},
			{0, 0},
		},
		{
			{1, 1},
			{1, 1},
		},
		{
			{2, 2},
			{2, 2},
		},
		{
			{3, 3},
			{3, 3},
		},
		{
			{4, 4},
			{4, 4},
		},
		{
			{5, 5},
			{5, 5},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestTurnUp(t *testing.T) {
	got := NewTestCube(2)
	got.TurnUp()

	expected := [6][][]int{
		{
			{5, 6},
			{7, 8},
		},
		{
			{9, 10},
			{11, 12},
		},
		{
			{13, 14},
			{15, 16},
		},
		{
			{3, 4},
			{1, 2},
		},
		{
			{18, 20},
			{17, 19},
		},
		{
			{23, 21},
			{24, 22},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestTurnDown(t *testing.T) {
	got := NewTestCube(2)
	got.TurnDown()

	expected := [6][][]int{
		{
			{15, 16},
			{13, 14},
		},
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
		{
			{9, 10},
			{11, 12},
		},
		{
			{19, 17},
			{20, 18},
		},
		{
			{22, 24},
			{21, 23},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestTurnLeft(t *testing.T) {
	got := NewTestCube(2)
	got.TurnLeft()

	expected := [6][][]int{
		{
			{19, 17},
			{20, 18},
		},
		{
			{7, 5},
			{8, 6},
		},
		{
			{23, 21},
			{24, 22},
		},
		{
			{14, 16},
			{13, 15},
		},
		{
			{11, 9},
			{12, 10},
		},
		{
			{3, 1},
			{4, 2},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestTurnright(t *testing.T) {
	got := NewTestCube(2)
	got.TurnRight()

	expected := [6][][]int{
		{
			{22, 24},
			{21, 23},
		},
		{
			{6, 8},
			{5, 7},
		},
		{
			{18, 20},
			{17, 19},
		},
		{
			{15, 13},
			{16, 14},
		},
		{
			{2, 4},
			{1, 3},
		},
		{
			{10, 12},
			{9, 11},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateLeft(t *testing.T) {
	got := NewTestCube(2)
	got.RotateLeft()

	expected := [6][][]int{
		{
			{3, 1},
			{4, 2},
		},
		{
			{21, 22},
			{23, 24},
		},
		{
			{10, 12},
			{9, 11},
		},
		{
			{20, 19},
			{18, 17},
		},
		{
			{5, 6},
			{7, 8},
		},
		{
			{16, 15},
			{14, 13},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateRight(t *testing.T) {
	got := NewTestCube(2)
	got.RotateRight()

	expected := [6][][]int{
		{
			{2, 4},
			{1, 3},
		},
		{
			{17, 18},
			{19, 20},
		},
		{
			{11, 9},
			{12, 10},
		},
		{
			{24, 23},
			{22, 21},
		},
		{
			{16, 15},
			{14, 13},
		},
		{
			{5, 6},
			{7, 8},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateLeftColumnUp(t *testing.T) {
	got := NewTestCube(2)
	got.RotateColumnUp(0)

	expected := [6][][]int{
		{
			{5, 2},
			{7, 4},
		},
		{
			{9, 6},
			{11, 8},
		},
		{
			{13, 10},
			{15, 12},
		},
		{
			{1, 14},
			{3, 16},
		},
		{
			{18, 20},
			{17, 19},
		},
		{
			{21, 22},
			{23, 24},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateRightColumnUp(t *testing.T) {
	got := NewTestCube(2)
	got.RotateColumnUp(1)

	expected := [6][][]int{
		{
			{1, 6},
			{3, 8},
		},
		{
			{5, 10},
			{7, 12},
		},
		{
			{9, 14},
			{11, 16},
		},
		{
			{13, 2},
			{15, 4},
		},
		{
			{17, 18},
			{19, 20},
		},
		{
			{23, 21},
			{24, 22},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateLeftColumnDown(t *testing.T) {
	got := NewTestCube(2)
	got.RotateColumnDown(0)

	expected := [6][][]int{
		{
			{13, 2},
			{15, 4},
		},
		{
			{1, 6},
			{3, 8},
		},
		{
			{5, 10},
			{7, 12},
		},
		{
			{9, 14},
			{11, 16},
		},
		{
			{19, 17},
			{20, 18},
		},
		{
			{21, 22},
			{23, 24},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateRightColumnDown(t *testing.T) {
	got := NewTestCube(2)
	got.RotateColumnDown(1)

	expected := [6][][]int{
		{
			{1, 14},
			{3, 16},
		},
		{
			{5, 2},
			{7, 4},
		},
		{
			{9, 6},
			{11, 8},
		},
		{
			{13, 10},
			{15, 12},
		},
		{
			{17, 18},
			{19, 20},
		},
		{
			{22, 24},
			{21, 23},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateTopRowLeft(t *testing.T) {
	got := NewTestCube(2)
	got.RotateRowLeft(0)

	expected := [6][][]int{
		{
			{3, 1},
			{4, 2},
		},
		{
			{21, 22},
			{7, 8},
		},
		{
			{9, 10},
			{11, 12},
		},
		{
			{13, 14},
			{18, 17},
		},
		{
			{5, 6},
			{19, 20},
		},
		{
			{16, 15},
			{23, 24},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateBottomRowLeft(t *testing.T) {
	got := NewTestCube(2)
	got.RotateRowLeft(1)

	expected := [6][][]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{23, 24},
		},
		{
			{10, 12},
			{9, 11},
		},
		{
			{20, 19},
			{15, 16},
		},
		{
			{17, 18},
			{7, 8},
		},
		{
			{21, 22},
			{14, 13},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateTopRowRight(t *testing.T) {
	got := NewTestCube(2)
	got.RotateRowRight(0)

	expected := [6][][]int{
		{
			{2, 4},
			{1, 3},
		},
		{
			{17, 18},
			{7, 8},
		},
		{
			{9, 10},
			{11, 12},
		},
		{
			{13, 14},
			{22, 21},
		},
		{
			{16, 15},
			{19, 20},
		},
		{
			{5, 6},
			{23, 24},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

func TestRotateBottomRowRight(t *testing.T) {
	got := NewTestCube(2)
	got.RotateRowRight(1)

	expected := [6][][]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{19, 20},
		},
		{
			{11, 9},
			{12, 10},
		},
		{
			{24, 23},
			{15, 16},
		},
		{
			{17, 18},
			{14, 13},
		},
		{
			{21, 22},
			{7, 8},
		},
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
				if gotSquare != expectedSquare {
					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
				}
			}
		}
	}
}

// func TestRotate3x3FaceRight(t *testing.T) {
// 	got := NewCube(3)
// 	got[front][0][0] = 1
// 	got[front][0][1] = 2
// 	got[front][0][2] = 3
// 	got[front][1][0] = 4
// 	got[front][1][1] = 5
// 	got[front][1][2] = 6
// 	got[front][2][0] = 7
// 	got[front][2][1] = 8
// 	got[front][2][2] = 9
// 	got.RotateFaceRight(front)

// 	expected := [6][][]int{
// 		{
// 			{7, 4, 1},
// 			{8, 5, 2},
// 			{9, 6, 3},
// 		},
// 		{
// 			{1, 1, 1},
// 			{1, 1, 1},
// 			{1, 1, 1},
// 		},
// 		{
// 			{2, 2, 2},
// 			{2, 2, 2},
// 			{2, 2, 2},
// 		},
// 		{
// 			{3, 3, 3},
// 			{3, 3, 3},
// 			{3, 3, 3},
// 		},
// 		{
// 			{4, 4, 4},
// 			{4, 4, 4},
// 			{4, 4, 4},
// 		},
// 		{
// 			{5, 5, 5},
// 			{5, 5, 5},
// 			{5, 5, 5},
// 		},
// 	}

// 	for i := 0; i < 6; i++ {
// 		for j := 0; j < 3; j++ {
// 			for k := 0; k < 3; k++ {
// 				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
// 				if gotSquare != expectedSquare {
// 					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
// 				}
// 			}
// 		}
// 	}
// }

// func TestRotate3x3FaceLeft(t *testing.T) {
// 	got := NewCube(3)
// 	got[front][0][0] = 1
// 	got[front][0][1] = 2
// 	got[front][0][2] = 3
// 	got[front][1][0] = 4
// 	got[front][1][1] = 5
// 	got[front][1][2] = 6
// 	got[front][2][0] = 7
// 	got[front][2][1] = 8
// 	got[front][2][2] = 9
// 	got.RotateFaceLeft(front)

// 	expected := [6][][]int{
// 		{
// 			{3, 6, 9},
// 			{2, 5, 8},
// 			{1, 4, 7},
// 		},
// 		{
// 			{1, 1, 1},
// 			{1, 1, 1},
// 			{1, 1, 1},
// 		},
// 		{
// 			{2, 2, 2},
// 			{2, 2, 2},
// 			{2, 2, 2},
// 		},
// 		{
// 			{3, 3, 3},
// 			{3, 3, 3},
// 			{3, 3, 3},
// 		},
// 		{
// 			{4, 4, 4},
// 			{4, 4, 4},
// 			{4, 4, 4},
// 		},
// 		{
// 			{5, 5, 5},
// 			{5, 5, 5},
// 			{5, 5, 5},
// 		},
// 	}

// 	for i := 0; i < 6; i++ {
// 		for j := 0; j < 3; j++ {
// 			for k := 0; k < 3; k++ {
// 				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
// 				if gotSquare != expectedSquare {
// 					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
// 				}
// 			}
// 		}
// 	}
// }

// func TestRotate4x4FaceLeft(t *testing.T) {
// 	got := NewCube(4)
// 	got[front][0][0] = 1
// 	got[front][0][1] = 2
// 	got[front][0][2] = 3
// 	got[front][0][3] = 4
// 	got[front][1][0] = 5
// 	got[front][1][1] = 6
// 	got[front][1][2] = 7
// 	got[front][1][3] = 8
// 	got[front][2][0] = 9
// 	got[front][2][1] = 10
// 	got[front][2][2] = 11
// 	got[front][2][3] = 12
// 	got[front][3][0] = 13
// 	got[front][3][1] = 14
// 	got[front][3][2] = 15
// 	got[front][3][3] = 16
// 	got.RotateFaceLeft(front)

// 	expected := [6][][]int{
// 		{
// 			{4, 8, 12, 16},
// 			{3, 7, 11, 15},
// 			{2, 6, 10, 14},
// 			{1, 5, 9, 13},
// 		},
// 		{
// 			{1, 1, 1, 1},
// 			{1, 1, 1, 1},
// 			{1, 1, 1, 1},
// 			{1, 1, 1, 1},
// 		},
// 		{
// 			{2, 2, 2, 2},
// 			{2, 2, 2, 2},
// 			{2, 2, 2, 2},
// 			{2, 2, 2, 2},
// 		},
// 		{
// 			{3, 3, 3, 3},
// 			{3, 3, 3, 3},
// 			{3, 3, 3, 3},
// 			{3, 3, 3, 3},
// 		},
// 		{
// 			{4, 4, 4, 4},
// 			{4, 4, 4, 4},
// 			{4, 4, 4, 4},
// 			{4, 4, 4, 4},
// 		},
// 		{
// 			{5, 5, 5, 5},
// 			{5, 5, 5, 5},
// 			{5, 5, 5, 5},
// 			{5, 5, 5, 5},
// 		},
// 	}

// 	for i := 0; i < 6; i++ {
// 		for j := 0; j < 4; j++ {
// 			for k := 0; k < 4; k++ {
// 				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
// 				if gotSquare != expectedSquare {
// 					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
// 				}
// 			}
// 		}
// 	}
// }

// func TestRotate4x4FaceRight(t *testing.T) {
// 	got := NewCube(4)
// 	got[front][0][0] = 1
// 	got[front][0][1] = 2
// 	got[front][0][2] = 3
// 	got[front][0][3] = 4
// 	got[front][1][0] = 5
// 	got[front][1][1] = 6
// 	got[front][1][2] = 7
// 	got[front][1][3] = 8
// 	got[front][2][0] = 9
// 	got[front][2][1] = 10
// 	got[front][2][2] = 11
// 	got[front][2][3] = 12
// 	got[front][3][0] = 13
// 	got[front][3][1] = 14
// 	got[front][3][2] = 15
// 	got[front][3][3] = 16
// 	got.RotateFaceRight(front)

// 	expected := [6][][]int{
// 		{
// 			{13, 9, 5, 1},
// 			{14, 10, 6, 2},
// 			{15, 11, 7, 3},
// 			{16, 12, 8, 4},
// 		},
// 		{
// 			{1, 1, 1, 1},
// 			{1, 1, 1, 1},
// 			{1, 1, 1, 1},
// 			{1, 1, 1, 1},
// 		},
// 		{
// 			{2, 2, 2, 2},
// 			{2, 2, 2, 2},
// 			{2, 2, 2, 2},
// 			{2, 2, 2, 2},
// 		},
// 		{
// 			{3, 3, 3, 3},
// 			{3, 3, 3, 3},
// 			{3, 3, 3, 3},
// 			{3, 3, 3, 3},
// 		},
// 		{
// 			{4, 4, 4, 4},
// 			{4, 4, 4, 4},
// 			{4, 4, 4, 4},
// 			{4, 4, 4, 4},
// 		},
// 		{
// 			{5, 5, 5, 5},
// 			{5, 5, 5, 5},
// 			{5, 5, 5, 5},
// 			{5, 5, 5, 5},
// 		},
// 	}

// 	for i := 0; i < 6; i++ {
// 		for j := 0; j < 4; j++ {
// 			for k := 0; k < 4; k++ {
// 				gotSquare, expectedSquare := got[i][j][k], expected[i][j][k]
// 				if gotSquare != expectedSquare {
// 					t.Errorf("got %v expected %v", gotSquare, expectedSquare)
// 				}
// 			}
// 		}
// 	}
// }

// func TestAllCombinationsPossible(t *testing.T) {
// 	actions := []func(cube cube){
// 		func(cube cube) { cube.RotateRowRight(0) },
// 		func(cube cube) { cube.RotateRowLeft(0) },
// 		func(cube cube) { cube.RotateRowRight(1) },
// 		func(cube cube) { cube.RotateRowLeft(1) },
// 		// func(cube cube) { cube.RotateColumnUp(0) },
// 		// func(cube cube) { cube.RotateColumnDown(0) },
// 		// func(cube cube) { cube.RotateColumnUp(1) },
// 		// func(cube cube) { cube.RotateColumnDown(1) },
// 	}

// 	states, track, combinations := [][24]int{}, map[[24]int]struct{}{}, 3674160

// 	initialState := flatten(NewCube(2))
// 	states, track[initialState] = append(states, initialState), struct{}{}

// 	start, end := -1, 0
// 	for start != end {
// 		start = len(track)
// 		for _, state := range states {
// 			next := [][24]int{}
// 			for _, action := range actions {
// 				view := reshape(state)
// 				action(view)
// 				flat := flatten(view)
// 				if _, ok := track[flat]; !ok {
// 					next = append(next, flat)
// 					track[flat] = struct{}{}
// 				}
// 			}
// 			states = next
// 		}
// 		end = len(track)
// 	}
// 	fmt.Println(end)
// 	fmt.Println(combinations)
// }

// func checkFaces(cube cube) [][24]int {
// 	faces := [][24]int{}
// 	faces = append(faces, checkRotations(cube)...)
// 	return faces
// }

// func checkRotations(cube cube) [][24]int {
// 	checks := [][24]int{}

// 	checks = append(checks, flatten(cube))
// 	cube.TurnUp()
// 	checks = append(checks, flatten(cube))
// 	cube.TurnUp()
// 	checks = append(checks, flatten(cube))
// 	cube.TurnUp()
// 	checks = append(checks, flatten(cube))
// 	cube.TurnUp()

// 	checks = append(checks, flatten(cube))
// 	cube.TurnRight()
// 	checks = append(checks, flatten(cube))
// 	cube.TurnRight()
// 	checks = append(checks, flatten(cube))
// 	cube.TurnRight()
// 	checks = append(checks, flatten(cube))
// 	cube.TurnRight()
// 	return checks
// }

// func flatten(cube cube) [24]int {
// 	flat := [24]int{}

// 	l := 0
// 	for i := range cube {
// 		for j := range cube[i] {
// 			for k := range cube[i][j] {
// 				flat[l] = cube[i][j][k]
// 				l += 1
// 			}
// 		}
// 	}
// 	return flat
// }

// func reshape(flat [24]int) cube {
// 	cube := NewCube(2)

// 	l := 0
// 	for i := range cube {
// 		for j := range cube[i] {
// 			for k := range cube[i][j] {
// 				cube[i][j][k] = flat[l]
// 				l += 1
// 			}
// 		}
// 	}
// 	return cube
// }
