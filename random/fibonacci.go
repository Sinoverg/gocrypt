package random

import (
	"fmt"
	"log"
)

func RandomFibonacci(a, b int, row []float32) []float32 {
	var k int
	i := len(row)
	for k < 10 {
		if row[i-a] >= row[i-b] {
			row = append(row, row[i-a]-row[i-b])
			fmt.Printf("k[%d] = k[%d] - k[%d] = %f\n", i, i-a, i-b, row[i-a]-row[i-b])
			k++
			i++
			continue
		}
		if row[i-a] < row[i-b] {
			row = append(row, row[i-a]-row[i-b]+1)
			fmt.Printf("k[%d] = k[%d] - k[%d] + 1 = %f\n", i, i-a, i-b, row[i-a]-row[i-b]+1)
			i++
			k++
			continue
		}
		log.Println("Something wrong")
		i++
		k++
	}
	return row
}
