package elgamal

import (
	"encoding/binary"
	"math"
	"math/rand"
)

type Elgamal struct {
	p int
	q int
	g int
	y int
	x int
}

func (e *Elgamal) Encrypt(m []byte) []byte {
	k := generateK(e.p)
	// a := powMod(e.g, k, e.p)
	b := powMod(e.y, k, e.p)

	ciphertext := make([]byte, len(m)*4) // Assuming 4 bytes for each integer
	for i := 0; i < len(m); i++ {
		// Convert byte to integer
		mInt := int(m[i])

		// Calculate ciphertext value
		cInt := powMod(b, mInt, e.p)

		// Convert integer back to bytes and store in ciphertext
		binary.BigEndian.PutUint32(ciphertext[i*4:], uint32(cInt))

	}

	return ciphertext
}

func generateK(p int) int {
	return rand.Intn(int(p-2)) + 1
}
func NewElgamal() *Elgamal {
	p := generateP()
	q := findQ(p)
	g := generateG(p, q)
	x := generateX(p)
	y := calculateY(g, p, x)
	return &Elgamal{
		p: p,
		q: q,
		g: g,
		y: y,
		x: x,
	}
}
func findQ(p int) int {
	pMinus1 := p - 1
	q := int(1)

	for i := int(2); i*i <= pMinus1; i++ {
		for pMinus1%i == 0 {
			q = i
			pMinus1 /= i
		}
	}

	if pMinus1 > 1 {
		q = pMinus1
	}

	return q
}
func calculateY(g, p, x int) int {
	return int(math.Pow(float64(g), float64(x))) % p
}
func generateG(p, q int) int {
	x := generateX(p)
	g := powMod(x, q, p)
	return g
}

func powMod(x, y, m int) int {
	if y == 0 {
		return 1
	}
	res := powMod(x, y/2, m)
	res = (res * res) % m
	if y%2 == 1 {
		res = (res * x) % m
	}
	return res
}

func generateX(p int) int {
	var x int
	for {
		if (int(1) < x) && (x < p-1) {
			return x
		}
		x = int(rand.Intn(int(p - 1)))
	}
}

func generateP() int {
	p := rand.Intn(100)
	for {
		if IsIntSimple(p) {
			return p
		}
		p++
	}
}
func IsIntSimple(x int) bool {
	for i := int(2); i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
