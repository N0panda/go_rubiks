package main

import (
	"flag"
	"fmt"
	"os"
	"rubiks_cube/rubik"
)

type Params struct {
	h_flag        bool
	d_flag        bool
	v_flag        bool
	g_flag        int
	dict_sequence []string
	sequence      string
	generator     bool
}

func argParsing() Params {
	params := Params{}

	flag.BoolVar(&params.h_flag, "h", false, "Display healp")
	flag.BoolVar(&params.d_flag, "d", false, "visualize movements")
	flag.BoolVar(&params.v_flag, "v", false, "verbose")
	flag.IntVar(&params.g_flag, "g", 0, "Positive int represente the number of moves to shuffle the cube instead of use a string, this param is overloarded by a moves string")
	flag.Parse()

	raw_sequence := flag.Args()
	if len(raw_sequence) > 0 {
		//given sequence
		params.sequence = raw_sequence[0]
		params.generator = false
	} else {
		//generator mode
		params.generator = true
	}

	if len(raw_sequence) > 1 {
		fmt.Printf("\nYou have to give only one string to define a shuffle or you can use the generator -g (see -h)\n")
		os.Exit(1)
	}
	if params.h_flag || params.g_flag < 0 {
		flag.Usage()
		os.Exit(1)
	}
	return params
}

func main() {
	params := argParsing()
	if !params.generator {
		_, err := rubik.ParseSequence(&params.sequence, &params.dict_sequence)
		if err != nil {
			fmt.Println("\nToken error :\n\n", err)
			os.Exit(1)
		}
	}
	// Create crube and shuffle it (generator or given sequence)
	cube := rubik.CreateCube(params.dict_sequence, params.generator, params.g_flag)

	// solving the cube
	rubik.SolveCube(cube, params.v_flag, params.d_flag)
}
