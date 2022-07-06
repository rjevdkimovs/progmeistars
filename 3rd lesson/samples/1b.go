package main

import (
	"fmt"
	"time"
	"github.com/nsf/termbox-go"
	"sync"
)
/*
type Job struct {
	x	int
	y	int
	msg uint64
}
*/
func tbprint(x, y int, bg termbox.Attribute) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, bg)
		x += 1
	}
}
func main() {
	termbox.Init()
	var jb(make chan int)
	var bg(make chan,2)
	var i, j int
	var fg, bg termbox.Attribute
	var colorRange []termbox.Attribute = []termbox.Attribute{
		termbox.ColorLightGreen,
	}

	var row, col int
	var text string
	for i, fg = range colorRange {
		for j, bg = range colorRange {
			row = i + 1
			col = j * 8
			text = fmt.Sprintf(" %02d/%02d ", fg, bg)
			tbprint(col, row+0, fg, bg, text)
		
		}
	}


		
		
		
}
