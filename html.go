package iu

import "strconv"

func ToHTMLEntities(line string) string {
	var converted string

	for _, r := range line {
		converted += runeToHTMLEntity(r)
	}

	return converted
}

func runeToHTMLEntity(r rune) string {
	if r < 128 {
		return string(r)
	}

	return "&#" + strconv.FormatInt(int64(r), 10) + ";"
}
