package main

import (
	"time"

	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	//FPS TIMERS
	nowDelta, lastDelta uint64
	Delta               float32
	//FPS
	setFPS     = 60
	targetFPS  float64
	frameDelay = float64(1000 / targetFPS) // Delay in milliseconds
	frameStart time.Time
	frameTime  float64
	FPS        float64
	frameCount int
	startTime  = time.Now()
	Seconds    int
)

func UP() {
	//MOUSE
	_, x, y := sdl.GetMouseState()
	MS.X, MS.Y = float32(x), float32(y)
	INP()
	uPL()
	getDelta()
}

func getDelta() {
	tickT := sdl.Ticks()
	Delta = float32(tickT-lastDelta) * 0.001
	lastDelta = tickT
}
