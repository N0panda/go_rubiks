package rubik

import (
	"fmt"
	"strings"
)

func CreatFuncMap() map[string]func(Cube) {
	func_map := map[string]func(Cube){
		"F":  F90,
		"R":  R90,
		"U":  U90,
		"B":  B90,
		"L":  L90,
		"D":  D90,
		"F2": F180,
		"R2": R180,
		"U2": U180,
		"B2": B180,
		"L2": L180,
		"D2": D180,
		"F'": F270,
		"R'": R270,
		"U'": U270,
		"B'": B270,
		"L'": L270,
		"D'": D270}
	return func_map
}

func CopyCube(cube Cube) Cube {

	cpy := InstanceCube()

	copy(cpy.down, cube.down)
	copy(cpy.up, cube.up)
	copy(cpy.left, cube.left)
	copy(cpy.right, cube.right)
	copy(cpy.front, cube.front)
	copy(cpy.back, cube.back)

	return cpy
}

func SolveCube(cube Cube, v bool, d bool) {
	all_moves := []string{}

	cube_test := CopyCube(cube)

	LayerOne(&cube, &all_moves, v)
	LayerTwo(&cube, &all_moves, v)
	LayerTree(&cube, &all_moves, v)

	result := SimplifyMoves(all_moves)

	SequenceShuffle(cube_test, result, d)

	// Final print
	fmt.Println(strings.Join(result, " "))
	if v {
		fmt.Println("Number of moves : ", len(result))
	}
}
