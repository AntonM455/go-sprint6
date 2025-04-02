package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Функция Convert конвертирует переданный текст в Морзе; и наоборот —
// если был передан код Морзе, функция должна переконвертировать его в
// обычный текст и вернуть.
func Convert(StringTxtOrMorse string) (string, error) {
	if StringTxtOrMorse == "" {
		// возвращаем ошибку в случае, если на вход пришла пустая строка
		return "", errors.New("input string is empty")
	}

	// Проверяем, состоит ли строка только из символов Морзе,
	// они могут быть только пробелом,"-" или "."
	if strings.ContainsFunc(StringTxtOrMorse, func(r rune) bool {
		return !strings.ContainsAny(string(r), ".- ") // если символ не '.', не '-', не пробел
	}) {
		// Если это обычный текст, конвертируем в Морзе
		return morse.ToMorse(StringTxtOrMorse), nil
	}

	// Если строка состоит только из символов Морзе, то выполняем конвертацию в текст
	return morse.ToText(StringTxtOrMorse), nil
}
