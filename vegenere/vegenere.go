package vegenere

import (
	"fmt"
	"strings"
)

type Vegener struct {
	table    [][]rune
	keyWord  []rune
	message  []rune
	eMessage []rune
	keys     []int
}

func NewVegener(message, key string) *Vegener {
	v := &Vegener{
		table:    make([][]rune, 32),
		message:  []rune(strings.ToLower(message)),
		keyWord:  []rune(strings.ToLower(key)),
		eMessage: make([]rune, len(message)),
		keys:     make([]int, 0),
	}
	v.deleteSpaces()

	return v
}

var russianMapRuneInt = map[rune]int{'а': 0, 'б': 1, 'в': 2, 'г': 3, 'д': 4, 'е': 5, 'ж': 6, 'з': 7, 'и': 8, 'й': 9, 'к': 10, 'л': 11, 'м': 12, 'н': 13, 'о': 14, 'п': 15, 'р': 16, 'с': 17, 'т': 18, 'у': 19, 'ф': 20, 'х': 21, 'ц': 22, 'ч': 23, 'ш': 24, 'щ': 25, 'ъ': 26, 'ы': 27, 'ь': 28, 'э': 29, 'ю': 30, 'я': 31}
var russianSlice = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}

func Shift(s int, slice []rune) []rune {
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
	// for i := range v.table {
	// 	fmt.Println(string(v.table[i]))
	// }
}

func (v *Vegener) deleteSpaces() {
	nText := make([]rune, len(v.message))
	for i := range v.message {
		if v.message[i] != ' ' {
			nText[i] = v.message[i]
		}
	}
	// fmt.Println("Message without spaces: ", string(nText))
	v.message = nText
	nKey := make([]rune, 0)
	for i := range v.keyWord {
		if v.keyWord[i] != ' ' {
			nKey = append(nKey, v.keyWord[i])
		}
	}
	// fmt.Println("Key without spaces: ", string(nKey))
	v.keyWord = nKey
}

func (v *Vegener) CreateKey() {
	nKey := make([]rune, 0)
	for range len(v.message) / 2 {
		nKey = append(nKey, v.keyWord...)

	}
	// fmt.Println(len(nKey))
	// fmt.Println(len(v.keyWord))
	// fmt.Println("Key: ", nKey)
	v.keyWord = nKey[:len(v.message)-1]
}

func (v *Vegener) Encrypt() {
	// fmt.Println("message: ", string(v.message))
	for i := range len(v.message) - 1 {
		for j := range v.table {
			if v.keyWord[i] == v.table[j][0] {
				v.keys = append(v.keys, j)
				v.eMessage[i] = v.table[j][russianMapRuneInt[v.message[i]]]
			}

		}
	}

}
func (v *Vegener) PrintData() {
	fmt.Printf("    %s\n", string(russianSlice))
	for i := range v.table {
		if i < 10 {
			fmt.Printf("%d:  ", i)
			fmt.Println(string(v.table[i]))
		}
		if i > 10 {
			fmt.Printf("%d: ", i)
			fmt.Println(string(v.table[i]))
		}
	}
	// fmt.Println("Key length: ", len(v.keyWord))
	// fmt.Println("Message length: ", len(v.message))
	// fmt.Println("Encrypted message length: ", len(v.eMessage))
	// fmt.Println("Message:   ", string(v.message))
	fmt.Println("KeyWord:   ", string(v.keyWord))
	fmt.Println("Keys:      ", v.keys)
	fmt.Println("Encrypted: ", string(v.eMessage))
}
