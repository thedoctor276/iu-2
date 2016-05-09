package iu

import "testing"

func TestHTMLEntities(t *testing.T) {
	s := "<div>j’aime les filles</div>"
	res := HTMLEntities(s)
	t.Log(res)
}
