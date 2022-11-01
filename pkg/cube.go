package rubikscube

const (
	top    = 0
	front  = 1
	bottom = 2
	back   = 3
	left   = 4
	right  = 5
)

type cube [6][][]int

func NewCube(size int) cube {
	cube := [6][][]int{}
	for i := 0; i < 6; i++ {
		cube[i] = make([][]int, size)
		for j := 0; j < size; j++ {
			cube[i][j] = make([]int, size)
			for k := 0; k < size; k++ {
				cube[i][j][k] = i
			}
		}
	}
	return cube
}

func NewTestCube(size int) cube {
	l, cube := 1, [6][][]int{}
	for i := 0; i < 6; i++ {
		cube[i] = make([][]int, size)
		for j := 0; j < size; j++ {
			cube[i][j] = make([]int, size)
			for k := 0; k < size; k++ {
				cube[i][j][k] = l
				l += 1
			}
		}
	}
	return cube
}

func (cube cube) Size() int {
	return len(cube[0])
}

func (cube *cube) TurnUp() {
	cube.Swap([][]int{{front, top}, {front, back}, {front, bottom}})
	cube.RotateFacesLeft(left)
	cube.RotateFacesRight(back, back, right)
	cube.ReverseRows(back)
}

func (cube *cube) TurnDown() {
	cube.Swap([][]int{{front, bottom}, {front, back}, {front, top}})
	cube.RotateFacesLeft(right)
	cube.RotateFacesRight(top, top, left)
	cube.ReverseRows(top)
}

func (cube *cube) TurnLeft() {
	cube.Swap([][]int{{left, top}, {left, right}, {left, bottom}})
	cube.RotateFacesLeft(back)
	cube.RotateFacesRight(top, front, bottom, left, right)
}

func (cube *cube) TurnRight() {
	cube.Swap([][]int{{right, top}, {right, left}, {right, bottom}})
	cube.RotateFacesLeft(top, front, bottom, left, right)
	cube.RotateFacesRight(back)
}

func (cube *cube) RotateLeft() {
	cube.Swap([][]int{{front, left}, {front, back}, {front, right}})
	cube.RotateFacesLeft(bottom, back, back)
	cube.RotateFacesRight(top, right, right)
}

func (cube *cube) RotateRight() {
	cube.Swap([][]int{{front, right}, {front, back}, {front, left}})
	cube.RotateFacesLeft(top, left, left)
	cube.RotateFacesRight(bottom, back, back)
}

func (cube *cube) RotateColumnUp(column int) {
	for row := range cube[0] {
		cube[front][row][column], cube[top][row][column] = cube[top][row][column], cube[front][row][column]
		cube[front][row][column], cube[back][row][column] = cube[back][row][column], cube[front][row][column]
		cube[front][row][column], cube[bottom][row][column] = cube[bottom][row][column], cube[front][row][column]
	}

	if column == 0 {
		cube.RotateFaceLeft(left)
	}

	if column == cube.Size()-1 {
		cube.RotateFaceRight(right)
	}
}

func (cube *cube) RotateColumnDown(column int) {
	for row := range cube[0] {
		cube[front][row][column], cube[bottom][row][column] = cube[bottom][row][column], cube[front][row][column]
		cube[front][row][column], cube[back][row][column] = cube[back][row][column], cube[front][row][column]
		cube[front][row][column], cube[top][row][column] = cube[top][row][column], cube[front][row][column]
	}

	if column == 0 {
		cube.RotateFaceRight(left)
	}

	if column == cube.Size()-1 {
		cube.RotateFaceLeft(right)
	}
}

func (cube *cube) RotateRowLeft(row int) {
	opposite := cube.Size() - row - 1
	cube[front][row], cube[left][row] = cube[left][row], cube[front][row]
	cube[front][row], cube[back][opposite] = cube[back][opposite], cube[front][row]
	cube.ReverseRow(back, opposite)
	cube[front][row], cube[right][row] = cube[right][row], cube[front][row]
	cube.ReverseRow(right, row)

	if row == 0 {
		cube.RotateFaceRight(top)
	}

	if row == cube.Size()-1 {
		cube.RotateFaceLeft(bottom)
	}
}

func (cube *cube) RotateRowRight(row int) {
	opposite := cube.Size() - row - 1
	cube[front][row], cube[right][row] = cube[right][row], cube[front][row]
	cube[front][row], cube[back][opposite] = cube[back][opposite], cube[front][row]
	cube.ReverseRow(back, opposite)
	cube[front][row], cube[left][row] = cube[left][row], cube[front][row]
	cube.ReverseRow(left, row)

	if row == 0 {
		cube.RotateFaceLeft(top)
	}

	if row == cube.Size()-1 {
		cube.RotateFaceRight(bottom)
	}
}

func (cube *cube) RotateFacesLeft(faces ...int) {
	for _, face := range faces {
		cube.RotateFaceLeft(face)
	}
}

func (cube *cube) RotateFaceLeft(face int) {
	surface, iterations := cube[face], len(cube[face])/2
	for size := cube.Size() - 1; iterations != 0; size-- {
		iterations -= 1

		skip := cube.Size() - (size + 1)
		for i := 0; i < size-skip; i++ {
			surface[skip][skip+i], surface[size-i][skip] = surface[size-i][skip], surface[skip][skip+i]
			surface[skip][skip+i], surface[size][size-i] = surface[size][size-i], surface[skip][skip+i]
			surface[skip][skip+i], surface[skip+i][size] = surface[skip+i][size], surface[skip][skip+i]
		}
	}
}

func (cube *cube) RotateFacesRight(faces ...int) {
	for _, face := range faces {
		cube.RotateFaceRight(face)
	}
}

func (cube *cube) RotateFaceRight(face int) {
	surface, iterations := cube[face], len(cube[face])/2
	for size := cube.Size() - 1; iterations != 0; size-- {
		iterations -= 1

		skip := cube.Size() - (size + 1)
		for i := 0; i < size-skip; i++ {
			surface[skip][skip+i], surface[skip+i][size] = surface[skip+i][size], surface[skip][skip+i]
			surface[skip][skip+i], surface[size][size-i] = surface[size][size-i], surface[skip][skip+i]
			surface[skip][skip+i], surface[size-i][skip] = surface[size-i][skip], surface[skip][skip+i]
		}
	}
}

func (cube *cube) Swap(faces [][]int) {
	for _, pair := range faces {
		cube[pair[0]], cube[pair[1]] = cube[pair[1]], cube[pair[0]]
	}
}

func (cube *cube) ReverseRows(face int) {
	for i := range cube[face] {
		cube.ReverseRow(face, i)
	}
}

func (cube *cube) ReverseRow(face, row int) {
	vector := cube[face][row]
	for i, j := 0, len(vector)-1; i < j; i, j = i+1, j-1 {
		vector[i], vector[j] = vector[j], vector[i]
	}
}
