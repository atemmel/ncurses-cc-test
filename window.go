package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

const (
	BorderLeftTop  = '┌'
	BorderTop      = '─'
	BorderRightTop = '┐'
	BorderRight    = '│'
	BorderLeftBot  = '└'
	BorderBot      = BorderTop
	BorderRightBot = '┘'
)

type Cell struct {
	Timestamp time.Time
	Content string
}

type Column struct {
	Title string
	Cells []Cell
}

type Table struct {
	Columns []Column
	Cx, Cy int
}

func makeTable(columns []Column) Table {
	return Table{
		Columns: columns,
		Cx: 0,
		Cy: 0,
	}
}

func (t *Table) Draw(x, y, w, h int) {
	n := len(t.Columns)
	cw := w / n
	for i, c := range t.Columns {
		focused := i == t.Cx
		t.DrawColumn(&c, x, y, cw, h, focused)
		x += cw
	}
}

func (t *Table) DrawColumn(c *Column, x, y, w, h int, focused bool) {
	mvprintscenter(x, y, w, c.Title)

	bg := termbox.ColorBlack
	fg := termbox.ColorWhite

	if focused {
		bg, fg = fg, bg
	} 

	line(x, y, w, bg, fg)

	y += 1
	for i, cell := range c.Cells {
		mvprints(x, y, cell.Timestamp.Format(time.Stamp))
		bg = termbox.ColorBlack
		fg = termbox.ColorWhite

		if focused && i  == t.Cy {
			bg, fg = fg, bg
		}

		line(x, y, w, bg, fg)
		y++
	}
}

func line(x, y, w int, bg, fg termbox.Attribute) {
	for i := x; i < w; i++ {
		termbox.SetBg(i, y, bg)
		termbox.SetFg(i, y, fg)
	}
}
