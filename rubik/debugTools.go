package rubik

import (
	"fmt"
)

func WriteFace(face *[3][3]string, cb []byte, letter string) {

	face[0][0] = Color[int8(cb[0])]
	face[0][1] = Color[int8(cb[1])]
	face[0][2] = Color[int8(cb[2])]
	face[1][2] = Color[int8(cb[3])]
	face[2][2] = Color[int8(cb[4])]
	face[2][1] = Color[int8(cb[5])]
	face[2][0] = Color[int8(cb[6])]
	face[1][0] = Color[int8(cb[7])]
	face[1][1] = letter
}

func DebbugDisplay(hc Hcube) {
	fmt.Println()
	fmt.Println("       ", hc.up[0])
	fmt.Println("       ", hc.up[1])
	fmt.Println("       ", hc.up[2])
	fmt.Println()
	fmt.Println(hc.left[0], hc.front[0], hc.right[0], hc.back[0])
	fmt.Println(hc.left[1], hc.front[1], hc.right[1], hc.back[1])
	fmt.Println(hc.left[2], hc.front[2], hc.right[2], hc.back[2])
	fmt.Println()
	fmt.Println("       ", hc.down[0])
	fmt.Println("       ", hc.down[1])
	fmt.Println("       ", hc.down[2])
	fmt.Println()
}

func DecodingCube(cube Cube) {
	hc := Hcube{}
	list_hface := []*[3][3]string{&hc.up, &hc.down, &hc.front, &hc.back, &hc.left, &hc.right}
	list_face := [][]byte{cube.up, cube.down, cube.front, cube.back, cube.left, cube.right}
	list_letter := []string{"R", "O", "W", "Y", "B", "G"}

	for i := 0; i < len(list_hface); i++ {

		WriteFace(list_hface[i], list_face[i], list_letter[i])
	}
	DebbugDisplay(hc)
}
