package rubik

import (
	"fmt"
	"os"
)

func WhiteCrossValid(cube Cube) bool {
	if cube.front[1] == WHITE && cube.front[3] == WHITE && cube.front[5] == WHITE && cube.front[7] == WHITE {
		return true
	} else {
		return false
	}
}

func DoWhiteCross(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string
	for !WhiteCrossValid(*incube) {
		if tmp.front[3] != WHITE && tmp.up[3] == WHITE {
			seq = "R'"
		} else if tmp.front[3] != WHITE && tmp.back[7] == WHITE {
			seq = "R2"
		} else if tmp.front[3] != WHITE && tmp.down[3] == WHITE {
			seq = "R"
		} else if tmp.front[7] != WHITE && tmp.up[7] == WHITE {
			seq = "L"
		} else if tmp.front[7] != WHITE && tmp.back[3] == WHITE {
			seq = "L2"
		} else if tmp.front[7] != WHITE && tmp.down[7] == WHITE {
			seq = "L'"
		} else if tmp.front[1] != WHITE && tmp.left[1] == WHITE {
			seq = "U'"
		} else if tmp.front[1] != WHITE && tmp.back[1] == WHITE {
			seq = "U2"
		} else if tmp.front[1] != WHITE && tmp.right[1] == WHITE {
			seq = "U"
		} else if tmp.front[5] != WHITE && tmp.left[5] == WHITE {
			seq = "D"
		} else if tmp.front[5] != WHITE && tmp.back[5] == WHITE {
			seq = "D2"
		} else if tmp.front[5] != WHITE && tmp.right[5] == WHITE {
			seq = "D'"
		} else if (tmp.front[3] == WHITE && tmp.up[3] == WHITE) || (tmp.front[3] == WHITE && tmp.back[7] == WHITE) || (tmp.front[3] == WHITE && tmp.down[3] == WHITE) ||
			(tmp.front[7] == WHITE && tmp.up[7] == WHITE) || (tmp.front[7] == WHITE && tmp.back[3] == WHITE) || (tmp.front[7] == WHITE && tmp.down[7] == WHITE) ||
			(tmp.front[1] == WHITE && tmp.left[1] == WHITE) || (tmp.front[1] == WHITE && tmp.back[1] == WHITE) || (tmp.front[1] == WHITE && tmp.right[1] == WHITE) ||
			(tmp.front[5] == WHITE && tmp.left[5] == WHITE) || (tmp.front[5] == WHITE && tmp.back[5] == WHITE) || (tmp.front[5] == WHITE && tmp.right[5] == WHITE) {
			seq = "F"
		} else if (tmp.up[5] == WHITE) || ((tmp.up[1] == WHITE) && (tmp.front[1] != WHITE)) {
			seq = "U"
		} else if (tmp.left[3]) == WHITE || ((tmp.left[7] == WHITE) && (tmp.front[7] != WHITE)) {
			seq = "L"
		} else if (tmp.down[1] == WHITE) || ((tmp.down[5] == WHITE) && (tmp.front[5] != WHITE)) {
			seq = "D"
		} else if (tmp.right[7] == WHITE) || ((tmp.right[3] == WHITE) && (tmp.front[3] != WHITE)) {
			seq = "R"
		} else {
			seq = "B"
		}
		DoMoves(incube, inmoves, seq)
	}
}

func OrientWhiteCrossValid(cube Cube) bool {
	if cube.up[5] == RED && cube.left[3] == BLUE && cube.right[7] == GREEN && cube.down[1] == ORANGE {
		return true
	} else {
		return false
	}
}

func OrientWhiteCross(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string

	for !OrientWhiteCrossValid(*incube) {
		if tmp.up[5] == BLUE && tmp.left[3] == ORANGE && tmp.down[1] == GREEN && tmp.right[7] == RED {
			seq = "F'"
		} else if tmp.up[5] == ORANGE && tmp.left[3] == GREEN && tmp.down[1] == RED && tmp.right[7] == BLUE {
			seq = "F2"
		} else if tmp.up[5] == GREEN && tmp.left[3] == RED && tmp.down[1] == BLUE && tmp.right[7] == ORANGE {
			seq = "F"
		} else if tmp.up[5] == ORANGE && tmp.down[1] == RED {
			seq = "U2 D2 B2 U2 D2"
		} else if tmp.left[3] == GREEN && tmp.right[7] == BLUE {
			seq = "L2 R2 B2 L2 R2"
		} else if tmp.up[5] == GREEN || tmp.right[7] == RED {
			seq = "U2 R2 B U2 B2 R2"
		} else if tmp.up[5] == BLUE || tmp.left[3] == RED {
			seq = "U2 L2 B' U2 B2 L2"
		} else if tmp.down[1] == GREEN || tmp.right[7] == ORANGE {
			seq = "D2 R2 B R2 B2 D2"
		} else if tmp.down[1] == BLUE || tmp.left[3] == ORANGE {
			seq = "D2 L2 B' L2 B2 D2"
		} else {
			fmt.Println("******************* PROBLEME OrientWhiteCrossValid *****************")
			os.Exit(1)
		}
		DoMoves(incube, inmoves, seq)
	}
}

func WhiteCornersValid(cube Cube) bool {
	if cube.front[0] == WHITE && cube.front[2] == WHITE && cube.front[4] == WHITE && cube.front[6] == WHITE &&
		cube.up[4] == RED && cube.up[6] == RED && cube.down[0] == ORANGE && cube.down[2] == ORANGE &&
		cube.left[2] == BLUE && cube.left[4] == BLUE && cube.right[0] == GREEN && cube.right[6] == GREEN {
		return true
	} else {
		return false
	}
}

func PlaceWhiteCorners(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string

	for !WhiteCornersValid(*incube) {
		if tmp.up[0] == WHITE && (tmp.left[0] == RED || tmp.left[0] == BLUE) && (tmp.back[2] == RED || tmp.back[2] == BLUE) {
			seq = "L U' L' U"
		} else if tmp.left[0] == WHITE && (tmp.up[0] == RED || tmp.up[0] == BLUE) && (tmp.back[2] == RED || tmp.back[2] == BLUE) {
			seq = "U' L U L'"
		} else if tmp.up[2] == WHITE && (tmp.right[2] == RED || tmp.right[2] == GREEN) && (tmp.back[0] == RED || tmp.back[0] == GREEN) {
			seq = "R' U R U'"
		} else if tmp.right[2] == WHITE && (tmp.up[2] == RED || tmp.up[2] == GREEN) && (tmp.back[0] == RED || tmp.back[0] == GREEN) {
			seq = "U R' U' R"
		} else if tmp.down[6] == WHITE && (tmp.left[6] == ORANGE || tmp.left[6] == BLUE) && (tmp.back[4] == ORANGE || tmp.back[4] == BLUE) {
			seq = "L' D L D'"
		} else if tmp.left[6] == WHITE && (tmp.down[6] == ORANGE || tmp.down[6] == BLUE) && (tmp.back[4] == ORANGE || tmp.back[4] == BLUE) {
			seq = "D L' D' L"
		} else if tmp.right[4] == WHITE && (tmp.down[4] == ORANGE || tmp.down[4] == GREEN) && (tmp.back[6] == ORANGE || tmp.back[6] == GREEN) {
			seq = "D' R D R'"
		} else if tmp.down[4] == WHITE && (tmp.right[4] == ORANGE || tmp.right[4] == GREEN) && (tmp.back[6] == ORANGE || tmp.back[6] == GREEN) {
			seq = "R D' R' D"
		} else if tmp.up[0] == WHITE || tmp.up[2] == WHITE || tmp.right[2] == WHITE || tmp.right[4] == WHITE ||
			tmp.down[4] == WHITE || tmp.down[6] == WHITE || tmp.left[6] == WHITE || tmp.left[0] == WHITE {
			seq = "B"
		} else if tmp.front[0] == WHITE && (tmp.up[6] != RED || tmp.left[2] != BLUE) {
			seq = "L' B' L"
		} else if tmp.front[2] == WHITE && (tmp.up[4] != RED || tmp.right[0] != GREEN) {
			seq = "R B R'"
		} else if tmp.front[4] == WHITE && (tmp.down[2] != ORANGE || tmp.right[6] != GREEN) {
			seq = "R' B' R"
		} else if tmp.front[6] == WHITE && (tmp.down[0] != ORANGE || tmp.left[4] != BLUE) {
			seq = "L B L'"
		} else if tmp.back[0] == WHITE && tmp.front[2] != WHITE {
			seq = "R B' R'"
		} else if tmp.back[2] == WHITE && tmp.front[0] != WHITE {
			seq = "L' B L"
		} else if tmp.back[4] == WHITE && tmp.front[6] != WHITE {
			seq = "L B' L'"
		} else if tmp.back[6] == WHITE && tmp.front[4] != WHITE {
			seq = "R' B R"
		} else if tmp.up[6] == WHITE {
			seq = "U B U'"
		} else if tmp.up[4] == WHITE {
			seq = "U' B' U"
		} else if tmp.right[0] == WHITE {
			seq = "R B R'"
		} else if tmp.right[6] == WHITE {
			seq = "R' B' R"
		} else if tmp.down[2] == WHITE {
			seq = "D B D'"
		} else if tmp.down[0] == WHITE {
			seq = "D' B' D"
		} else if tmp.left[4] == WHITE {
			seq = "L B L'"
		} else if tmp.left[2] == WHITE {
			seq = "L' B' L"
		} else {
			seq = "B"
		}
		DoMoves(incube, inmoves, seq)
	}
}

func LayerOne(incube *Cube, inmoves *[]string, v bool) {

	DoWhiteCross(incube, inmoves)
	if v {
		fmt.Println("\nWhite cross Done !")
		DecodingCube(*incube)
	}
	OrientWhiteCross(incube, inmoves)
	if v {
		fmt.Println("White cross Oriented !")
		DecodingCube(*incube)
	}

	PlaceWhiteCorners(incube, inmoves)
	if v {
		fmt.Println("White corners placed !")
		DecodingCube(*incube)
	}
}
