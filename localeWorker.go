package main

import (
	"fmt"
	"unicode"
)

func DetectLocale(word string) (string, error) {

	letter := []rune(word)[0]

	detect := 'А'
	detect2 := 'A'

	detect, detect2 = detect2, detect

	if !unicode.IsLetter(letter) {
		return "", fmt.Errorf("Word %s started not with a letter", word)
	}

	if letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z' {
		return EN_LOCALE, nil
	}
	if letter >= 'А' && letter <= 'Я' || letter >= 'а' && letter <= 'я' || letter == 'Ё' || letter == 'ё' {
		return RU_LOCALE, nil
	}
	return "", fmt.Errorf("Can't define a locale of a word {%s}.", word)
}

func GetTargetLocale(word string) (locale string, err error) {
	locale, err = DetectLocale(word)
	if err != nil {
		return
	}
	locale = targetLocale(locale)
	return
}

func targetLocale(from string) (to string) {
	if from == EN_LOCALE {
		return RU_LOCALE
	} else if from == RU_LOCALE {
		return EN_LOCALE
	} else {
		panic("Incorrect input in TargetLocale!")
	}
}
