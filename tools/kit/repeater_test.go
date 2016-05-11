package kit

import (
	"testing"

	"github.com/maxence-charriere/iu"
)

type SimpleText struct {
	Text string
}

func (t *SimpleText) Template() string {
	return `<p>{{.Text}}</p>`
}

func TestRepeaterSetSource(t *testing.T) {
	r := NewRepeater(func(src interface{}) iu.Component {
		return &SimpleText{
			Text: src.(string),
		}
	})

	d := iu.NewDriverTest(r, iu.DriverConfig{})
	defer d.Close()

	src := []string{
		"Maxence",
		"Achille",
		"Taesung",
		"Nam",
		"Damien",
	}

	r.SetSource(src)
	iu.RenderComponent(r)
}

func TestRepeaterSetSourceNotSlice(t *testing.T) {
	defer func() { recover() }()

	r := NewRepeater(func(src interface{}) iu.Component {
		return &SimpleText{
			Text: src.(string),
		}
	})

	d := iu.NewDriverTest(r, iu.DriverConfig{})
	defer d.Close()

	src := "Maxence"

	r.SetSource(src)
	iu.RenderComponent(r)
	t.Error("should have panic")
}

func TestNewRepeater(t *testing.T) {
	r := NewRepeater(func(src interface{}) iu.Component {
		return &SimpleText{
			Text: src.(string),
		}
	})

	d := iu.NewDriverTest(r, iu.DriverConfig{})
	defer d.Close()
}
