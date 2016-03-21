package iu

import "testing"

func brTest() *Br {
	return &Br{}

}

func TestTestBrSetDataContext(t *testing.T) {
	testComponentSetDataContext(t, brTest())
}

func TestBrInit(t *testing.T) {
	testComponentInit(t, brTest())
}

func TestBrClose(t *testing.T) {
	testComponentClose(t, brTest())
}

func TestBrRender(t *testing.T) {
	testComponentRender(t, brTest())
}

func TestBrSync(t *testing.T) {
	testComponentSync(t, brTest())
}
