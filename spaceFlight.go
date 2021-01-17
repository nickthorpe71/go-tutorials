package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 800

// -- ENUM
type gameState int

const (
	start gameState = iota
	play
)

var state = start

var shipGraphic = []byte{
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1,
	1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1,
	1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1,
	1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1,
}

type color struct {
	r, g, b byte
}

type position struct {
	x, y float32
}

type debris struct {
	position
	radius    float32
	xvelocity float32
	yvelocity float32
	color     color
}

func (debris *debris) draw(pixels []byte) {
	for y := -debris.radius; y < debris.radius; y++ {
		for x := -debris.radius; x < debris.radius; x++ {
			if x*x+y*y < debris.radius*debris.radius {
				setPixel(int(debris.x+x), int(debris.y+y), debris.color, pixels)
			}
		}
	}
}

func getCenter() position {
	return position{float32(winWidth) / 2, float32(winHeight) / 2}
}

func (debris *debris) update(ship *ship, elapsedTime float32) {
	debris.x += debris.xvelocity * elapsedTime
	debris.y += debris.yvelocity * elapsedTime

	// handle collisions
	if debris.y-debris.radius > float32(winHeight) {
		debris.y = float32(-20 - rand.Intn(40))
		debris.x = 20 + float32(rand.Intn(winWidth-40))
		debris.radius = 20 + float32(rand.Intn(20))
		debris.yvelocity = 50 + float32(rand.Intn(400))
	}

	if debris.x-debris.radius < ship.x+ship.width/3 && debris.x+debris.radius > ship.x-ship.width/3 {
		if debris.y+debris.radius > ship.y-ship.height/3 && debris.y-debris.radius < ship.y+ship.height/3 {
			ship.alive = false
		}
	}
	//  -----
}

type ship struct {
	position
	width  float32
	height float32
	speed  float32
	alive  bool
	color  color
}

func lerp(a float32, b float32, pct float32) float32 {
	return a + pct*(b-a)
}

func (ship *ship) draw(pixels []byte) {
	if ship.alive {
		startX := ship.position.x - ship.width/2
		startY := ship.position.y - ship.height/2

		for i, v := range shipGraphic {
			if v == 1 {
				for y := startY; y < startY+4; y++ {
					for x := startX; x < startX+4; x++ {
						setPixel(int(x), int(y), ship.color, pixels)
					}
				}
			}
			startX += 4
			if (i+1)%16 == 0 {
				startY += 4
				startX -= 4 * 16
			}
		}
	}
}

func (ship *ship) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_W] != 0 {
		if ship.y-ship.height/2 > 0 {
			ship.y -= ship.speed * elapsedTime
		}
	}
	if keyState[sdl.SCANCODE_S] != 0 {
		if ship.y+ship.height/2 < float32(winHeight) {
			ship.y += ship.speed * elapsedTime
		}
	}
	if keyState[sdl.SCANCODE_A] != 0 {
		if ship.x-ship.width/2 > 0 {
			ship.x -= ship.speed * elapsedTime
		}
	}
	if keyState[sdl.SCANCODE_D] != 0 {
		if ship.x+ship.width/2 < float32(winHeight) {
			ship.x += ship.speed * elapsedTime
		}
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}

}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Saving Ana", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	ship := ship{position{400, 700}, 64, 64, 300, true, color{255, 255, 255}}
	debris1 := debris{position{20 + float32(rand.Intn(winWidth-40)), 0}, 20 + float32(rand.Intn(20)), 0, 100 + float32(rand.Intn(400)), color{255, 222, 222}}
	debris2 := debris{position{20 + float32(rand.Intn(winWidth-40)), -20 - float32(rand.Intn(20))}, 20 + float32(rand.Intn(20)), 0, 50 + float32(rand.Intn(300)), color{255, 222, 222}}
	debris3 := debris{position{20 + float32(rand.Intn(winWidth-40)), -30 - float32(rand.Intn(20))}, 20 + float32(rand.Intn(20)), 0, 50 + float32(rand.Intn(300)), color{255, 222, 222}}
	debris4 := debris{position{20 + float32(rand.Intn(winWidth-40)), -50 - float32(rand.Intn(20))}, 20 + float32(rand.Intn(20)), 0, 50 + float32(rand.Intn(300)), color{255, 222, 222}}
	debris5 := debris{position{20 + float32(rand.Intn(winWidth-40)), -80 - float32(rand.Intn(20))}, 20 + float32(rand.Intn(20)), 0, 50 + float32(rand.Intn(300)), color{255, 222, 222}}

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32

	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		ship.update(keyState, elapsedTime)
		debris1.update(&ship, elapsedTime)
		debris2.update(&ship, elapsedTime)
		debris3.update(&ship, elapsedTime)
		debris4.update(&ship, elapsedTime)
		debris5.update(&ship, elapsedTime)

		clear(pixels)
		ship.draw(pixels)
		debris1.draw(pixels)
		debris2.draw(pixels)
		debris3.draw(pixels)
		debris4.draw(pixels)
		debris5.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime/1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}

	}

}
