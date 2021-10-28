package main

import (
	"bufio"
	"fmt"
	"os"
)

const EN_LOCALE = "En"
const RU_LOCALE = "Ru"

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	locale, err := GetTargetLocale(input)

	tr, err := TranslateText(locale, input)
	fmt.Println(tr, err)

}
