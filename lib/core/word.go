package core

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

const EN_LOCALE = "En"
const RU_LOCALE = "Ru"
