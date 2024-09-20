package vegenere

type Vegener struct {
	table [][]rune
}

func Shift(s int, slice []rune) []rune {
	// new slice for swap
	nSlice := make([]rune, len(slice))
	for i := range nSlice {
		if i == len(slice)-1 {
			nSlice[i] = slice[0]
			continue
		}
		nSlice[i] = slice[i+1]
	}
	if s > 1 {
		return Shift(s-1, nSlice)
	}
	return nSlice
}
