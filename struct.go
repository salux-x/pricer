package pricer

import (
	"math/big"
	"sync"
)

type Price struct {
	default_type string

	Price_source string
	Price        string
	Price_type   string
	Price_rat    *big.Rat
}

// [source][destination]ratio
type ConvertPriceType struct {
	sync.RWMutex
	d map[string]map[string]*big.Rat
}

type SearchPriceType struct {
	Search         string
	Result         string
	Description_ru string
}

var PriceTypes = []SearchPriceType{
	{"р", "RUB", "Рубль"},
	{string('\u20BD'), "RUB", "Рубль"}, // char "₽"
	{"€", "EUR", "Евро"},
	{"$", "USD", "Доллар США"},
	{"EUR", "EUR", "Евро"},
	{"USD", "USD", "Доллар США"},
	{"AUD", "AUD", "Австралийский доллар"},
	{"GBP", "GBP", "Фунт стерлингов Соединенного королевства"},
	{"BYR", "BYR", "Белорусских рублей"},
	{"Br", "BYR", "Белорусских рублей"},
	{"DKK", "DKK", "Датских крон"},
	{"KZT", "KZT", "Казахстанских тенге"},
	{"CAD", "CAD", "Канадский доллар"},
	{"CNY", "CNY", "Китайских юаней"},
	{"NOK", "NOK", "Норвежских крон"},
	{"XDR", "XDR", "СДР (специальные права заимствования)"},
	{"SGD", "SGD", "Сингапурский доллар"},
	{"TRY", "TRY", "Турецкая лира"},
	{"UAH", "UAH", "Украинских гривен"},
	{"SEK", "SEK", "Шведских крон"},
	{"CHF", "CHF", "Швейцарский франк"},
	{"JPY", "JPY", "Японских иен"},
	{"AZN", "AZN", "Азербайджанский манат"},
	{"AMD", "AMD", "Армянских драмов"},
	{"BGN", "BGN", "Болгарский лев"},
	{"BRL", "BRL", "Бразильский реал"},
	{"HUF", "HUF", "Венгерских форинтов"},
	{"INR", "INR", "Индийских рупий"},
	{"KGS", "KGS", "Киргизских сомов"},
	{"MDL", "MDL", "Молдавских леев"},
	{"PLN", "PLN", "Польский злотый"},
	{"RON", "RON", "Румынский лей"},
	{"TJS", "TJS", "Таджикских сомони"},
	{"TMT", "TMT", "Новый туркменский манат"},
	{"UZS", "UZS", "Узбекских сумов"},
	{"CZK", "CZK", "Чешских крон"},
	{"Kč", "CZK", "Чешских крон"},
	{"ZAR", "ZAR", "Южноафриканских рэндов"},
	{"KRW", "KRW", "Вон Республики Корея"},
}
