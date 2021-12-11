package main

import (
	"bufio"
	"fmt"
	"golangDict/lib/core"
	db "golangDict/lib/database"
	tr "golangDict/lib/translation"
	"os"
	"strings"
)

func main() {
	ConsoleReader()
}

func ConsoleReader() {

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("Insert the word which you want to translate: ")
		input, err := reader.ReadString('\n')
		os.Stdout.Sync()
		if err != nil {
			fmt.Printf("Invalid input!\n %s\n", err.Error())
		}
		input = strings.ToLower(strings.TrimSpace(input))

		res := TranslateWord(input)

		fmt.Printf("%s -> %s \n", input, res)
	}
}

func TranslateWord(input string) (output string) {
	originalLocale, translateLocale, err := tr.GetTargetLocale(input)

	if err != nil {
		return err.Error()
	}
	word, err := db.GetWord(originalLocale, input)
	if err == nil {
		return word.GetResult(originalLocale)
	} else {
		if err.Error() == db.EMPTY_RESULT_CAPTION {
			res, translationErr := tr.TranslateText(translateLocale, input)
			if translationErr != nil {
				return translationErr.Error()
			}
			res = strings.ToLower(res)
			if originalLocale == core.RU_LOCALE {
				db.AddWord(input, res)
			} else {
				db.AddWord(res, input)
			}
			return res
		}
		return err.Error()
	}
}
