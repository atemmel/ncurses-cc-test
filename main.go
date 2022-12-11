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
	//w, h := termbox.Size()
	w, h := 40, 6
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
				{
					Timestamp: time.Now(),
					Content: "papaya",
				},
				{
					Timestamp: time.Now(),
					Content: "papaya",
				},
				{
					Timestamp: time.Now(),
					Content: "papaya",
				},
				{
					Timestamp: time.Now(),
					Content: "papaya",
				},
			},
		},
		{
			Title: "Böcker",
			Cells: []Cell{
				{
					Timestamp: time.Now(),
					Content: "Illiaden",
				},
				{
					Timestamp: time.Now(),
					Content: "Odysséen",
				},
				{
					Timestamp: time.Now(),
					Content: "Håkan Bråkan och Roboten Rex",
				},
				{
					Timestamp: time.Now(),
					Content: "Horus Rising",
				},
				{
					Timestamp: time.Now(),
					Content: "Twilight",
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
					table.Left()
				case 'j':
					table.Down()
				case 'k':
					table.Up()
				case 'l':
					table.Right()
				case 'g':
					table.Top()
				case 'G':
					table.Bottom()
				default:
					switch event.Key {
						case termbox.KeyArrowLeft:
							table.Left()
						case termbox.KeyArrowDown:
							table.Down()
						case termbox.KeyArrowUp:
							table.Up()
						case termbox.KeyArrowRight:
							table.Right()
						case termbox.KeyHome:
							table.Top()
						case termbox.KeyEnd:
							table.Bottom()
					}
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
