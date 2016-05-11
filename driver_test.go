package iu

import "testing"

func TestDriverBaseRender(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{}
	d := NewDriverTest(root, conf)
	defer d.Close()

	t.Log(d.Render())
}

func TestDriverBaseRenderWithConf(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{
		ID:   "Test",
		Lang: "fr",
		CSS:  []string{"test.css"},
		JS:   []string{"test.js"},
	}
	d := NewDriverTest(root, conf)
	defer d.Close()

	t.Log(d.Render())
}

func TestNewDriverBase(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{}
	d := NewDriverBase(root, conf)

	t.Log(d.Config())
	t.Log(d.Nav())
}

func TestRegisterDriver(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()
}

func TestRegisterRegisteredDriver(t *testing.T) {
	defer func() { recover() }()

	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	NewDriverTest(&Foo{}, DriverConfig{})
	t.Error("should have panic")
}

func TestUnregisterDriver(t *testing.T) {
	d := NewDriverTest(&Foo{}, DriverConfig{})
	defer d.Close()

	UnregisterDriver(d)
}

func TestDriverByID(t *testing.T) {
	id := DriverToken("SuperDriver")
	d := NewDriverTest(&Foo{}, DriverConfig{ID: id})
	defer d.Close()

	if d2, _ := DriverByID(id); d2 != d {
		t.Errorf("d2 should be %#v: %#v", d, d2)
	}
}

func TestDriverByComponent(t *testing.T) {
	root := &Bar{
		Foo: &Foo{},
	}

	conf := DriverConfig{}
	d := NewDriverTest(root, conf)
	defer d.Close()

	if driver := DriverByComponent(root); driver != d {
		t.Errorf("driver should be %#v: %#v", d, driver)
	}
}
