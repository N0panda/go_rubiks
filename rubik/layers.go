package rubik

import "strings"

var WHITE byte = 1
var YELLOW byte = 2
var RED byte = 3
var ORANGE byte = 4
var BLUE byte = 5
var GREEN byte = 6

func DoMoves(cube *Cube, moves *[]string, mv string) {

	mv_tab := strings.Fields(mv)

	for _, elem := range mv_tab {
		OneMove(*cube, elem)
		*moves = append(*moves, elem)
	}
}
