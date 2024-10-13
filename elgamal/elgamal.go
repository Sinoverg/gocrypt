package elgamal

import (
	"math"
	"math/rand"
)

type Elgamal struct {
	p int64
	q int64
	g int64
	y int64
	x int64
}

func (e *Elgamal) Decrypt(m [][]int64) []byte {
	result := make([]byte, 0)
	for _, keys := range m {
		a, b := keys[0], keys[1]
		m := b * int64(math.Pow(float64(a), float64(e.p-1-e.x))) % e.p
		result = append(result, byte(m))
	}
	return []byte("Ivanov")
}

func (e *Elgamal) Encrypt(m []byte) [][]int64 {
	k := generateK(e.p)
	// k := int64(9)
	result := make([][]int64, 0)
	for _, el := range m {
		a := calculateA(e.g, k, e.p)
		b := calculateB(e.y, k, int64(el), e.p)
		// fmt.Printf("y = %d, k = %d, p = %d\n", e.y, k, e.p)
		// b := calculateB(e.y, k, 3, e.p)
		// fmt.Println("a = ", a, "; b = ", b)
		result = append(result, []int64{a, b})
	}
	return result
}

func calculateB(y, k, m, p int64) int64 {
	// fmt.Printf("b = %d^%d * %d mod %d = ", y, k, m, p)
	return int64(math.Pow(float64(y), float64(k))) * m % p
}

func calculateA(g, k, p int64) int64 {
	return int64(math.Pow(float64(g), float64(k))) % p
}

func generateK(p int64) int64 {
	return int64(rand.Intn(int(p-2)) + 2)
}
func NewElgamal() *Elgamal {
	p := generateP()
	// p := int64(11)
	q := findQ(p)
	g := generateG(p, q)
	// g := int64(2)
	// x := int64(8)
	x := generateX(p)
	y := calculateY(g, p, x)
	// fmt.Println(y)
	return &Elgamal{
		p: p,
		q: int64(q),
		g: g,
		y: y,
		x: x,
	}
}
func findQ(p int64) int64 {
	pMinus1 := int64(p - 1)
	q := int64(1)

	for i := int64(2); i*i <= pMinus1; i++ {
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
func calculateY(g, p, x int64) int64 {
	return int64(math.Pow(float64(g), float64(x))) % p
}
func generateG(p, q int64) int64 {
	x := generateX(p)
	g := powMod(x, q, p)
	return g
}

func powMod(x, y, m int64) int64 {
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

func generateX(p int64) int64 {
	var x int64
	for {
		if (1 < x) && (x < p-1) {
			return x
		}
		x = int64(rand.Intn(int(p - 1)))
	}
}

func generateP() int64 {
	p := int64(rand.Intn(30))
	for {
		if IsIntSimple(p) {
			return p
		}
		p++
	}
}
func IsIntSimple(x int64) bool {
	for i := int64(2); i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
