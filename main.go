package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Cirqach/gocrypt/vegenere"
)

func main() {
	var keyWord string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter key word: ")
	fmt.Scan(&keyWord)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	v := vegenere.NewVegener(text, keyWord)
	v.CreateTable(9)
	v.CreateKey()
	v.Encrypt()
	v.PrintData()
}

// fmt.Println(string(vegenere.Shift(5, []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'})))
// fmt.Println(caesar.EncryptWithShift("МЫ  ДОЛЖНЫ  ПРИЗНАТЬ  ОЧЕВИДНОЕ:  ПОНИМАЮТ  ЛИШЬ  ТЕ,  КТО  ХОЧЕТ  ПОНЯТЬ", shiftedSlice))
//
// fmt.Print("Enter keyA: ")
// fmt.Scan(&keyA)
// fmt.Println("Encrypted: ", string(caesar.EncryptAffineTable(caesar.CreateAffineTable(keyA, keyB), text)))
// fmt.Println(string(c.Encrypt("привет")))

// }
