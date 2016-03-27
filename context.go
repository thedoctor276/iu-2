package iu

type Context interface {
	Name() string

	CurrentPage() *Page

	InjectComponent(component *Component)

	Navigate(page *Page)
}
