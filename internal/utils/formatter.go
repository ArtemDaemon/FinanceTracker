package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatCurrency(value float64) string {
	p := message.NewPrinter(language.Russian)
	return p.Sprintf("%d ₽", int(value))
}

func GetColorClass(value float64) string {
	if value < 0 {
		return "color-red"
	}
	return "color-green"
}
