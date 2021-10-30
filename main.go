package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DBWord struct {
	ID            int64
	EnTranslation string
	RuTranslation string
	AppealCounter int64
}

/*
Returns a correct result for translating function.

For example: If you've typed a word in RU language it'll
return you an English translation of this words and vice versa
*/
func (word DBWord) GetResult(locale string) string {
	if locale == RU_LOCALE {
		return word.EnTranslation
	}
	return word.RuTranslation
}

func main() {
	ConnectDatabase()
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
		input = strings.ToLower(strings.Trim(input, " \r\n"))

		res := TranslateWord(input)

		fmt.Printf("%s -> %s \n", input, res)
	}
}

func TranslateWord(input string) (output string) {
	originalLocale, translateLocale, err := GetTargetLocale(input)

	if err != nil {
		return err.Error()
	}
	word, err := GetWord(originalLocale, input)
	if err == nil {
		return word.GetResult(originalLocale)
	} else {
		if err.Error() == EMPTY_RESULT_CAPTION {
			res, translationErr := TranslateText(translateLocale, input)
			if translationErr != nil {
				return translationErr.Error()
			}
			res = strings.ToLower(res)
			if originalLocale == RU_LOCALE {
				AddWord(input, res)
			} else {
				AddWord(res, input)
			}
			return res
		}
		return err.Error()
	}
}
