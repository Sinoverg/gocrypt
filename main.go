package main

import (
	"fmt"

	des "github.com/Cirqach/gocrypt/DES"
)

func main() {
	var a int = 909
	d := des.NewDES(a)
	d.GenerateKeys()
	// fmt.Printf("Key K1 = %s\nKey K2 = %s", des.PrintD(d.K1), des.PrintD(d.K2))
	var t string
	fmt.Scan(&t)
	b := []byte(t)
	fmt.Println(b)
	d.Encrypt(b)

}
