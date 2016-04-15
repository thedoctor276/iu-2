package iu

import "testing"

func TestNewPage(t *testing.T) {
	c := PageConfig{
		Title: "Page test",
	}

	v := &World{}

	p := NewPage(v, c)
	defer p.Close()

	if conf := p.Config(); conf.Title != c.Title {
		t.Errorf("conf should be %v: %v", c, conf)
	}

	if conf := p.Config(); conf.Title != c.Title {
		t.Errorf("conf should be %v: %v", c, conf)
	}

	if mainView := p.MainView(); mainView != v {
		t.Errorf("mainView should be %v: %v", v, mainView)
	}
}

func TestNewPageWithLoopView(t *testing.T) {
	v := &HelloLoop{
		Greeting: &World{},
	}

	v.Parent = v

	p := NewPage(v, PageConfig{
		Title: "Page test",
	})

	defer p.Close()
}

func TestPageRender(t *testing.T) {
	v := &Hello{
		Greeting: &World{},
	}

	p := NewPage(v, PageConfig{
		Title: "Page test",
	})
	defer p.Close()

	p.Context = &EmptyContext{}
	t.Log(p.Render())
}
