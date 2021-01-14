package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}

type position struct {
	x, y float32
}

type ball struct {
	position
	radius    int
	xvelocity float32
	yvelocity float32
	color     color
}

func (ball *ball) draw(pixels []byte) {
	//YAGNI Ya Aint Gonna Need It - meaning: instead of spending days optimizing how to draw a circle, just make it work and optimize it when necessary

	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	ball.x += ball.xvelocity
	ball.y += ball.yvelocity

	// handle collisions
	if int(ball.y)-ball.radius < 0 || int(ball.y)+ball.radius > winHeight {
		ball.yvelocity = -ball.yvelocity
	}

	if ball.x < 0 || int(ball.x) > winWidth {
		ball.x = 300
		ball.y = 300
	}

	if int(ball.x) < int(leftPaddle.x)+leftPaddle.width/2 {
		if int(ball.y) > int(leftPaddle.y)-leftPaddle.height/2 && int(ball.y) < int(leftPaddle.y)+leftPaddle.height/2 {
			ball.xvelocity -= ball.xvelocity
		}
	}

	if int(ball.x) > int(rightPaddle.x)+rightPaddle.width/2 {
		if int(ball.y) > int(rightPaddle.y)-rightPaddle.height/2 && int(ball.y) < int(rightPaddle.y)+rightPaddle.height/2 {
			ball.xvelocity -= ball.xvelocity
		}
	}

}

type paddle struct {
	position
	width  int
	height int
	color  color
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.width/2
	startY := int(paddle.y) - paddle.height/2

	for y := 0; y < paddle.height; y++ {
		for x := 0; x < paddle.width; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}
}

func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y--
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y++
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

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {

	// Required for Mac
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
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

	player1 := paddle{position{50, 100}, 20, 100, color{255, 255, 255}}
	player2 := paddle{position{float32(winWidth) - 50, 700}, 20, 100, color{255, 255, 255}}
	ball := ball{position{300, 300}, 20, 2, 2, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	// For Mac
	// OSX retuires that you consume events for windows to open and work properly
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		clear(pixels)

		player1.update(keyState)
		player2.aiUpdate(&ball)
		ball.update(&player1, &player2)

		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)
	}

}
