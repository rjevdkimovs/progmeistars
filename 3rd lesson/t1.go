package main

import(
	"fmt"
	"github.com/nsf/termbox-go"
	
)


func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += 1
	}
}

func main() {
	termbox.Init()

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
	for j, bg = range colorRange {
		tbprint(j*8, 0, termbox.ColorDefault, bg, "       ")
		tbprint(j*8, i+2, termbox.ColorDefault, bg, "       ")
	}

	tbprint(15, i+4, termbox.ColorDefault, termbox.ColorDefault,
		"Press any key to close...")
	termbox.Flush()
	termbox.PollEvent()
	termbox.Close()
}
