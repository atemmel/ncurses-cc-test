package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func mvprint(x, y int, args... any) {
	str := fmt.Sprint(args...)
	mvprints(x, y, str)
}

func mvprintf(x, y int, msg string, args... any) {
	str := fmt.Sprintf(msg, args...)
	mvprints(x, y, str)
}

func mvprintscenter(x, y, w int, msg string) {
	x += w / 2 - len(msg) / 2
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
		x += runewidth.RuneWidth(c)
	}
}

func mvprints(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
		x += runewidth.RuneWidth(c)
	}
}
