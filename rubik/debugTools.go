package rubik

import (
	"fmt"
)

var S_WHITE string = "\033[38;5;231m"
var S_YELLOW string = "\033[38;5;226m"
var S_RED string = "\033[38;5;160m"
var S_ORANGE string = "\033[38;5;208m"
var S_BLUE string = "\033[38;5;33m"
var S_GREEN string = "\033[38;5;40m"

var S_color = map[int8]string{
	0x01: S_WHITE,
	0x02: S_YELLOW,
	0x03: S_RED,
	0x04: S_ORANGE,
	0x05: S_BLUE,
	0x06: S_GREEN,
}

func WriteFace(face *[3][3]string, cb []byte, letter string, letter_color string) {

	face[0][0] = S_color[int8(cb[0])] + Color[int8(cb[0])] + "\033[0m"
	face[0][1] = S_color[int8(cb[1])] + Color[int8(cb[1])] + "\033[0m"
	face[0][2] = S_color[int8(cb[2])] + Color[int8(cb[2])] + "\033[0m"
	face[1][2] = S_color[int8(cb[3])] + Color[int8(cb[3])] + "\033[0m"
	face[2][2] = S_color[int8(cb[4])] + Color[int8(cb[4])] + "\033[0m"
	face[2][1] = S_color[int8(cb[5])] + Color[int8(cb[5])] + "\033[0m"
	face[2][0] = S_color[int8(cb[6])] + Color[int8(cb[6])] + "\033[0m"
	face[1][0] = S_color[int8(cb[7])] + Color[int8(cb[7])] + "\033[0m"
	face[1][1] = letter_color + letter + "\033[0m"
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
	list_scolor := []string{S_RED, S_ORANGE, S_WHITE, S_YELLOW, S_BLUE, S_GREEN}

	for i := 0; i < len(list_hface); i++ {

		WriteFace(list_hface[i], list_face[i], list_letter[i], list_scolor[i])
	}
	DebbugDisplay(hc)
}
