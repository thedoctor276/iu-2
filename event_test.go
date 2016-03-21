package iu

import "testing"

func TestTryCallViewEvent(t *testing.T) {
	var ctx = &EmptyContext{}
	var msg = "It works"
	var result string
	var onLoaded = func() { result = msg }

	var view = &Page{
		Title:    "Test",
		OnLoaded: onLoaded,
	}

	view.Init(ctx)
	TryCallViewEvent(view, "OnLoaded")

	if result != msg {
		t.Errorf("result should be %v: %v", msg, result)
	}
}

func TestTryCallComponentEvent(t *testing.T) {
	var ctx = &EmptyContext{}
	var arg = MouseEvent{}
	var msg = "It works"
	var result string
	var onClick = func(event MouseEvent) { result = msg }

	var view = &Page{
		Title: "Test",
		Body: []Component{
			&Div{
				OnClick: onClick,
			},
		},
	}

	view.Init(ctx)
	TryCallComponentEvent(view.Body[0], "OnClick", arg)

	if result != msg {
		t.Errorf("result should be %v: %v", msg, result)
	}
}

func TestTryCallEventNil(t *testing.T) {
	var ctx = &EmptyContext{}

	var view = &Page{
		Title: "Test",
		Body: []Component{
			&Div{},
		},
	}

	view.Init(ctx)
	TryCallComponentEvent(view.Body[0], "OnClick", nil)
}

func TestTryCallEventInvalidArg(t *testing.T) {
	var ctx = &EmptyContext{}
	var msg = "It works"
	var result string
	var onClick = func(event MouseEvent) { result = msg }

	var view = &Page{
		Title: "Test",
		Body: []Component{
			&Div{
				OnClick: onClick,
			},
		},
	}

	defer func() { recover() }()

	view.Init(ctx)
	TryCallComponentEvent(view.Body[0], "OnClick", "Boo")
	t.Error("should have panic")
}

func TestTryCallEventDifferentArgLen(t *testing.T) {
	var ctx = &EmptyContext{}
	var arg = MouseEvent{}
	var msg = "It works"
	var result string
	var onClick = func(event MouseEvent) { result = msg }

	var view = &Page{
		Title: "Test",
		Body: []Component{
			&Div{
				OnClick: onClick,
			},
		},
	}

	defer func() { recover() }()

	view.Init(ctx)
	TryCallComponentEvent(view.Body[0], "OnClick", arg, "Boo")
	t.Error("should have panic")
}

func TestTryCallInvalidEvent(t *testing.T) {
	var ctx = &EmptyContext{}

	var view = &Page{
		Title: "Test",
		Body: []Component{
			&Div{},
		},
	}

	defer func() { recover() }()

	view.Init(ctx)
	TryCallComponentEvent(view.Body[0], "Title")
	t.Error("should have panic")
}

func TestTryCallComponentEventWithArg(t *testing.T) {
	var JSONArg = `{"AltKey":false,"Button":0,"ClientX":627,"ClientY":344,"CtrlKey":false,"Detail":1,"MetaKey":false,"PageX":627,"PageY":344,"ScreenX":-1489,"ScreenY":557,"ShiftKey":false}`
	var ctx = &EmptyContext{}
	var ok bool
	var onClick = func(event MouseEvent) { ok = true }

	var view = &Page{
		Title: "Test",
		Body: []Component{
			&Div{
				OnClick: onClick,
			},
		},
	}

	view.Init(ctx)
	tryCallComponentEventWithArg(view.Body[0], "OnClick", JSONArg)

	if !ok {
		t.Error("ok should be true")
	}
}
