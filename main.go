package main

import (
	"github.com/Zyko0/go-sdl3/bin/binimg"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

var (
	SCRW, SCRH = 1920, 1080
	RND        *sdl.Renderer
	WIN        *sdl.Window
	CNT        sdl.FPoint
	TEX        *sdl.Texture
	RUN        = true
)

func main() {

	defer binsdl.Load().Unload()
	defer binttf.Load().Unload()
	defer binimg.Load().Unload()
	defer sdl.Quit()

	ttf.Init()
	ttf.LoadLibrary(ttf.Path())
	defer ttf.CloseLibrary()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	WIN, RND, _ = sdl.CreateWindowAndRenderer("SDL3", SCRW, SCRH, 0)
	if err != nil {
		panic(err)
	}
	defer WIN.Destroy()
	defer RND.Destroy()

	RND.SetVSync(1)
	TEX, _ = RND.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, SCRW, SCRH)
	TEX.SetBlendMode(sdl.BLENDMODE_BLEND)
	RND.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	RND.SetRenderTarget(TEX)

	INITIAL()

	for RUN {

		B4()
		RND.SetDrawColor(COL2RGBA(BLACK()))
		RND.Clear()

		dLEV()

		AFTER()

	}
}
