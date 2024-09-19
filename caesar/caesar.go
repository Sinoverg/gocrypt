package caesar

import (
	"fmt"
	"strings"
)

var (
	russianSlice = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
	russianMap   = map[rune]int{
		'а': 0,
		'б': 1,
		'в': 2,
		'г': 3,
		'д': 4,
		'е': 5,
		'ж': 6,
		'з': 7,
		'и': 8,
		'й': 9,
		'к': 10,
		'л': 11,
		'м': 12,
		'н': 13,
		'о': 14,
		'п': 15,
		'р': 16,
		'с': 17,
		'т': 18,
		'у': 19,
		'ф': 20,
		'х': 21,
		'ц': 22,
		'ч': 23,
		'ш': 24,
		'щ': 25,
		'ъ': 26,
		'ы': 27,
		'ь': 28,
		'э': 29,
		'ю': 30,
		'я': 31,
	}
)

type Caesar struct {
	Table   map[int]rune
	key     int
	keyWord []rune
}

// TODO: delete print functions
func PrintIntRune(table map[int]rune) string {
	result := ""
	for i := 0; i < 32; i++ {
		result += fmt.Sprintf("%d:%s ", i, string(table[i]))
	}
	return result
}
func PrintRuneInt(table map[rune]int) string {
	result := ""
	for i := range table {
		result += fmt.Sprintf("%d:%s ", table[i], string(i))
	}
	return result
}
func NewCaesar(lang string, key int, keyWord string) *Caesar {
	lang = strings.ToLower(lang)
	keyWord = strings.ToLower(keyWord)
	return &Caesar{
		Table:   make(map[int]rune),
		key:     key,
		keyWord: []rune(keyWord),
	}
}

func (c *Caesar) Encrypt(text string) string {
	text = strings.ToLower(text)
	// textRunes := []rune(text)
	// fmt.Println("Text: ", textRunes)
	encrypted := ""
	for _, v := range text {
		// fmt.Println(string(v))
		if num, ok := russianMap[v]; ok {
			// fmt.Println("num = ", num, " ; c.Table[num] = ", c.Table[num])
			encrypted += string(c.Table[num])
			// fmt.Println("encrypted = ", encrypted)
		} else {
			encrypted += string(v)
		}
	}
	return encrypted
}

func (c *Caesar) CreateTable() {
	// fmt.Printf("Now table: %v\n", PrintIntRune(c.Table))
	// create pointer to key postion in alphabet
	counter := c.key - 1
	// fmt.Println("KeyWord: ", string(c.keyWord), "; counter = ", counter)
	newTable := make(map[int]rune)
	alredyIn := make(map[rune]bool, 0)
	for i := range c.keyWord {
		if c.keyWord[i] == ' ' {
			continue
		}
		if _, ok := alredyIn[c.keyWord[i]]; ok {
			// fmt.Printf("Counter: %d; Symbol %s already in table: %v\n", counter, string(c.keyWord[i]), PrintIntRune(newTable))
			continue
		}
		// fmt.Printf("%s added to table: %v\n", string(c.keyWord[i]), PrintIntRune(newTable))
		newTable[counter] = c.keyWord[i]
		alredyIn[c.keyWord[i]] = true
		counter++
	}
	// fmt.Printf("New table: %v\n", PrintIntRune(newTable))
	// fmt.Printf("Alredy in table: %v\n", alredyIn)

	// filling table with other symbols
	for i := range russianSlice {
		if _, ok := alredyIn[russianSlice[i]]; !ok {
			if counter >= 32 {
				counter = 0
			}
			newTable[counter] = russianSlice[i]
			counter++
		}
	}
	// fmt.Printf("Table with other symbols: %v\n", PrintIntRune(newTable))
	c.Table = newTable
}
