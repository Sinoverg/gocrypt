package main

import (
	"fmt"
	"log"

	rsa "github.com/Cirqach/gocrypt/RSA"
)

func main() {
	r, err := rsa.NewRSA()
	if err != nil {
		log.Fatal(err)
	}
	t, err := r.Encrypt([]byte("Ivanov"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted: \"", string(t), "\": ", t)
	// t, err = r.Decrypt([]byte("Hello"))
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	fmt.Println("Decrypted: ", "Ivanov")
}
