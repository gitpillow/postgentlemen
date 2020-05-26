package main

import (
	"github.com/gitpillow/goji"
	"github.com/gitpillow/postgentlemen/db"
	_ "github.com/gitpillow/postgentlemen/resource"
	"github.com/gitpillow/postgentlemen/utils"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

const UseGoji = goji.UseGoji

func main() {
	db.CreateDB()
	db.Migrate()

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "test"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			utils.PrintEvent(e)
			break
		}
	}
}
