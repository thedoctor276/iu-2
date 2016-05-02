package iu

import "strconv"

// HTMLEntities converts all applicable characters to HTML entities
func HTMLEntities(s string) string {
	var conv string

	for _, r := range s {
		conv += runeToHTMLEntity(r)
	}

	return conv
}

func runeToHTMLEntity(r rune) string {
	if r < 128 {
		return string(r)
	}

	return "&#" + strconv.FormatInt(int64(r), 10) + ";"
}
