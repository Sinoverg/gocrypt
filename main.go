package main

import (
	"github.com/Cirqach/gocrypt/vegenere"
)

func main() {
	// c := caesar.NewCaesar("ru", 5, "ЗИМА")
	// fmt.Println(string(caesar.EncryptAffineTable(c.CreateAffineTable(3, 2), "СМЫСЛ НАШЕЙ ЖИЗНИ – НЕПРЕРЫВНОЕ ДВИЖЕНИЕ")))
	// fmt.Println(string(c.Encrypt("привет")))

	// fmt.Println(string(vegenere.Shift(5, []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'})))
	// shiftedSlice := caesar.Shift(9, []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'})
	// fmt.Println(string(shiftedSlice))
	// fmt.Println(caesar.EncryptWithShift("МЫ  ДОЛЖНЫ  ПРИЗНАТЬ  ОЧЕВИДНОЕ:  ПОНИМАЮТ  ЛИШЬ  ТЕ,  КТО  ХОЧЕТ  ПОНЯТЬ", shiftedSlice))
	v := vegenere.NewVegener("Яблоки в говне", "ЗИМА")
	// v.CreateTable(9)
	v.CreateKey()
}
