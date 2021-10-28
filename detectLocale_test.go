package main

import (
	"testing"
)

func TestEnLocales(t *testing.T) {
	engAlphabet := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	for _, letter := range engAlphabet {
		answer, err := DetectLocale(string(letter))
		if answer != EN_LOCALE {
			t.Fatalf("Wanted {%s}, but recieved {%s} with err {%s}", EN_LOCALE, answer, err.Error())
		}

	}
}
