package rubik

import "fmt"

func LayerTwoValid(cube Cube) bool {
	if cube.up[7] == RED && cube.up[3] == RED && cube.right[1] == GREEN && cube.right[5] == GREEN &&
		cube.down[3] == ORANGE && cube.down[7] == ORANGE && cube.left[1] == BLUE && cube.left[5] == BLUE {
		return true
	} else {
		return false
	}
}

func IsGraviting(cube Cube, c1 byte, c2 byte) bool {
	if (cube.up[1] == c1 && cube.back[1] == c2) || (cube.left[7] == c1 && cube.back[3] == c2) ||
		(cube.down[5] == c1 && cube.back[5] == c2) || (cube.right[3] == c1 && cube.back[7] == c2) {
		return true
	} else {
		return false
	}

}

func LayerTwo(incube *Cube, inmoves *[]string, v bool) {
	tmp := *incube
	var seq string

	for !LayerTwoValid(*incube) {
		if tmp.up[1] == RED && tmp.back[1] == BLUE {
			seq = "B' L' B L B U B' U'"
		} else if tmp.left[7] == BLUE && tmp.back[3] == ORANGE {
			seq = "B' D' B D B L B' L'"
		} else if tmp.right[3] == GREEN && tmp.back[7] == RED {
			seq = "B' U' B U B R B' R'"
		} else if tmp.down[5] == ORANGE && tmp.back[5] == GREEN {
			seq = "B' R' B R B D B' D'"
		} else if tmp.up[1] == RED && tmp.back[1] == GREEN {
			seq = "B R B' R' B' U' B U"
		} else if tmp.left[7] == BLUE && tmp.back[3] == RED {
			seq = "B U B' U' B' L' B L"
		} else if tmp.right[3] == GREEN && tmp.back[7] == ORANGE {
			seq = "B D B' D' B' R' B R"
		} else if tmp.down[5] == ORANGE && tmp.back[5] == BLUE {
			seq = "B L B' L' B' D' B D"

		} else if (tmp.up[3] != RED || tmp.right[1] != GREEN) && tmp.up[3] != YELLOW && tmp.right[1] != YELLOW &&
			(tmp.up[1] == YELLOW || tmp.back[1] == YELLOW) && !IsGraviting(tmp, RED, GREEN) {
			seq = "B R B' R' B' U' B U"
		} else if (tmp.up[7] != RED || tmp.left[1] != BLUE) && tmp.up[7] != YELLOW && tmp.left[1] != YELLOW &&
			(tmp.up[1] == YELLOW || tmp.back[1] == YELLOW) && !IsGraviting(tmp, RED, BLUE) {
			seq = "B' L' B L B U B' U'"
		} else if (tmp.up[7] != RED || tmp.left[1] != BLUE) && tmp.up[7] != YELLOW && tmp.left[1] != YELLOW &&
			(tmp.left[7] == YELLOW || tmp.back[3] == YELLOW) && !IsGraviting(tmp, BLUE, RED) {
			seq = "B U B' U' B' L' B L"
		} else if (tmp.left[5] != BLUE || tmp.down[7] != ORANGE) && tmp.left[5] != YELLOW && tmp.down[7] != YELLOW &&
			(tmp.left[7] == YELLOW || tmp.back[3] == YELLOW) && !IsGraviting(tmp, BLUE, ORANGE) {
			seq = "B' D' B D B L B' L'"
		} else if (tmp.left[5] != BLUE || tmp.down[7] != ORANGE) && tmp.left[5] != YELLOW && tmp.down[7] != YELLOW &&
			(tmp.down[5] == YELLOW || tmp.back[5] == YELLOW) && !IsGraviting(tmp, ORANGE, BLUE) {
			seq = "B L B' L' B' D' B D"
		} else if (tmp.down[3] != ORANGE || tmp.right[5] != GREEN) && tmp.down[3] != YELLOW && tmp.right[5] != YELLOW &&
			(tmp.down[5] == YELLOW || tmp.back[5] == YELLOW) && !IsGraviting(tmp, ORANGE, GREEN) {
			seq = "B' R' B R B D B' D'"
		} else if (tmp.down[3] != ORANGE || tmp.right[5] != GREEN) && tmp.down[3] != YELLOW && tmp.right[5] != YELLOW &&
			(tmp.right[3] == YELLOW || tmp.back[7] == YELLOW) && !IsGraviting(tmp, GREEN, ORANGE) {
			seq = "B D B' D' B' R' B R"
		} else if (tmp.up[3] != RED || tmp.right[1] != GREEN) && tmp.up[3] != YELLOW && tmp.right[1] != YELLOW &&
			(tmp.right[3] == YELLOW || tmp.back[7] == YELLOW) && !IsGraviting(tmp, GREEN, RED) {
			seq = "B' U' B U B R B' R'"
		} else {
			seq = "B"
		}

		DoMoves(incube, inmoves, seq)
	}
	if v {
		fmt.Println("Edges of Layer 2 Done !")
		DecodingCube(*incube)
	}
}
