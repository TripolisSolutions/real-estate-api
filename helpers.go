package main

import "strconv"

func ParseIntWithFallback(v string, def int) int {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return def
	}
	return int(i)
}

var SupportedLanguages = []string{"vietnamese", "english"}

func isLanguageSupported(language string) bool {
	for _, lang := range SupportedLanguages {
		if lang == language {
			return true
		}
	}

	return false
}

func langCodeToLanguage(lang string) string {
	switch lang {
	case "vi":
		fallthrough
	case "vi-VN":
		return "vietnamese"
	default:
		return "english"
	}
}

var SupportedCurrencies = []string{"VND", "USD"}

func isCurrencySupported(currerncy string) bool {
	for _, cur := range SupportedCurrencies {
		if currerncy == cur {
			return true
		}
	}

	return false
}
