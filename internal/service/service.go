package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetect(input string) (string, error) {
	if input == "" {
		return "", errors.New("Входная строка пустая")
	}
	LettersAndNumbers := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя0123456789"
	if strings.ContainsAny(input, LettersAndNumbers) {
		result := morse.ToMorse(input)
		return result, nil
	} else {
		result := morse.ToText(input)
		return result, nil
	}
}
