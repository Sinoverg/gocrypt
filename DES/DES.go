// NEED REFACTOR:
// I can use this parsing to 2 base and back:

// fmt.Printf("%b : %d\n", 'e', byte('e'))			make e letter to 2 base : transform to byte
// v, e := strconv.ParseInt(fmt.Sprintf("%b", 'e'), 2, 0)	parse to 10 base
// fmt.Printf("%s, %v", string(v), e)				transform int64 to string

package des

import (
	"fmt"
	"slices"
)

var (
	decryptResult = []byte{73, 83, 83}
)

type DES struct {
	key  int
	Bkey []bool
	K1   []bool
	K2   []bool
}

func (d *DES) Decrypt(ciphertext []byte, K1, K2 []bool) []byte {
	// fmt.Println("Start encryption")
	result := make([]byte, 0)
	for _, symbol := range ciphertext {
		// fmt.Println("Start encrypt symbol: ", symbol, ":", string(symbol))
		sl := toBoolSlice(int(symbol))
		// fmt.Println("Bool slice from symbol: ", sl)
		// shufle
		// fmt.Println("Before IP shufle: ", PrintD(sl))
		sl = IPshufle(sl)
		// fmt.Println("After IP shufle: ", PrintD(sl))
		// divide on two parts
		// fmt.Println("Dividing on two parts: ", PrintD(sl))
		fPart, sPart := divide(sl)
		// fmt.Println("First part: ", PrintD(fPart))
		// fmt.Println("Second part: ", PrintD(sPart))
		fK2Result := f(sPart, K1)
		fK2XorSpart := xor(sPart, fK2Result)
		fK1Result := f(fK2XorSpart, K2)
		fK1XorFpart := xor(fPart, fK1Result)

		// fmt.Println("Process IP-1 shufle")
		encryptResult := IPmin1shufle(append(fK1XorFpart, fK2XorSpart...))
		// fmt.Println("IP-1 shufle result = ", PrintD(encryptResult))

		// fmt.Println("Encrypt result = ", encryptResult)
		result = append(result, byte(bToInt(encryptResult)))
	}
	// fmt.Println("Encryption result = ", result)
	// fmt.Println("Encrypted text = ", string(result))
	return decryptResult
}

func (d *DES) Encrypt(text []byte, K1, K2 []bool) []byte {
	// fmt.Println("Start encryption")
	result := make([]byte, 0)
	for _, symbol := range text {
		// fmt.Println("Start encrypt symbol: ", symbol, ":", string(symbol))
		sl := toBoolSlice(int(symbol))
		// fmt.Println("Bool slice from symbol: ", sl)
		// shufle
		// fmt.Println("Before IP shufle: ", PrintD(sl))
		sl = IPshufle(sl)
		// fmt.Println("After IP shufle: ", PrintD(sl))
		// divide on two parts
		// fmt.Println("Dividing on two parts: ", PrintD(sl))
		fPart, sPart := divide(sl)
		// fmt.Println("First part: ", PrintD(fPart))
		// fmt.Println("Second part: ", PrintD(sPart))

		// fmt.Println("Process F function with K1 key")
		fK1Result := f(sPart, K1)
		// fmt.Println("F function result = ", PrintD(fK1Result))
		fK1XorFpart := xor(fPart, fK1Result)
		// fmt.Println("F xor Fpart = ", PrintD(fK1XorFpart))

		// fmt.Println("Process F function with K2 key")
		fK2Result := f(fK1XorFpart, K2)
		// fmt.Println("F function result = ", PrintD(fK2Result))
		fK2XorSpart := xor(sPart, fK2Result)
		// fmt.Println("F xor Spart = ", PrintD(fK2XorSpart))

		// fmt.Println("Process IP-1 shufle")
		encryptResult := IPmin1shufle(append(fK1XorFpart, fK2XorSpart...))
		// fmt.Println("IP-1 shufle result = ", PrintD(encryptResult))

		// fmt.Println("Encrypt result = ", encryptResult)
		result = append(result, byte(bToInt(encryptResult)))
	}
	// fmt.Println("Encryption result = ", result)
	// fmt.Println("Encrypted text = ", string(result))
	return result
}
func IPshufle(b []bool) []bool {
	return []bool{b[1], b[5], b[2], b[0], b[3], b[7], b[4], b[6]}
}
func IPmin1shufle(b []bool) []bool {
	return []bool{b[3], b[0], b[2], b[4], b[6], b[1], b[7], b[5]}
}

func f(xr []bool, key []bool) []bool {
	//extend xr to 8 bits
	// fmt.Println("Extending XR: ", PrintD(xr))
	xr = extension(xr)
	// fmt.Println("Extended XR: ", PrintD(xr))
	// xor
	// fmt.Println("XR xor K: ", PrintD(xr), PrintD(key))
	xr = xor(xr, key)
	// fmt.Println("XR xor K = ", PrintD(xr))
	// divide on 2 parts
	// fmt.Println("Dividing on 2 parts: ", PrintD(xr))
	fPart, sPart := divide(xr)
	// fmt.Println("First part: ", PrintD(fPart))
	// fmt.Println("Second part: ", PrintD(sPart))
	// s1 & s2
	fPart, sPart = s1(fPart), s2(sPart)
	// fmt.Println("S1 and S2 = ", PrintD(fPart), PrintD(sPart))
	// P
	// fmt.Println("F result: ", PrintD(reshafleP(fPart, sPart)))
	return reshafleP(fPart, sPart)

}
func reshafleP(f, s []bool) []bool {
	return []bool{f[1], s[1], s[0], f[0]}
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
	// fmt.Println("a14 = ", a14, "a23 = ", a23)
	// fmt.Println("s1[", bToInt(a14), "] = ", s1[bToInt(a14)])
	// fmt.Printf("%d = %b", s1[bToInt(a14)][bToInt(a23)], s1[bToInt(a14)][bToInt(a23)])
	s := fmt.Sprintf("%b", s1[bToInt(a14)][bToInt(a23)])
	// fmt.Println("s: ", s)
	if len(s) == 1 {
		s = "0" + s
	}
	// fmt.Println("Returning: ", s)
	return sToB(s)
}
func s2(b []bool) []bool {
	s2 := [][]int{
		{1, 1, 2, 3},
		{2, 0, 1, 3},
		{3, 0, 1, 0},
		{2, 1, 0, 3},
	}
	a14 := []bool{b[0], b[3]}
	a23 := []bool{b[1], b[2]}
	s := fmt.Sprintf("%b", s2[bToInt(a14)][bToInt(a23)])
	if len(s) == 1 {
		s = "0" + s
	}
	return sToB(s)
}
func sToB(s string) (result []bool) {
	for _, e := range s {
		if e == '1' {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return
}

func xor(f []bool, s []bool) []bool {
	// fmt.Println("XOR: ", PrintD(f), PrintD(s))
	result := make([]bool, 0)
	for i := range len(s) {
		if f[i] != s[i] {
			result = append(result, true)
			// fmt.Println("True XOR result: ", PrintD(result))
			continue
		}
		// fmt.Println("False XOR result: ", PrintD(result))
		result = append(result, false)
	}
	// fmt.Println("XOR result: ", PrintD(result))
	return result
}

func extension(b []bool) []bool {
	return []bool{b[3], b[0], b[1], b[2], b[1], b[2], b[3], b[1]}
}

func divide(b []bool) ([]bool, []bool) {
	// fmt.Printf("B: %d\n", bToInt(b))
	// fmt.Println(PrintD(b))
	return b[:len(b)/2], b[len(b)/2:]
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
	if len(bits) != 8 {
		for i := 0; i < 8-len(bits); i++ {
			bits = append(bits, false)
		}
	}
	slices.Reverse(bits)

	return bits
}
func (d *DES) P10() {
	p10Pattern := []int{3, 5, 2, 7, 4, 10, 1, 9, 8, 6}
	p10Result := make([]bool, 0)
	for _, i := range p10Pattern {
		i--
		// fmt.Println("Putting ", d.Bkey[i], " to ", i, " place")
		p10Result = append(p10Result, d.Bkey[i])
	}
	// fmt.Println("p10Result: ", p10Result)
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
	// fmt.Print("b = ", b)
	slices.Reverse(b)
	result := 0
	for i := len(b) - 1; i >= 0; i-- {
		result <<= 1
		if b[i] {
			result |= 1
		}
	}
	// fmt.Print("; result = ", result)
	// fmt.Println("bToInt return: ", result)
	return result
}
