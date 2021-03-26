package rubik

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func OneMove(cube Cube, move string) {

	switch move {
	case "F":
		F90(cube)
	case "B":
		B90(cube)
	case "U":
		U90(cube)
	case "D":
		D90(cube)
	case "R":
		R90(cube)
	case "L":
		L90(cube)
	case "F'":
		F270(cube)
	case "B'":
		B270(cube)
	case "U'":
		U270(cube)
	case "D'":
		D270(cube)
	case "R'":
		R270(cube)
	case "L'":
		L270(cube)
	case "F2":
		F180(cube)
	case "B2":
		B180(cube)
	case "U2":
		U180(cube)
	case "D2":
		D180(cube)
	case "R2":
		R180(cube)
	case "L2":
		L180(cube)
	default:
		fmt.Println("An error occured during the sequence of shuffling movement with the token : ", move)
		os.Exit(1)
	}
}

func GenerateShuffle(cube Cube, nb_iter int) {
	// take nb_iter times a random move in the list and apply it to the cube
	dict := strings.Fields("F R U B L D F2 R2 U2 B2 L2 D2 F' R' U' B' L' D'")
	max := len(dict)

	for i := 0; i < nb_iter; i++ {
		random := dict[rand.Intn(max)]
		OneMove(cube, random)
	}
}

func SequenceShuffle(cube Cube, sequence []string) {
	for _, move := range sequence {
		OneMove(cube, move)
	}
}
