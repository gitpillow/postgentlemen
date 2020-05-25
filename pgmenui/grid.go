package pgmenui

import (
	ui "github.com/gizak/termui/v3"
)

func Grid() {
	grid := ui.NewGrid()
	w, h := ui.TerminalDimensions()
	grid.SetRect(0, 0, w, h)
}
