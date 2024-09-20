package vegenere

import (
	"fmt"
	"strings"
)

type Vegener struct {
	table   [][]rune
	key     []rune
	message []rune
}

func NewVegener(message, key string) *Vegener {
	v := &Vegener{
		table:   make([][]rune, 32),
		message: []rune(strings.ToLower(message)),
		key:     []rune(strings.ToLower(key)),
	}
	v.deleteSpaces()

	return v
}

var russianSlice = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}

func Shift(s int, slice []rune) []rune {
	// new slice for swap
	nSlice := make([]rune, len(slice))
	for i := range nSlice {
		if i == len(slice)-1 {
			nSlice[i] = slice[0]
			continue
		}
		nSlice[i] = slice[i+1]
	}
	if s > 1 {
		return Shift(s-1, nSlice)
	}
	return nSlice
}

func (v *Vegener) CreateTable(shift int) {
	for i := 0; i < 32; i++ {
		v.table[i] = Shift(i, russianSlice)
		// fmt.Println(string(v.table[i]))
	}
	v.table[0] = russianSlice
	for i := range v.table {
		fmt.Println(string(v.table[i]))
	}
}

func (v *Vegener) deleteSpaces() {
	nText := make([]rune, len(v.message))
	for i := range v.message {
		if v.message[i] != ' ' {
			nText[i] = v.message[i]
		}
	}
	fmt.Println("Message without spaces: ", string(nText))
	v.message = nText
	nKey := make([]rune, len(v.message))
	for i := range v.key {
		if v.key[i] != ' ' {
			nKey[i] = v.key[i]
		}
	}
	fmt.Println("Key without spaces: ", string(nKey))
	v.message = nKey
}

func (v *Vegener) CreateKey() {
	nKey := make([]rune, len(v.message))
	fmt.Printf("Length nkey: %d, len messasge: %d\n", len(nKey), len(v.message))
	for i := range v.message {
		nKey[i] = v.key[i%len(v.key)]
	}
	fmt.Println(string(nKey))
}

func (v *Vegener) Encrypt(text string) {

	// tSlice := []rune(text)
	// encrypted := ""
	// for i := range tSlice {
	//
	// }

}
