package pgmenui

import "github.com/gizak/termui/v3/widgets"

func LeftSideBar() *widgets.List {
	l := widgets.NewList()
	l.Title = "resources"
	return l
}
