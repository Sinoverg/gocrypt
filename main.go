package main

import (
	"fmt"

	des "github.com/Cirqach/gocrypt/DES"
)

func main() {
	var a int
	fmt.Scan(&a)
	d := des.NewDES(a)
	d.Encrypt()
	fmt.Printf("Key K1 = %s\nKey K2 = %s", des.PrintD(d.K1), des.PrintD(d.K2))

}
