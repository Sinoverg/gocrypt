package rsa

import (
	"fmt"
	"math"
	"math/rand"
)

type RSA struct {
	p, q int64
	n    int64
	f    int64
	e    int64
	d    int64
}

func NewRSA() (*RSA, error) {
	// fmt.Println("Generating keys...")
	p, q := generatePQ()
	// fmt.Println("p = ", p, " q = ", q)
	n := p * q
	// fmt.Println("n = ", n)
	f := (p - 1) * (q - 1)
	// fmt.Println("f = ", f)
	e := generateE(f)
	if e == 0 {
		return nil, fmt.Errorf("Error due generating e")
	}
	// fmt.Println("e = ", e)
	d := calculateD(e, f)
	if d == 0 {
		return nil, fmt.Errorf("Error due generating d")
	}
	// fmt.Println("d = ", d)
	// p, q := int64(3), int64(7)
	// n := int64(21)
	// f := int64(12)
	// e := int64(5)
	// d := int64(17)
	return &RSA{
		p: p,
		q: q,
		n: n,
		f: f,
		e: e,
		d: d,
	}, nil
}
func powInt(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}
func (r *RSA) Encrypt(text []byte) ([]byte, error) {
	// fmt.Println("Encrypting...")
	result := make([]byte, len(text))
	for i, t := range text {
		b := powInt(int64(t), r.e)
		// fmt.Println("b = ", b)
		result[i] = byte(int64(b) % r.n)
		// fmt.Println("result[i] = ", result[i])
		// fmt.Println("Encrypted symbol(", t, ":", string(t), ") = ", result[i], ":", string(result[i]))
	}
	return result, nil
}

func (r *RSA) Decrypt(text []byte) ([]byte, error) {
	// text = []byte{8, 6, 13, 18, 15, 1, 3}
	// fmt.Println("Decrypting...")
	result := make([]byte, len(text))
	for i, t := range text {
		// fmt.Println("powInt(float64(t), float64(r.d)) % r.n = ", int64(math.Pow(float64(t), float64(r.d))), " % ", r.n, " = ", int64(math.Pow(float64(t), float64(r.d)))%r.n)
		result[i] = byte(int64(powInt(int64(t), r.d)) % r.n)
		// fmt.Println("result[i] = ", result[i])
		// fmt.Println("Decrypted symbol(", t, ":", string(t), ") = ", result[i], ":", string(result[i]))
	}
	return result, nil
}

func calculateD(e, f int64) int64 {
	// fmt.Println("Calculating d...")
	for d := int64(2); d <= f; d++ {
		// fmt.Println("i = ", d, "(", d, "*", e, ")%", f, " = ", (d*e)%f)
		if (d*e)%f == 1 && d != e {
			return d
		}
	}
	return 0
}

func generateE(fN int64) int64 {
	// fmt.Println("Generating e...")
	for i := int64(0); i < fN; i++ {
		if isItSimple(i) && i != 2 {
			return i
		}
	}
	return 0
}

func generatePQ() (p, q int64) {
	// fmt.Println("Generating p and q...")
	x, y := int64(rand.Intn(100)), int64(rand.Intn(100))
	// fmt.Println("x = ", x, " y = ", y)
	for !isItSimple(x) {
		x++
	}
	for !isItSimple(y) {
		y++
	}
	p, q = x, y
	return
}

func isItSimple(x int64) bool {
	// fmt.Println("Checking ", x)
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	for i := int64(2); i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
