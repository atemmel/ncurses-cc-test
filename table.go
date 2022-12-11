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
	yOffset int
	w, h int
}

func makeTable(columns []Column) Table {
	return Table{
		Columns: columns,
		Cx: 0,
		Cy: 0,
		yOffset: 0,
		w: 0,
		h: 0,
	}
}

func (t *Table) Draw(x, y, w, h int) {
	t.w = w
	t.h = h
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

	y++
	for i := t.yOffset; i < len(c.Cells); i++ {
		if y >= t.h {
			break
		}
		cell := &c.Cells[i]
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

func (t *Table) Left() {
	t.Cx -= 1
	if t.Cx < 0 {
		t.Cx = 0
	}
	t.fitCursor()
}

func (t *Table) Up() {
	t.Cy -= 1
	if t.Cy < 0 {
		t.Cy = 0
	}
	if t.yOffset > t.Cy {
		t.yOffset--
	}
}

func (t *Table) Down() {
	maxY := t.currColLen()
	t.Cy += 1
	reachedBottom := t.Cy >= maxY
	if reachedBottom {
		t.Cy = maxY - 1
	}
	if t.Cy + 2 >= t.h {
		t.yOffset++
		if reachedBottom {
			t.yOffset--
		}
	}
}

func (t *Table) Right() {
	maxX := len(t.Columns)
	t.Cx += 1
	if t.Cx >= maxX {
		t.Cx = maxX - 1
	}
	t.fitCursor()
}

func (t *Table) Top() {
	t.Cy = 0
	t.yOffset = 0
}

func (t *Table) Bottom() {
	maxY := t.currColLen()
	t.Cy = maxY - 1
	t.yOffset = 0
	if t.Cy > t.h {
		t.yOffset = t.h - 2
	}
	if t.Cy < 0 {
		t.Cy = 0
		t.yOffset = 0
	}
}

func (t *Table) fitCursor() {
	maxY := t.currColLen()
	if t.Cy >= maxY {
		t.Cy = maxY - 1
	} 
}

func (t *Table) currColLen() int {
	return len(t.Columns[t.Cx].Cells)
}
