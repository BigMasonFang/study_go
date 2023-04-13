package funcs

import (
	"strings"
)

// titled func is exportable
func Cap(text string) string {
	return strings.Title(text)
}

// not titled func is not exportable
func repeat(text string, n int) string {
	const limit = 100
	return strings.Repeat(text, n)
}

func Swap(a, b int) (int, int) {
	return b, a
}
