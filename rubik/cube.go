package rubik

import "encoding/binary"

type Cube struct {
	up    []byte // RED
	down  []byte // ORANGE
	front []byte // WHITE
	back  []byte // YELLOW
	left  []byte // BLUE
	right []byte // GREEN
}

type Hcube struct {
	up    [3][3]string
	down  [3][3]string
	front [3][3]string
	back  [3][3]string
	left  [3][3]string
	right [3][3]string
}

var Color = map[int8]string{
	0x01: "W", // WHITE
	0x02: "Y", // YELLOW
	0x03: "R", // RED
	0x04: "O", // ORANGE
	0x05: "B", // BLUE
	0x06: "G", // GREEN
}

func CubeToString(cube Cube) string {
	var str string = string(cube.right) + string(cube.left) + string(cube.down) + string(cube.back) + string(cube.front) + string(cube.up)
	return str
}

func InstanceCube() Cube {
	cube := Cube{}
	cube.up = make([]byte, 8)
	cube.down = make([]byte, 8)
	cube.left = make([]byte, 8)
	cube.right = make([]byte, 8)
	cube.front = make([]byte, 8)
	cube.back = make([]byte, 8)
	return cube
}

func InitCube() Cube {
	cube := InstanceCube()
	binary.LittleEndian.PutUint64(cube.up, 0x0303030303030303)
	binary.LittleEndian.PutUint64(cube.down, 0x0404040404040404)
	binary.LittleEndian.PutUint64(cube.front, 0x0101010101010101)
	binary.LittleEndian.PutUint64(cube.back, 0x0202020202020202)
	binary.LittleEndian.PutUint64(cube.left, 0x0505050505050505)
	binary.LittleEndian.PutUint64(cube.right, 0x0606060606060606)

	return cube
}

func CreateCube(sequence []string, generator bool, nb_iter int) Cube {
	// sequence : list of moves given by the user
	// generator : use generator or not
	// shuffle : given number of iteration if generator is true
	start_cube := InitCube()
	if generator {
		GenerateShuffle(start_cube, nb_iter)
	} else {
		SequenceShuffle(start_cube, sequence, false)
	}
	return start_cube
}
