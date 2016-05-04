package iu

import "testing"

func TestHTMLEntities(t *testing.T) {
	s := "<div>j’aime les filles</div>"
	HTMLEntities(s)
}
