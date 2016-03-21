package iu

import "testing"

func textTest() *Text {
	return &Text{
		Value:                "I love 42",
		OnDataContextChanged: func(dataCtx interface{}) {},
	}
}

func TestTestTextSetDataContext(t *testing.T) {
	testComponentSetDataContext(t, textTest())
}

func TestTextInit(t *testing.T) {
	testComponentInit(t, textTest())
}

func TestTextClose(t *testing.T) {
	testComponentClose(t, textTest())
}

func TestTextRender(t *testing.T) {
	testComponentRender(t, textTest())
}

func TestTextSync(t *testing.T) {
	testComponentSync(t, textTest())
}
