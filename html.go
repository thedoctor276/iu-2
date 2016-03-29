package iu

import "strconv"

func ToHTMLEntities(line string) string {
	var converted string = ""
	for _, rune_value := range line {
		converted += RuneToHTMLEntity(rune_value)
	}
	return converted
}

func RuneToHTMLEntity(entity rune) string {
	if entity < 128 {
		return string(entity)
	} else {
		return "&#" + strconv.FormatInt(int64(entity), 10) + ";"
	}
}
