package iu

import "fmt"

type Composer interface {
	Template() string
}

type composerMap map[string]interface{}

func (m composerMap) OnEvent(eventName string, arg string) string {
	return fmt.Sprintf("CallEventHandler(this.id, '%v', %v)", eventName, arg)
}
