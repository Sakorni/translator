package main

import (
	"fmt"
	"unicode"
)

const EN_LOCALE = "En"
const RU_LOCALE = "Ru"

func main() {
	word := "a"
	DetectLocale(string(word))
	tr, err := TranslateText("Ru", "cute")
	fmt.Println(tr, err)

}

func DetectLocale(word string) (string, error) {

	letter := rune(word[0])

	detect := 'a'

	detect++

	if !unicode.IsLetter(letter) {
		return "", fmt.Errorf("Words are should start with a letter")
	}

	if letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z' {
		return EN_LOCALE, nil
	}
	if letter >= 'А' && letter <= 'Я' || letter >= 'а' && letter <= 'я' {
		return RU_LOCALE, nil
	}
	return "", fmt.Errorf("Can't define a locale of a word {%s}.", word)
}
