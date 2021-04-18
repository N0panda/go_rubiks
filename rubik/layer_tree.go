package rubik

import (
	"fmt"
	"os"
)

func YellowCrossValid(cube Cube) bool {
	if cube.back[1] == YELLOW && cube.back[3] == YELLOW && cube.back[5] == YELLOW && cube.back[7] == YELLOW {
		return true
	} else {
		return false
	}
}

func DoYellowCross(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string

	for !YellowCrossValid(*incube) {
		if tmp.back[1] != YELLOW && tmp.back[3] != YELLOW && tmp.back[5] != YELLOW && tmp.back[7] != YELLOW {
			seq = "U R B R' B' U'"

		} else if tmp.back[1] == YELLOW && tmp.back[3] == YELLOW {
			seq = "R D B D' B' R'"
		} else if tmp.back[3] == YELLOW && tmp.back[5] == YELLOW {
			seq = "U R B R' B' U'"
		} else if tmp.back[5] == YELLOW && tmp.back[7] == YELLOW {
			seq = "L U B U' B' L'"
		} else if tmp.back[7] == YELLOW && tmp.back[1] == YELLOW {
			seq = "D L B L' B' D'"

		} else if tmp.back[1] == YELLOW && tmp.back[5] == YELLOW {
			seq = "R D B D' B' R'"
		} else if tmp.back[3] == YELLOW && tmp.back[7] == YELLOW {
			seq = "U R B R' B' U'"
		}
		DoMoves(incube, inmoves, seq)
	}
}

func YellowEdgesValid(cube Cube) bool {
	if cube.up[1] == RED && cube.left[7] == BLUE && cube.right[3] == GREEN && cube.down[5] == ORANGE {
		return true
	} else {
		return false
	}
}

func CountYellowEdges(cube Cube) int {
	count := 0

	if cube.up[1] == RED {
		count++
	}
	if cube.left[7] == BLUE {
		count++
	}
	if cube.right[3] == GREEN {
		count++
	}
	if cube.down[5] == ORANGE {
		count++
	}
	return count
}

func PlaceEdges(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string

	for !YellowEdgesValid(*incube) {
		if CountYellowEdges(*incube) < 2 {
			seq = "B"
		} else if tmp.up[1] == RED && tmp.down[5] == ORANGE {
			seq = "R B R' B R B2 R'"
		} else if tmp.left[7] == BLUE && tmp.right[3] == GREEN {
			seq = "D B D' B D B2 D'"
		} else if tmp.right[3] == GREEN && tmp.down[5] == ORANGE {
			seq = "R B R' B R B2 R'"
		} else if tmp.left[7] == BLUE && tmp.down[5] == ORANGE {
			seq = "D B D' B D B2 D'"
		} else if tmp.up[1] == RED && tmp.right[3] == GREEN {
			seq = "U B U' B U B2 U'"
		} else if tmp.up[1] == RED && tmp.left[7] == BLUE {
			seq = "L B L' B L B2 L'"
		} else {
			fmt.Println("PLACE YELLOW EDGE :  ERROR CASE")
			os.Exit(1)
		}
		DoMoves(incube, inmoves, seq)
	}
}

func YellwoCornerValid(cube Cube) int {
	placed := 0

	if (cube.up[2] == RED || cube.up[2] == GREEN || cube.up[2] == YELLOW) &&
		(cube.right[2] == RED || cube.right[2] == GREEN || cube.right[2] == YELLOW) &&
		(cube.back[0] == RED || cube.back[0] == GREEN || cube.back[0] == YELLOW) {
		placed++
	}
	if (cube.up[0] == RED || cube.up[0] == BLUE || cube.up[0] == YELLOW) &&
		(cube.left[0] == RED || cube.left[0] == BLUE || cube.left[0] == YELLOW) &&
		(cube.back[2] == RED || cube.back[2] == BLUE || cube.back[2] == YELLOW) {
		placed++
	}
	if (cube.down[6] == ORANGE || cube.down[6] == BLUE || cube.down[6] == YELLOW) &&
		(cube.left[6] == ORANGE || cube.left[6] == BLUE || cube.left[6] == YELLOW) &&
		(cube.back[4] == ORANGE || cube.back[4] == BLUE || cube.back[4] == YELLOW) {
		placed++
	}
	if (cube.down[4] == ORANGE || cube.down[4] == GREEN || cube.down[4] == YELLOW) &&
		(cube.right[4] == ORANGE || cube.right[4] == GREEN || cube.right[4] == YELLOW) &&
		(cube.back[6] == ORANGE || cube.back[6] == GREEN || cube.back[6] == YELLOW) {
		placed++
	}
	return placed
}

func PlaceYellowCorner(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string

	for YellwoCornerValid(*incube) != 4 {
		placed := YellwoCornerValid(*incube)
		if placed == 0 {
			seq = "B R B' L' B R' B' L"
		} else if (tmp.up[2] == RED || tmp.up[2] == GREEN || tmp.up[2] == YELLOW) &&
			(tmp.right[2] == RED || tmp.right[2] == GREEN || tmp.right[2] == YELLOW) &&
			(tmp.back[0] == RED || tmp.back[0] == GREEN || tmp.back[0] == YELLOW) {
			seq = "B R B' L' B R' B' L"
		} else if (tmp.down[6] == ORANGE || tmp.down[6] == BLUE || tmp.down[6] == YELLOW) &&
			(tmp.left[6] == ORANGE || tmp.left[6] == BLUE || tmp.left[6] == YELLOW) &&
			(tmp.back[4] == ORANGE || tmp.back[4] == BLUE || tmp.back[4] == YELLOW) {
			seq = "B L B' R' B L' B' R"
		} else if (tmp.up[0] == RED || tmp.up[0] == BLUE || tmp.up[0] == YELLOW) &&
			(tmp.left[0] == RED || tmp.left[0] == BLUE || tmp.left[0] == YELLOW) &&
			(tmp.back[2] == RED || tmp.back[2] == BLUE || tmp.back[2] == YELLOW) {
			seq = "B U B' D' B U' B' D"
		} else if (tmp.down[4] == ORANGE || tmp.down[4] == GREEN || tmp.down[4] == YELLOW) &&
			(tmp.right[4] == ORANGE || tmp.right[4] == GREEN || tmp.right[4] == YELLOW) &&
			(tmp.back[6] == ORANGE || tmp.back[6] == GREEN || tmp.back[6] == YELLOW) {
			seq = "B D B' U' B D' B' U"
		}
		DoMoves(incube, inmoves, seq)
	}
}

func FullValid(cube Cube) bool {
	if cube.up[0] == RED && cube.up[2] == RED && cube.left[0] == BLUE && cube.left[6] == BLUE &&
		cube.right[2] == GREEN && cube.right[4] == GREEN && cube.down[4] == ORANGE && cube.down[6] == ORANGE &&
		cube.back[0] == YELLOW && cube.back[2] == YELLOW && cube.back[4] == YELLOW && cube.back[6] == YELLOW {
		return true
	} else {
		return false
	}
}

func OrientYellowCorner(incube *Cube, inmoves *[]string) {
	tmp := *incube
	var seq string

	for !FullValid(*incube) {
		if tmp.back[0] != YELLOW {
			seq = "R' F' R F R' F' R F"
		} else {
			seq = "B"
		}

		DoMoves(incube, inmoves, seq)
	}
}

func LayerTree(incube *Cube, inmoves *[]string, v bool) {
	DoYellowCross(incube, inmoves)
	if v {
		fmt.Println("Yellow cross Done !")
		DecodingCube(*incube)
	}

	PlaceEdges(incube, inmoves)
	if v {
		fmt.Println("Yellow edges placed !")
		DecodingCube(*incube)
	}

	PlaceYellowCorner(incube, inmoves)
	if v {
		fmt.Println("Yellow corners placed !")
		DecodingCube(*incube)
	}

	OrientYellowCorner(incube, inmoves)
	if v {
		fmt.Println("Yellow corners oriented !")
		DecodingCube(*incube)
		fmt.Println("Done !")
	}
}
