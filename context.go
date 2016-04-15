package iu

var Path func(...string) string

type Context interface {
	CurrentPage() *Page

	Navigate(page *Page)

	InjectComponent(component *Component)

	ShowContextMenu(menus []Menu, compoID string)
}
