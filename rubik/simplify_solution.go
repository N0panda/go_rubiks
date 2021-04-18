package rubik

func Equalcase(mv *[]string, i int) {
	tmp := (*mv)[i]

	if len(tmp) == 1 {
		*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
		end := append([]string{(tmp + "2")}, (*mv)[i-1:]...)
		*mv = append((*mv)[:i-1], end...)
	} else if tmp[1] == '2' {
		*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
	} else if tmp[1] == '\'' {
		*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
		end := append([]string{(tmp[:1] + "2")}, (*mv)[i-1:]...)
		*mv = append((*mv)[:i-1], end...)
	}
}

func SameNatureCase(mv *[]string, i int) {
	tmp1 := (*mv)[i-1]
	tmp2 := (*mv)[i]

	if len(tmp1) == 1 {
		if tmp2[1] == '\'' {
			*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
		} else if tmp2[1] == '2' {
			*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
			end := append([]string{(tmp1 + "'")}, (*mv)[i-1:]...)
			*mv = append((*mv)[:i-1], end...)
		}
	} else if len(tmp2) == 1 {
		if tmp1[1] == '\'' {
			*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
		} else if tmp1[1] == '2' {
			*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
			end := append([]string{(tmp2 + "'")}, (*mv)[i-1:]...)
			*mv = append((*mv)[:i-1], end...)
		}
	} else if tmp1[1] == '\'' {
		*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
		end := append([]string{string(tmp1[0])}, (*mv)[i-1:]...)
		*mv = append((*mv)[:i-1], end...)
	} else if tmp2[1] == '\'' {
		*mv = append((*mv)[:(i-1)], (*mv)[i+1:]...)
		end := append([]string{string(tmp2[0])}, (*mv)[i-1:]...)
		*mv = append((*mv)[:i-1], end...)
	}
}

func DeleteMoves(moves *[]string) bool {

	for i := 1; i < len(*moves); i++ {
		if (*moves)[i] == (*moves)[i-1] {
			Equalcase(moves, i)
			return true
		} else if (*moves)[i][0] == (*moves)[i-1][0] {
			SameNatureCase(moves, i)
		}
	}
	return false
}

func SimplifyMoves(moves []string) []string {

	result := make([]string, len(moves))
	copy(result, moves)

	for DeleteMoves(&result) {

	}
	return result
}
