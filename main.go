package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func main() {
	defer termbox.Close()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	printCenter("Hej hallå ncurses")
	termbox.Flush()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Ch == 'q' {
			break
		} 
	}
}

func printCenter(msg string) {
	w, h := termbox.Size()
	y := h / 2
	x := w / 2 - len(msg) / 2
	tbprint(x, y, msg)
}

func tbprint(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorBlack)
		x += runewidth.RuneWidth(c)
	}
}
