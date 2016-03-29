package iu

import "path/filepath"

var Path func(...string) string = filepath.Join

type Context interface {
	Name() string

	CurrentPage() *Page

	InjectComponent(component *Component)

	Navigate(page *Page)
}
