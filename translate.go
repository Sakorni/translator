package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TranslateText(targetLanguage, text string) (string, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("secrets/translateKey.json")

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx, opt)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)

	fmt.Println(client.SupportedLanguages(ctx, lang))

	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}