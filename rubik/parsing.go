package rubik

import (
	"fmt"
	"strings"
)

type ErrToken struct {
	info string
	err  string
}

func (err ErrToken) Error() string {
	return fmt.Sprintf("Your Token [%s] is a bad token\nPlease only use tokens from this list : \n%s\n", err.err, err.info)
}

var dict_map map[string]bool

func SimplifySequence(dict_sequence *[]string) {
	for i := 0; i < len(*dict_sequence); i++ {
		switch (*dict_sequence)[i] {
		case "U'2", "U2'":
			(*dict_sequence)[i] = "U2"
		case "D'2", "D2'":
			(*dict_sequence)[i] = "D2"
		case "F'2", "F2'":
			(*dict_sequence)[i] = "F2"
		case "B'2", "B2'":
			(*dict_sequence)[i] = "B2"
		case "R'2", "R2'":
			(*dict_sequence)[i] = "R2"
		case "L'2", "L2'":
			(*dict_sequence)[i] = "L2"
		}
	}
}

func ParseSequence(sequence *string, dict_sequence *[]string) (int, error) {

	var dict string = "F R U B L D F2 R2 U2 B2 L2 D2 F' R' U' B' L' D' F2' R2' U2' B2' L2' D2' F'2 R'2 U'2 B'2 L'2 D'2"
	dict_map = make(map[string]bool)

	*dict_sequence = strings.Fields(*sequence)

	for _, j := range strings.Fields(dict) {
		dict_map[j] = true
	}
	for i, j := range *dict_sequence {
		(*dict_sequence)[i] = strings.ToUpper(j)
		j = (*dict_sequence)[i]
		if !dict_map[j] {
			return 0, ErrToken{dict, j}
		}
	}
	SimplifySequence(dict_sequence)
	return 0, nil
}
