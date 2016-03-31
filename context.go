package iu

var Path func(...string) string

type Context interface {
	Name() string

	CurrentPage() *Page

	InjectComponent(component *Component)

	Navigate(page *Page)
}
