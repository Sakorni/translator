package translation

import (
	. "golangDict/lib/core"
	"testing"
)

func TestEnLocale(t *testing.T) {
	engAlphabet := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	for _, letter := range engAlphabet {
		answer, err := DetectLocale(string(letter))
		if answer != EN_LOCALE {
			t.Fatalf("Wanted {%s}, but recieved {%s} with err {%s}", EN_LOCALE, answer, err.Error())
		}
	}
}

func TestRuLocale(t *testing.T) {
	rusAlphabet := "АаБбВвГгДдЕеЁёЖжЗзИиЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯя"
	for _, letter := range rusAlphabet {
		answer, err := DetectLocale(string(letter))
		if answer != RU_LOCALE {
			t.Fatalf("Wanted {%s}, but recieved {%s} with err {%s}", RU_LOCALE, answer, err.Error())
		}
	}
}
