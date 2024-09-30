package des

import (
	"fmt"
	"slices"
)

type DES struct {
	key  int
	Bkey []bool
	K1   []bool
	K2   []bool
}

func (d *DES) Encrypt(text []byte) {
	for _, symbol := range text {
		sl := toBoolSlice(int(symbol))
		// divide on two parts
		fPart, sPart := divide(sl)
		fmt.Printf("First part: %s; Second part: %s\n", PrintD(fPart), PrintD(sPart))
		// extentend right part

	}
}

func (d *DES) f(xr []bool) {
	//extend xr to 8 bits
	xr = extension(xr)
	// xor
	xr = xor(xr, d.K1)
	// divide on 2 parts
	// fPart, sPart := divide(xr)
	// s1 & s2

}
func s1(b []bool) []bool {
	s1 := [][]int{
		{1, 0, 3, 2},
		{3, 2, 1, 0},
		{0, 2, 1, 3},
		{3, 1, 3, 2},
	}
	a14 := []bool{b[0], b[3]}
	a23 := []bool{b[1], b[2]}
	return toBoolSlice(s1[bToInt(a14)][bToInt(a23)])
}

func xor(f []bool, s []bool) []bool {
	result := make([]bool, 0)
	for i := range result {
		if f[i] != s[i] {
			result = append(result, true)
			continue
		}
		result = append(result, false)
	}
	return result
}

func extension(b []bool) []bool {
	return []bool{b[3], b[0], b[1], b[2], b[1], b[2], b[3], b[1]}
}

func divide(b []bool) ([]bool, []bool) {
	fmt.Printf("B: %d\n", bToInt(b))
	fmt.Println(PrintD(b))
	return b[:len(b)/2+1], b[len(b)/2:]
}

func NewDES(key int) *DES {
	Bkey := toBoolSlice(key)
	return &DES{
		key:  key,
		Bkey: Bkey,
	}
}

func toBoolSlice(key int) []bool {
	bits := make([]bool, 0)
	for key != 0 {
		bits = append(bits, key&1 == 1)
		key >>= 1
	}
	slices.Reverse(bits)
	return bits
}
func (d *DES) P10() {
	p10Pattern := []int{3, 5, 2, 7, 4, 10, 1, 9, 8, 6}
	p10Result := make([]bool, 0)
	for _, i := range p10Pattern {
		i--
		// log.Println("Putting ", d.Bkey[i], " to ", i, " place")
		p10Result = append(p10Result, d.Bkey[i])
	}
	// log.Println("p10Result: ", p10Result)
	d.Bkey = p10Result

}
func PrintD(dslice []bool) (result string) {
	result = ""
	for _, v := range dslice {
		if v {
			result += "1"
		} else {
			result += "0"
		}
	}
	return
}
func (d *DES) GenerateKeys() {
	// fmt.Println("before P10: ", PrintD(d.Bkey))
	d.P10()
	// fmt.Println("after P10: ", PrintD(d.Bkey))
	// fmt.Println("Left part: ", PrintD(d.Bkey[:5]))
	// fmt.Println("Right part: ", PrintD(d.Bkey[5:]))
	fPart := d.Bkey[:5]
	sPart := d.Bkey[5:]
	// fmt.Println("First part before <<1: ", PrintD(fPart))
	fPart = d.lShift(1, fPart)
	// fmt.Println("First part after <<1: ", PrintD(fPart))
	// fmt.Println("Second part before <<1: ", PrintD(sPart))
	sPart = d.lShift(1, sPart)
	// fmt.Println("Second part after <<1: ", PrintD(sPart))
	// fmt.Println("before P8: ", PrintD(d.Bkey))
	nBkey := append(fPart, sPart...)
	b8key := d.P8(nBkey)
	// fmt.Println("after P8: ", PrintD(b8key))
	// fmt.Println("new K1: ", PrintD(b8key))
	d.K1 = b8key
	// fmt.Println("First part before <<2: ", PrintD(fPart))
	fPart = d.lShift(2, fPart)
	// fmt.Println("First part after <<2: ", PrintD(fPart))
	// fmt.Println("Second part before <<2: ", PrintD(sPart))
	sPart = d.lShift(2, sPart)
	// fmt.Println("Second part after <<2: ", PrintD(sPart))
	d.K2 = d.P8(append(fPart, sPart...))
	// fmt.Println("new K2: ", PrintD(d.K2))
}
func (d *DES) lShift(shift int, s []bool) []bool {
	r := make([]bool, len(s))
	for i := range s {
		if i == len(r)-1 {
			r[i] = s[0]
			continue
		}
		r[i] = s[i+1]
	}
	shift--
	if shift >= 1 {
		return d.lShift(shift, r)
	}
	return r
}
func (d *DES) P8(bk []bool) []bool {
	p8Pattern := []int{6, 3, 7, 4, 8, 5, 10, 9}
	nBkey := make([]bool, 0)
	for _, v := range p8Pattern {
		nBkey = append(nBkey, bk[v-1])
	}
	return nBkey
}
func bToInt(b []bool) int {
	result := 0
	for i := len(b) - 1; i >= 0; i-- {
		result <<= 1
		if b[i] {
			result |= 1
		}
	}
	return result
}
