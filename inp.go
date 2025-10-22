package main

import (
	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	MS sdl.FPoint
)

func INP() {
	sdl.PumpEvents()
	var event sdl.Event
	for sdl.PollEvent(&event) {
		switch event.Type {
		case sdl.EVENT_KEY_DOWN:
			k := event.KeyboardEvent().Scancode
			switch k {

			case sdl.SCANCODE_ESCAPE:
				EXIT()
			case sdl.SCANCODE_F1:
				dBUG = !dBUG
			case sdl.SCANCODE_F2:
				RESTART()
			}
		case sdl.EVENT_QUIT:
			EXIT()
		}

	}
	movePL()

}
