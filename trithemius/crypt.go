package trithemius

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	russianAlphabet = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
)

func Decrypt(tableSize, text, key string) (string, error) {
	table, err := createTable(tableSize, key)
	if err != nil {
		log.Println("Error in encrypt: " + err.Error())
		return "", err
	}
	// printTable(table)
	encryptedText := []rune(text)
	// log.Println("Encrypted text: ", string(encryptedText))
	// for decrypt we need to use symbols which upper than our symbol in table
	// a b c d
	// ^
	// |
	// e f g h
	decryptedText := ""
	// go through created table
	for k := range encryptedText {
		for i := range table {
			for j := range table[i] {
				// go through encrypted text
				// if we find our symbol in table
				if table[i][j] == encryptedText[k] {
					// log.Println("Processing ", string(encryptedText[k]), " symbol; table item = ", string(table[i][j]))
					// if we are not in last row
					if i == 0 {
						// log.Println("first row; i = ", i)
						decryptedText += string(table[len(table)-1][j])
						// log.Println("Adding new item to decrypted text: ", decryptedText)
						continue
					}
					decryptedText += string(table[i-1][j])
					// log.Println("Adding new item to decrypted text: ", decryptedText)
					encryptedText = encryptedText[0:]
				}
			}
		}
	}
	// log.Println("Decrypted text: ", decryptedText)
	return decryptedText, nil

}

func Encrypt(tableSize, text, key string) (string, error) {
	table, err := createTable(tableSize, key)
	if err != nil {
		log.Println("Error in encrypt: " + err.Error())
		return "", err
	}
	// printTable(table)
	decryptedText := []rune(text)
	// log.Println("Encrypted text: ", string(encryptedText))
	// for decrypt we need to use symbols which upper than our symbol in table
	// a b c d
	// ^
	// |
	// e f g h
	encryptedText := ""
	// go through created table
	for k := range decryptedText {
		for i := range table {
			for j := range table[i] {
				// go through encrypted text
				// if we find our symbol in table
				if table[i][j] == decryptedText[k] {
					// log.Println("Processing ", string(encryptedText[k]), " symbol; table item = ", string(table[i][j]))
					// if we are not in last row
					if i == len(table)-1 {
						// log.Println("first row; i = ", i)
						encryptedText += string(table[0][j])
						// log.Println("Adding new item to decrypted text: ", decryptedText)
						continue
					}
					encryptedText += string(table[i+1][j])
					// log.Println("Adding new item to decrypted text: ", decryptedText)
					decryptedText = decryptedText[0:]
				}
			}
		}
	}
	// log.Println("Decrypted text: ", decryptedText)
	return encryptedText, nil

}

func createTable(tableSize string, key string) ([][]rune, error) {
	size := strings.Split(tableSize, "x")
	row, err := strconv.Atoi(size[0])
	if err != nil {
		log.Println("Error due creating table: " + err.Error())
		return [][]rune{}, fmt.Errorf("Error due creating table: %s", err)
	}
	col, err := strconv.Atoi(size[1])
	if err != nil {
		log.Println("Error due creating table: " + err.Error())
		return [][]rune{}, fmt.Errorf("Error due creating table: %s", err)
	}
	table := make([][]rune, row)
	for i := range table {
		table[i] = make([]rune, col)
	}
	// printTable(table)
	splitedKey := []rune(key)
	k := 0
	alphabet := make(map[rune]bool)

	// putting key symbols in table and add it to alphabet map
	for i := range table {
		for j := range table[i] {
			if k == len(splitedKey) {
				break
			}
			// log.Println("i = ", i, " j = ", j, " k = ", k, "symbol = ", string(splitedKey[k]), " table[i][j] = ", string(table[i][j]))
			if _, ok := alphabet[splitedKey[k]]; ok {
				// log.Println("Symbol alredy in table: ", string(table[i][j]))
				k++
				continue
			}
			alphabet[splitedKey[k]] = true
			table[i][j] = splitedKey[k]
			k++

		}
	}
	// log.Println("Table with key: ", table)

	// filling table with other symbols
	symbolsToPut := make([]rune, 0)

	// saving russian symbols in slice
	// and adding them to alphabet
	// this need for solve problem with position of symbols in table
	for _, symbol := range russianAlphabet {
		if _, ok := alphabet[symbol]; !ok {
			alphabet[symbol] = true
			symbolsToPut = append(symbolsToPut, symbol)
		}
	}
	// log.Println("Symbols to put: ", symbolsToPut)
	// now we fill table with other symbols
	for i := range table {
		for j := range table[i] {
			if table[i][j] == 0 {
				table[i][j] = symbolsToPut[0]
				symbolsToPut = symbolsToPut[1:]
			}
		}
	}
	// log.Println("Alphabet: ", alphabet)
	// log.Println("Table: ", table)
	return table, nil
}

func printTable(table [][]rune) {
	for i := range table {
		for j := range table[i] {
			fmt.Print(string(table[i][j]) + " ")
		}
		fmt.Println()
	}
}
