package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func redraw(table *Table) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	drawHeader()
	mvprintf(5, 5, "kör! %s", "kobra")
	w, h := termbox.Size()
	table.Draw(0, 1, w, h)
	termbox.Flush()
}

func main() {
	defer termbox.Close()

	table := makeTable([]Column{
		{
			Title: "Stor-hasse",
			Cells: []Cell{
				{
					Timestamp: time.Now(),
					Content: "tjaba",
				},
				{
					Timestamp: time.Now(),
					Content: "tjena",
				},
				{
					Timestamp: time.Now(),
					Content: "hallå",
				},
			},
		},
		{
			Title: "Bra frukter",
			Cells: []Cell{
				{
					Timestamp: time.Now(),
					Content: "äpple",
				},
				{
					Timestamp: time.Now(),
					Content: "appelsin",
				},
				{
					Timestamp: time.Now(),
					Content: "päron",
				},
				{
					Timestamp: time.Now(),
					Content: "mango",
				},
				{
					Timestamp: time.Now(),
					Content: "papaya",
				},
			},
		},
	})

	LOOP: for {
		redraw(&table)
		event := termbox.PollEvent()
		switch event.Type {
			case termbox.EventKey: {
				switch event.Ch {
				case 'q':
					break LOOP
				case 'h':
					table.Cx -= 1
				case 'j':
					table.Cy += 1
				case 'k':
					table.Cy -= 1
				case 'l':
					table.Cx += 1
				}

			}

		}
	}
}

func drawHeader() {
	w, _ := termbox.Size()
	const msg = "monitor"
	y := 0

	x := w / 2 - len(msg) / 2
	mvprints(x, y, msg)

	for x := 0; x < w; x++ {
		termbox.SetBg(x, y, termbox.ColorBlue)
		termbox.SetFg(x, y, termbox.ColorWhite)
	}
}
