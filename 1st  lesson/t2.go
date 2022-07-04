package main
import (
	"os"
	"github.com/nsf/termbox-go"
	"time"
)
	type (
		Runer struct {
			x, y int
			vx, vy int
			color termbox.Attribute
			delay time.Duration
	 }
 )
 func NewRunner (x, y int, vx, vy int, color termbox.Attribute, delay time.Duration) {
		var r Runer
		r.x, r.y = x, y
		r.vx, r.vy = vx, vy
		r.color = color
		r.delay = delay
		return r
		
	}
	func (r*Runer) Step() {
		width
	}
