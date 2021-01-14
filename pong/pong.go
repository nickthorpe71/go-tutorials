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
				setPixel(int(ball.x)+x, int(ball.y)+y, color{255, 255, 255}, pixels)
			}
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

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(startX+x, startY+y, color{255, 255, 255}, pixels)
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

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
		}
	}

	tex.Update(nil, pixels, winWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	// For Mac
	// OSX retuires that you consume events for windows to open and work properly
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		sdl.Delay(16)
	}

}
