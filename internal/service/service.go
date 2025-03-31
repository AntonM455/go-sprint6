package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Функция Convert конвертирует переданный текст в Морзе; и наоборот — если был передан код Морзе, функция должна переконвертировать его в обычный текст и вернуть.
func Convert(StringTxtOrMorse string) (string, error) {
	if StringTxtOrMorse == "" {
		return "", errors.New("входящая строка пустая") // возвращаем ошибку в случае, если на вход пришла пустая строка
	}

	// Проверяем, состоит ли строка только из символов Морзе, они могут быть только пробелом, - или .
	isMorse := true
	for _, ch := range StringTxtOrMorse {
		if ch != '-' && ch != '.' && ch != ' ' {
			isMorse = false
			break
		}
	}

	// Выполняем конвертацию, если true, то вызовет ToText, если false - ToMorse

	if isMorse {
		return morse.ToText(StringTxtOrMorse), nil
	} else {
		return morse.ToMorse(StringTxtOrMorse), nil
	}
}
