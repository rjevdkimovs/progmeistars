package main

import (
	"os"
	"math/rand"
	"time"
	"github.com/nsf/termbox-go"
)
var counter int
var counterW chan int
var counterB chan int
func main() {
	counterW = make(chan int, 100)
    counterB = make(chan int, 100)
    counter = 0
    
		if err := termbox.Init(); err != nil {
			os.Exit(1)
		}
		defer termbox.Close()
		
		width, height := termbox.Size()
		
		x1, y1, vx1, vy1 := rand.Intn(width), rand.Intn(height), 1, -1
		x2, y2, vx2, vy2 := rand.Intn(width), rand.Intn(height), 2, -2
		
		for t:= time.Now(); time.Since(t) < 3*time.Second;	{
			termbox.SetCell(x1, y1, ' ', termbox.ColorDefault, termbox.ColorDefault)
			x1 += vx1
			y1 += vy1
			termbox.SetCell(x1, y1, ' ', termbox.ColorDefault, termbox.ColorWhite)
			termbox.SetCell(x2, y2, ' ', termbox.ColorDefault, termbox.ColorDefault)
			x2 += vx2
			y2 += vy2
			termbox.SetCell(x2, y2, ' ', termbox.ColorDefault, termbox.ColorBlue)			
			termbox.Flush()
			time.Sleep(30 * time.Millisecond)
			if x1<=0 || x1>= width-1 {
				vx1 = - vx1
			}
			if y1<=0 || y1>=height-1 {
				vy1 = - vy1
			}
			if x2<=0 || x2>= width-1 {
				vx2 = - vx2
			}
			if y2<=0 || y2>=height-1 {
				vy2 = - vy2
			}
		}
			
	}
		
			
