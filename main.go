package main

import (
	"fmt"

	"github.com/Cirqach/gocrypt/elgamal"
)

func main() {
	e := elgamal.NewElgamal()
	msg := e.Encrypt([]byte("Hello"))
	fmt.Println(msg)
	fmt.Println(string(msg))
}
