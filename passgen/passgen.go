package passgen

import (
	"fmt"
	"math/rand"
)

const (
	Numbers = "0123456789"
	Special = "!@#$%^&*()_+"
)

func Generate(length int, symbols string) string {
	fmt.Printf("Generating password with length %d and symbols: %s\n", length, symbols)
	var password string
	for range length {
		rd := rand.Intn(3)
		// fmt.Printf("rd: %d\n", rd)
		switch rd {
		case 0:
			password += string(symbols[rand.Intn(len(symbols))])
		case 1:
			password += string(Numbers[rand.Intn(len(Numbers))])
		case 2:
			password += string(Special[rand.Intn(len(Special))])
		}
	}
	return password
}
