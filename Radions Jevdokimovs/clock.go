package main
import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"math"
	"time"
)
var winTitle string = "TEST"
var winWidth, winHeight int32 = 800, 600
func drawcircle(renderer *sdl.Renderer, cx, cy, r int32) {
	var px, py, ax, ay, bx, by int32
	ax = cx
	ay = cy - r
	px = ax
	py = ay
	var i float64
	renderer.SetDrawColor(255, 255, 255, 255)
	for j := 0; j < 365; j++ {
		i = float64(2*math.Pi * float64(j) / 360)
		//bx = int32(float64(ax) * math.Cos(i) - float64(cx) * math.Cos(i) - float64(ay) * math.Sin(i) + float64(cy) * math.Sin(i) + float64(cx))
		//by = int32(float64(ax) * math.Sin(i) - float64(cx) * math.Sin(i) + float64(ay) * math.Cos(i) - float64(cy) * math.Cos(i) + float64(cy))
		bx = int32(float64(ax) * math.Cos(i) - float64(ay) * math.Sin(i) + (float64(cx) * (1 - math.Cos(i)) + float64(cy) * math.Sin(i)))
		by = int32(float64(ax) * math.Sin(i) + float64(ay) * math.Cos(i) + (float64(-cx) * math.Sin(i) + float64(cy) * (1 - math.Cos(i))))
		renderer.DrawLine(bx, by, px, py)
		px = bx
		py = by
	}
}

func drawtime(renderer *sdl.Renderer, cx, cy, r int32) {
	renderer.SetDrawColor(255, 255, 255, 255)
	var ssec, smin, shour, ax, ay, bx, by int32
	var sec, min, hour int
	ax = cx
	ssec = r / 1
	smin = r / 2
	shour = r / 3
	hour, min, sec = time.Now().Clock()
	var i float64
	ay = cy - ssec
	i = float64(2*math.Pi * float64(sec) / 60)
	bx = int32(float64(ax) * math.Cos(i) - float64(ay) * math.Sin(i) + (float64(cx) * (1 - math.Cos(i)) + float64(cy) * math.Sin(i)))
	by = int32(float64(ax) * math.Sin(i) + float64(ay) * math.Cos(i) + (float64(-cx) * math.Sin(i) + float64(cy) * (1 - math.Cos(i))))
	renderer.DrawLine(bx, by, cx, cy)
	ay = cy - smin
	i = float64(2*math.Pi * float64(min) / 60)
	bx = int32(float64(ax) * math.Cos(i) - float64(ay) * math.Sin(i) + (float64(cx) * (1 - math.Cos(i)) + float64(cy) * math.Sin(i)))
	by = int32(float64(ax) * math.Sin(i) + float64(ay) * math.Cos(i) + (float64(-cx) * math.Sin(i) + float64(cy) * (1 - math.Cos(i))))
	renderer.DrawLine(bx, by, cx, cy)
	ay = cy - shour
	i = float64(2*math.Pi * (float64(hour) / 12 - 1))
	bx = int32(float64(ax) * math.Cos(i) - float64(ay) * math.Sin(i) + (float64(cx) * (1 - math.Cos(i)) + float64(cy) * math.Sin(i)))
	by = int32(float64(ax) * math.Sin(i) + float64(ay) * math.Cos(i) + (float64(-cx) * math.Sin(i) + float64(cy) * (1 - math.Cos(i))))
	renderer.DrawLine(bx, by, cx, cy)
}

func drawotmetki(renderer *sdl.Renderer, cx, cy, r int32) {
	var dx, dy, ax, ay, bx, by int32
	renderer.SetDrawColor(255, 255, 255, 255)
	ax = cx
	ay = cy - r
	var d int32
	d = r/4
	var i float64
	for j := 0; j < 12; j++ {
		ay = cy - r
		i = float64(2*math.Pi * float64(j) / 12)
		bx = int32(float64(ax) * math.Cos(i) - float64(ay) * math.Sin(i) + (float64(cx) * (1 - math.Cos(i)) + float64(cy) * math.Sin(i)))
		by = int32(float64(ax) * math.Sin(i) + float64(ay) * math.Cos(i) + (float64(-cx) * math.Sin(i) + float64(cy) * (1 - math.Cos(i))))
		
		ay = cy - r + d
		dx = int32(float64(ax) * math.Cos(i) - float64(ay) * math.Sin(i) + (float64(cx) * (1 - math.Cos(i)) + float64(cy) * math.Sin(i)))
		dy = int32(float64(ax) * math.Sin(i) + float64(ay) * math.Cos(i) + (float64(-cx) * math.Sin(i) + float64(cy) * (1 - math.Cos(i))))
		renderer.DrawLine(bx, by, dx, dy)
	}
}

func main() {	
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	defer sdl.Quit()
	var window *sdl.Window
	var renderer *sdl.Renderer
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return
	}
	defer renderer.Destroy()
	var ccx, ccy, r int32
	ccx = 400
	ccy = 300
	r = 200
	for {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		drawcircle(renderer, ccx, ccy, r)
		drawotmetki(renderer, ccx, ccy, r)
		drawtime(renderer, ccx, ccy, r)
		renderer.Present()
		sdl.PollEvent()
	}
}

