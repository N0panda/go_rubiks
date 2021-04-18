package rubik

import (
	"encoding/binary"
	"math/bits"
)

func RotateFace(face *[]byte, rotation int) {
	tmp_face := bits.RotateLeft64(uint64(binary.LittleEndian.Uint64(*face)), rotation)
	binary.LittleEndian.PutUint64(*face, tmp_face)
}

func F90(cube Cube) {

	// Rotation de la face
	RotateFace(&cube.front, 16)

	// Gestion des cotés
	tmp_left := [3]byte{cube.left[2], cube.left[3], cube.left[4]}
	cube.left[2], cube.left[3], cube.left[4] = cube.down[0], cube.down[1], cube.down[2]
	cube.down[0], cube.down[1], cube.down[2] = cube.right[6], cube.right[7], cube.right[0]
	cube.right[6], cube.right[7], cube.right[0] = cube.up[4], cube.up[5], cube.up[6]
	cube.up[4], cube.up[5], cube.up[6] = tmp_left[0], tmp_left[1], tmp_left[2]
}

func F180(cube Cube) {

	RotateFace(&cube.front, 32)
	cube.left[2], cube.left[3], cube.left[4], cube.right[6], cube.right[7], cube.right[0] = cube.right[6], cube.right[7], cube.right[0], cube.left[2], cube.left[3], cube.left[4]
	cube.up[6], cube.up[5], cube.up[4], cube.down[2], cube.down[1], cube.down[0] = cube.down[2], cube.down[1], cube.down[0], cube.up[6], cube.up[5], cube.up[4]
}

func F270(cube Cube) {

	RotateFace(&cube.front, 48)
	tmp_left := [3]byte{cube.left[2], cube.left[3], cube.left[4]}
	cube.left[2], cube.left[3], cube.left[4] = cube.up[4], cube.up[5], cube.up[6]
	cube.up[4], cube.up[5], cube.up[6] = cube.right[6], cube.right[7], cube.right[0]
	cube.right[6], cube.right[7], cube.right[0] = cube.down[0], cube.down[1], cube.down[2]
	cube.down[0], cube.down[1], cube.down[2] = tmp_left[0], tmp_left[1], tmp_left[2]
}

func B90(cube Cube) {

	RotateFace(&cube.back, 16)
	tmp := [3]byte{cube.right[2], cube.right[3], cube.right[4]}
	cube.right[2], cube.right[3], cube.right[4] = cube.down[4], cube.down[5], cube.down[6]
	cube.down[4], cube.down[5], cube.down[6] = cube.left[6], cube.left[7], cube.left[0]
	cube.left[6], cube.left[7], cube.left[0] = cube.up[0], cube.up[1], cube.up[2]
	cube.up[0], cube.up[1], cube.up[2] = tmp[0], tmp[1], tmp[2]
}

func B180(cube Cube) {

	RotateFace(&cube.back, 32)
	cube.left[6], cube.left[7], cube.left[0], cube.right[2], cube.right[3], cube.right[4] = cube.right[2], cube.right[3], cube.right[4], cube.left[6], cube.left[7], cube.left[0]
	cube.up[0], cube.up[1], cube.up[2], cube.down[4], cube.down[5], cube.down[6] = cube.down[4], cube.down[5], cube.down[6], cube.up[0], cube.up[1], cube.up[2]
}

func B270(cube Cube) {

	RotateFace(&cube.back, 48)
	tmp := [3]byte{cube.right[2], cube.right[3], cube.right[4]}
	cube.right[2], cube.right[3], cube.right[4] = cube.up[0], cube.up[1], cube.up[2]
	cube.up[0], cube.up[1], cube.up[2] = cube.left[6], cube.left[7], cube.left[0]
	cube.left[6], cube.left[7], cube.left[0] = cube.down[4], cube.down[5], cube.down[6]
	cube.down[4], cube.down[5], cube.down[6] = tmp[0], tmp[1], tmp[2]
}

func L90(cube Cube) {

	RotateFace(&cube.left, 16)
	tmp := [3]byte{cube.back[2], cube.back[3], cube.back[4]}
	cube.back[2], cube.back[3], cube.back[4] = cube.down[6], cube.down[7], cube.down[0]
	cube.down[6], cube.down[7], cube.down[0] = cube.front[6], cube.front[7], cube.front[0]
	cube.front[6], cube.front[7], cube.front[0] = cube.up[6], cube.up[7], cube.up[0]
	cube.up[6], cube.up[7], cube.up[0] = tmp[0], tmp[1], tmp[2]
}

func L180(cube Cube) {

	RotateFace(&cube.left, 32)
	cube.back[2], cube.back[3], cube.back[4], cube.front[6], cube.front[7], cube.front[0] = cube.front[6], cube.front[7], cube.front[0], cube.back[2], cube.back[3], cube.back[4]
	cube.up[6], cube.up[7], cube.up[0], cube.down[6], cube.down[7], cube.down[0] = cube.down[6], cube.down[7], cube.down[0], cube.up[6], cube.up[7], cube.up[0]
}

func L270(cube Cube) {

	RotateFace(&cube.left, 48)
	tmp := [3]byte{cube.back[2], cube.back[3], cube.back[4]}
	cube.back[2], cube.back[3], cube.back[4] = cube.up[6], cube.up[7], cube.up[0]
	cube.up[6], cube.up[7], cube.up[0] = cube.front[6], cube.front[7], cube.front[0]
	cube.front[6], cube.front[7], cube.front[0] = cube.down[6], cube.down[7], cube.down[0]
	cube.down[6], cube.down[7], cube.down[0] = tmp[0], tmp[1], tmp[2]
}

func R90(cube Cube) {

	RotateFace(&cube.right, 16)
	tmp := [3]byte{cube.front[2], cube.front[3], cube.front[4]}
	cube.front[2], cube.front[3], cube.front[4] = cube.down[2], cube.down[3], cube.down[4]
	cube.down[2], cube.down[3], cube.down[4] = cube.back[6], cube.back[7], cube.back[0]
	cube.back[6], cube.back[7], cube.back[0] = cube.up[2], cube.up[3], cube.up[4]
	cube.up[2], cube.up[3], cube.up[4] = tmp[0], tmp[1], tmp[2]
}

func R180(cube Cube) {

	RotateFace(&cube.right, 32)
	cube.back[6], cube.back[7], cube.back[0], cube.front[2], cube.front[3], cube.front[4] = cube.front[2], cube.front[3], cube.front[4], cube.back[6], cube.back[7], cube.back[0]
	cube.up[2], cube.up[3], cube.up[4], cube.down[2], cube.down[3], cube.down[4] = cube.down[2], cube.down[3], cube.down[4], cube.up[2], cube.up[3], cube.up[4]
}

func R270(cube Cube) {

	RotateFace(&cube.right, 48)
	tmp := [3]byte{cube.front[2], cube.front[3], cube.front[4]}
	cube.front[2], cube.front[3], cube.front[4] = cube.up[2], cube.up[3], cube.up[4]
	cube.up[2], cube.up[3], cube.up[4] = cube.back[6], cube.back[7], cube.back[0]
	cube.back[6], cube.back[7], cube.back[0] = cube.down[2], cube.down[3], cube.down[4]
	cube.down[2], cube.down[3], cube.down[4] = tmp[0], tmp[1], tmp[2]
}

func U90(cube Cube) {

	RotateFace(&cube.up, 16)
	tmp := [3]byte{cube.left[0], cube.left[1], cube.left[2]}
	cube.left[0], cube.left[1], cube.left[2] = cube.front[0], cube.front[1], cube.front[2]
	cube.front[0], cube.front[1], cube.front[2] = cube.right[0], cube.right[1], cube.right[2]
	cube.right[0], cube.right[1], cube.right[2] = cube.back[0], cube.back[1], cube.back[2]
	cube.back[0], cube.back[1], cube.back[2] = tmp[0], tmp[1], tmp[2]
}

func U180(cube Cube) {

	RotateFace(&cube.up, 32)
	cube.right[0], cube.right[1], cube.right[2], cube.left[0], cube.left[1], cube.left[2] = cube.left[0], cube.left[1], cube.left[2], cube.right[0], cube.right[1], cube.right[2]
	cube.back[0], cube.back[1], cube.back[2], cube.front[0], cube.front[1], cube.front[2] = cube.front[0], cube.front[1], cube.front[2], cube.back[0], cube.back[1], cube.back[2]
}

func U270(cube Cube) {

	RotateFace(&cube.up, 48)
	tmp := [3]byte{cube.left[0], cube.left[1], cube.left[2]}
	cube.left[0], cube.left[1], cube.left[2] = cube.back[0], cube.back[1], cube.back[2]
	cube.back[0], cube.back[1], cube.back[2] = cube.right[0], cube.right[1], cube.right[2]
	cube.right[0], cube.right[1], cube.right[2] = cube.front[0], cube.front[1], cube.front[2]
	cube.front[0], cube.front[1], cube.front[2] = tmp[0], tmp[1], tmp[2]
}

func D90(cube Cube) {

	RotateFace(&cube.down, 16)
	tmp := [3]byte{cube.left[4], cube.left[5], cube.left[6]}
	cube.left[4], cube.left[5], cube.left[6] = cube.back[4], cube.back[5], cube.back[6]
	cube.back[4], cube.back[5], cube.back[6] = cube.right[4], cube.right[5], cube.right[6]
	cube.right[4], cube.right[5], cube.right[6] = cube.front[4], cube.front[5], cube.front[6]
	cube.front[4], cube.front[5], cube.front[6] = tmp[0], tmp[1], tmp[2]
}

func D180(cube Cube) {

	RotateFace(&cube.down, 32)
	cube.right[4], cube.right[5], cube.right[6], cube.left[4], cube.left[5], cube.left[6] = cube.left[4], cube.left[5], cube.left[6], cube.right[4], cube.right[5], cube.right[6]
	cube.back[4], cube.back[5], cube.back[6], cube.front[4], cube.front[5], cube.front[6] = cube.front[4], cube.front[5], cube.front[6], cube.back[4], cube.back[5], cube.back[6]
}

func D270(cube Cube) {

	RotateFace(&cube.down, 48)
	tmp := [3]byte{cube.left[4], cube.left[5], cube.left[6]}
	cube.left[4], cube.left[5], cube.left[6] = cube.front[4], cube.front[5], cube.front[6]
	cube.front[4], cube.front[5], cube.front[6] = cube.right[4], cube.right[5], cube.right[6]
	cube.right[4], cube.right[5], cube.right[6] = cube.back[4], cube.back[5], cube.back[6]
	cube.back[4], cube.back[5], cube.back[6] = tmp[0], tmp[1], tmp[2]
}
