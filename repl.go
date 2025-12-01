package main

import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	words := strings.Fields(text)
	return words
}