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

func main() {
	ConnectDatabase()
}

func TranslateWord(input string) (output string) {
}

func ConsoleReader() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		//Todo: do something, idk
	}
	input = strings.ToLower(input)
	locale, err := GetTargetLocale(input)

	tr, err := TranslateText(locale, input)
	fmt.Println(tr, err)
}
