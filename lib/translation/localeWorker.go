package translation

import (
	"fmt"
	c "golangDict/lib/core"
	"unicode"
)

func DetectLocale(word string) (string, error) {

	letter := []rune(word)[0]

	if !unicode.IsLetter(letter) {
		return "", fmt.Errorf("word %s started not with a letter", word)
	}

	if letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z' {
		return c.EN_LOCALE, nil
	}
	if letter >= 'А' && letter <= 'Я' || letter >= 'а' && letter <= 'я' || letter == 'Ё' || letter == 'ё' {
		return c.RU_LOCALE, nil
	}
	return "", fmt.Errorf("can't define a locale of a word {%s}", word)
}

func GetTargetLocale(word string) (originalLocale, translateLocale string, err error) {
	originalLocale, err = DetectLocale(word)
	if err != nil {
		return
	}
	translateLocale = targetLocale(originalLocale)
	return
}

func targetLocale(from string) (to string) {

	if from == c.EN_LOCALE {
		return c.RU_LOCALE
	} else if from == c.RU_LOCALE {
		return c.EN_LOCALE
	} else {
		panic("Incorrect input in TargetLocale!")
	}
}
