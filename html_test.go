package iu

import "testing"

func TestToHTMLEntities(t *testing.T) {
	s := "<div>j’aime les filles</div>"
	ToHTMLEntities(s)
}
