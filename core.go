package main

import (
	"time"

	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	UN = float32(32)
)

func RESTART() {
	mLEV()
	mPL()
}

func INITIAL() {

	//dBUG = true

	CNT = sdl.FPoint{float32(SCRW / 2), float32(SCRH / 2)}
	mFonts()

	mIMG()
	mLEV()
	mPL()

	//FPS
	targetFPS = float64(setFPS)
	frameDelay = 1000 / targetFPS // Delay in milliseconds
}

func B4() {
	RND.SetRenderTarget(TEX)
	frameStart = time.Now()
}

func AFTER() {
	if dBUG {
		DEBUG()
	}
	RND.SetRenderTarget(nil)
	RND.RenderTexture(TEX, nil, &sdl.FRect{0, 0, float32(SCRW), float32(SCRH)})
	RND.Present()
	UP()

	//FPS
	frameCount++

	frameTime = float64(time.Since(frameStart).Milliseconds())
	if frameTime < frameDelay {
		sdl.Delay(uint32(frameDelay - frameTime)) // Wait to achieve target FPS
	}
	if time.Since(startTime).Seconds() >= 1 {
		FPS = float64(frameCount) / time.Since(startTime).Seconds()
		frameCount = 0
		startTime = time.Now()
	}

	if frameCount == 0 {
		Seconds++
	}

}

func EXIT() {
	TEX.Destroy()
	RND.Destroy()
	WIN.Destroy()
	RUN = false
}
