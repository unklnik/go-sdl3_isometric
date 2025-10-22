package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	dBUG, dbgBOOL bool
	dbgNUM        int
	dbgRecW       = float32(300)
)

func DEBUG() {

	dRecFillAlpha(sdl.FRect{0, 0, dbgRecW, float32(SCRH)}, RED(), 70)
	var x, y float32 = 4, 4
	dTxtXY("dbgBOOL "+fmt.Sprint(dbgBOOL)+" dbgNUM "+fmt.Sprint(dbgNUM), x, y, 1)
	y += FONT1.hS
	dTxtXY("MS.X "+fmt.Sprint(MS.X)+" MS.Y "+fmt.Sprint(MS.Y), x, y, 1)
	y += FONT1.hS
	dTxtXY("pl.cnt.X "+fmt.Sprint(pl.cnt.X)+" pl.cnt.Y "+fmt.Sprint(pl.cnt.Y), x, y, 1)
	y += FONT1.hS
	dTxtXY("len(LEV) "+fmt.Sprint(len(LEV))+" frameCount "+fmt.Sprint(frameCount), x, y, 1)
	y += FONT1.hS
	dTxtXY("pl.zi "+fmt.Sprint(pl.zi)+" frameCount "+fmt.Sprint(frameCount), x, y, 1)
	y += FONT1.hS

}
