package iu

// package main

// type Component struct {
//     id  int
//     val Element
// }

// func NewComponent(val Element) *Component {
//     return &Compnent{
//         id:  generateID(),
//         val: val,
//     }
// }

// func (self *Component) Render() string {
//     m := map[string]interface{}{}

//     // ...

//     m["ID"] = self.id

//     m["OnEvent"] = func(eventType string, eventArg interface{}) {
//         v := reflect.ValueOf(self.val)
//         f := v.MethodByName(eventType)
//         reflect.Call(f, eventArg)
//     }

//     return self.val.Template().Render(m)
// }

// type Element interface {
//     Template() Template
// }

// type HelloWorld struct {
//     Hello string
//     World string
// }

// func (self *HelloWorld) Template() string {
//     return ""
// }

// func (self *HelloWorld) OnInputChanged(val string) {

// }

// func main() {
//     c := NewComponent(&HelloWorld{
//         Hello: "hello",
//         World: "world",
//     })
// }
