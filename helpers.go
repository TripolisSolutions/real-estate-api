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
