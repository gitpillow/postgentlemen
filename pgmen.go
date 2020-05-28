package main

import (
	"github.com/gitpillow/goji"
	"github.com/gitpillow/postgentlemen/db"
	"github.com/gitpillow/postgentlemen/pgmenui"
	_ "github.com/gitpillow/postgentlemen/resource"
	"github.com/jroimartin/gocui"
	"log"
)

const UseGoji = goji.UseGoji

func main() {
	db.CreateDB()
	db.Migrate()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	pgmenui.Load(g)
}
