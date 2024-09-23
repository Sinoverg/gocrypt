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

}
