package main

// TODO
// Frame rate independance
// Score
// Game over state
// 2 Player
// AI needs imperfections
// window resizing

import (
	"fmt"
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

type ball struct {
	position
	radius    float32
	xvelocity float32
	yvelocity float32
	color     color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x+x), int(ball.y+y), ball.color, pixels)
			}
		}
	}
}

func getCenter() position {
	return position{float32(winWidth) / 2, float32(winHeight) / 2}
}

func (ball *ball) update(leftship *ship, elapsedTime float32) {
	ball.x += ball.xvelocity * elapsedTime
	ball.y += ball.yvelocity * elapsedTime

	// handle collisions
	if ball.y-ball.radius < 0 || ball.y+ball.radius > float32(winHeight) {
		ball.yvelocity = -ball.yvelocity
	}

	if ball.x < 0 {
		ball.position = getCenter()
		state = start
	} else if ball.x > float32(winWidth) {
		leftship.score++
		ball.position = getCenter()
		state = start
	}

	if ball.x-ball.radius < leftship.x+leftship.width/2 {
		if ball.y > leftship.y-leftship.height/2 && ball.y < leftship.y+leftship.height/2 {
			ball.xvelocity = -ball.xvelocity
			ball.x = leftship.x + leftship.width/2.0 + ball.radius
		}
	}
}

type ship struct {
	position
	width  float32
	height float32
	speed  float32
	color  color
}

func lerp(a float32, b float32, pct float32) float32 {
	return a + pct*(b-a)
}

func (ship *ship) draw(pixelsize int, pixels []byte) {
	startX := int(ship.position.x) - (pixelsize*16)/2
	startY := int(ship.position.y) - (pixelsize*16)/2

	for i, v := range shipGraphic {
		if v == 1 {
			for y := startY; y < startY+pixelsize; y++ {
				for x := startX; x < startX+pixelsize; x++ {
					setPixel(x, y, ship.color, pixels)
				}
			}
		}
		startX += pixelsize
		if (i+1)%16 == 0 {
			startY += pixelsize
			startX -= pixelsize * 16
		}
	}
}

func (ship *ship) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_W] != 0 {
		ship.y -= ship.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_S] != 0 {
		ship.y += ship.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_A] != 0 {
		ship.x -= ship.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_D] != 0 {
		ship.x += ship.speed * elapsedTime
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

	ship := ship{position{50, 100}, 20, 100, 300, color{255, 255, 255}}
	ball := ball{position{300, 300}, 20, 400, 400, color{255, 255, 255}}

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
		ball.update(&ship, elapsedTime)

		clear(pixels)
		ship.draw(4, pixels)
		ball.draw(pixels)

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
