package trithemius

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	russianAlphabet = "абвгдежзийклмнопрстуфхцчшщъыьэюя"
	englishAlphabet = "abcdefghijklmnopqrstuvwxyz"
)

func Decrypt(tableSize, text, key string) (string, error) {
	table, err := createTable(tableSize, key)
	if err != nil {
		log.Println("Error in encrypt: " + err.Error())
		return "", err
	}
	log.Println("Table: ", table)
	encryptedText := text
	// for decrypt we need to use symbols which upper than our symbol in table
	// a b c d
	// ^
	// |
	// e f g h
	decryptedText := ""
	// go through created table
	for i := range table {
		for j := range table[i] {
			// go through encrypted text
			for k := range encryptedText {
				// if we find our symbol in table
				if table[i][j] == string(encryptedText[k]) {
					log.Println("Table element = element from encrypted text: ", table[i][j], encryptedText[k])
					// if we are not in last row
					if i < len(table)-1 {
						log.Println("i = ", i)
						decryptedText += table[i+1][j]
						// if we are in the first row - we can't use i+1 row, last row
					} else if i == 0 {
						log.Println("i = ", i)
						decryptedText += table[len(table)-1][j]
						// if something wrong
					} else {
						log.Println("WTF, i = ", i)
					}
				}
			}
		}
	}
	log.Println("Decrypted text: ", decryptedText)
	return decryptedText, nil

}

// TODO: create encrypt function
func Encrypt(tableSize, text, key string) (string, error) {
	return "", nil
}

func createTable(tableSize string, key string) ([][]string, error) {
	size := strings.Split(tableSize, "x")
	row, err := strconv.Atoi(size[0])
	if err != nil {
		log.Println("Error due creating table: " + err.Error())
		return [][]string{}, fmt.Errorf("Error due creating table: %s", err)
	}
	col, err := strconv.Atoi(size[1])
	if err != nil {
		log.Println("Error due creating table: " + err.Error())
		return [][]string{}, fmt.Errorf("Error due creating table: %s", err)
	}
	table := make([][]string, row)
	for i := range table {
		table[i] = make([]string, col)
	}
	log.Println("Creating table, table: ", table)
	splitedKey := strings.Split(key, "")
	k := 0
	alphabet := make(map[string]bool)

	// putting key symbols in table and add it to alphabet map
	for i := range table {
		for j := range table[i] {
			if k == len(splitedKey) {
				break
			}
			table[i][j] = splitedKey[k]
			alphabet[table[i][j]] = true
			k++
		}
	}
	log.Println("Table with key: ", table)

	// filling table with other symbols
	russianSymbols := strings.Split(russianAlphabet, "")
	symbolsToPut := make([]string, 0)

	// saving russian symbols in slice
	// and adding them to alphabet
	// this need for solve problem with position of symbols in table
	for _, symbol := range russianSymbols {
		if _, ok := alphabet[symbol]; !ok {
			alphabet[symbol] = true
			symbolsToPut = append(symbolsToPut, symbol)
		}
	}
	log.Println("Symbols to put: ", symbolsToPut)
	// now we fill table with other symbols
	for i := range table {
		for j := range table[i] {
			if table[i][j] == "" {
				table[i][j] = symbolsToPut[0]
				symbolsToPut = symbolsToPut[1:]
			}
		}
	}
	log.Println("Alphabet: ", alphabet)
	log.Println("Table: ", table)
	return table, nil
}
