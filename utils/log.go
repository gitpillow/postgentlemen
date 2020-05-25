package utils

import (
	"github.com/gizak/termui/v3"
	"log"
)

var EventNames map[int]string = map[int]string{
	1: "KeyboardEvent",
	2: "MouseEvent",
	3: "ResizeEvent",
}

func PrintEvent(e termui.Event) {
	log.Printf("event type: %v, event id: %v, event payload: %v", EventNames[int(e.Type)], e.ID, e.Payload)
}
