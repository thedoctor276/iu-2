package iu

import (
	"fmt"

	"github.com/maxence-charriere/iu-log"
)

var compoM = newCompoManager()

type compoManager struct {
	components map[View]*Component
	views      map[string]View
}

func (manager *compoManager) Register(v View) (err error) {
	var c *Component
	var ok bool

	if c, ok = manager.components[v]; ok {
		err = fmt.Errorf("%#v is already registered in compoManager %p", v, manager)
		return
	}

	c = NewComponent(v)
	manager.components[v] = c
	manager.views[c.ID()] = v
	return
}

func (manager *compoManager) Unregister(v View) (err error) {
	var c *Component
	var ok bool

	if c, ok = manager.components[v]; !ok {
		err = fmt.Errorf("%#v is not found in compoManager %p", v, manager)
		return
	}

	delete(manager.components, v)
	delete(manager.views, c.ID())
	return
}

func (manager *compoManager) Component(v View) (c *Component) {
	var ok bool

	if c, ok = manager.components[v]; !ok {
		iulog.Panicf("no component resgistered for %#v", v)
	}

	return
}

func (manager *compoManager) View(compoID string) (v View) {
	var ok bool

	if v, ok = manager.views[compoID]; !ok {
		iulog.Panicf("no view registered for %#v", compoID)
	}

	return
}

func newCompoManager() *compoManager {
	return &compoManager{
		components: map[View]*Component{},
		views:      map[string]View{},
	}
}

func RegisterView(v View) error {
	return compoM.Register(v)
}

func UnregisterView(v View) error {
	return compoM.Unregister(v)
}

func ViewFromComponentID(ID string) View {
	return compoM.View(ID)
}
