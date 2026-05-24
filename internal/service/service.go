package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(t string) (string, error) {
	t = strings.TrimSpace(t)
	if t == "" {
		return "", errors.New("Пустая строка")
	}
	clean := strings.Trim(t, ".- /")
	if clean == "" {
		text := morse.ToText(t)
		return text, nil
	}
	morseCode := morse.ToMorse(t)
	return morseCode, nil
}
