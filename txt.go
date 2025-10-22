package main

import (
	"strings"

	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

var (
	FONT1 FONT

	txtS, txtM, txtL, txtXL float32 = 16, 18, 24, 48

	standardCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:;<=>?!#$%&'()*+,-./@[]^_`{|}~'\"' "
)

type FONT struct {
	cS, cM, cL, cXL []CHAR
	hS, hM, hL, hXL float32
	f               ttf.Font
}
type CHAR struct {
	t   string
	r   sdl.FRect
	tex *sdl.Texture
}

// MARK: DRAW
func dTxtXY(t string, x, y float32, siz1234 int) {
	t2 := strings.Split(t, "")
	for i := range t2 {
		for j := range FONT1.cS {
			if t2[i] == FONT1.cS[j].t {
				switch siz1234 {
				case 1:
					rTex(FONT1.cS[j].r, sdl.FRect{x, y, FONT1.cS[j].r.W, FONT1.cS[j].r.H}, FONT1.cS[j].tex)
					x += FONT1.cS[j].r.W
				case 2:
					rTex(FONT1.cM[j].r, sdl.FRect{x, y, FONT1.cM[j].r.W, FONT1.cM[j].r.H}, FONT1.cM[j].tex)
					x += FONT1.cM[j].r.W
				case 3:
					rTex(FONT1.cL[j].r, sdl.FRect{x, y, FONT1.cL[j].r.W, FONT1.cL[j].r.H}, FONT1.cL[j].tex)
					x += FONT1.cL[j].r.W
				case 4:
					rTex(FONT1.cXL[j].r, sdl.FRect{x, y, FONT1.cXL[j].r.W, FONT1.cXL[j].r.H}, FONT1.cXL[j].tex)
					x += FONT1.cXL[j].r.W
				}
			}
		}
	}
}
func dTxtXYcolor(t string, x, y float32, siz1234 int, c sdl.Color) {
	t2 := strings.Split(t, "")
	for i := range t2 {
		for j := range FONT1.cS {
			if t2[i] == FONT1.cS[j].t {
				switch siz1234 {
				case 1:
					FONT1.cS[j].tex = texColor(FONT1.cS[j].tex, c)
					rTex(FONT1.cS[j].r, sdl.FRect{x, y, FONT1.cS[j].r.W, FONT1.cS[j].r.H}, FONT1.cS[j].tex)
					x += FONT1.cS[j].r.W
					FONT1.cS[j].tex = texRevertColor(FONT1.cS[j].tex)
				case 2:
					FONT1.cM[j].tex = texColor(FONT1.cM[j].tex, c)
					rTex(FONT1.cM[j].r, sdl.FRect{x, y, FONT1.cM[j].r.W, FONT1.cM[j].r.H}, FONT1.cM[j].tex)
					x += FONT1.cM[j].r.W
					FONT1.cM[j].tex = texRevertColor(FONT1.cM[j].tex)
				case 3:
					FONT1.cL[j].tex = texColor(FONT1.cL[j].tex, c)
					rTex(FONT1.cL[j].r, sdl.FRect{x, y, FONT1.cL[j].r.W, FONT1.cL[j].r.H}, FONT1.cL[j].tex)
					x += FONT1.cL[j].r.W
					FONT1.cL[j].tex = texRevertColor(FONT1.cL[j].tex)
				case 4:
					FONT1.cXL[j].tex = texColor(FONT1.cXL[j].tex, c)
					rTex(FONT1.cXL[j].r, sdl.FRect{x, y, FONT1.cXL[j].r.W, FONT1.cXL[j].r.H}, FONT1.cXL[j].tex)
					x += FONT1.cXL[j].r.W
					FONT1.cXL[j].tex = texRevertColor(FONT1.cXL[j].tex)
				}
			}
		}
	}
}

// MARK: MAKE
func mFonts() {
	t := strings.Split(standardCharacters, "")
	for i := range t {
		FONT1.cS = append(FONT1.cS, mChar(t[i], "Rubik-Medium.ttf", 1))
		FONT1.cM = append(FONT1.cM, mChar(t[i], "Rubik-Medium.ttf", 2))
		FONT1.cL = append(FONT1.cL, mChar(t[i], "Rubik-Medium.ttf", 3))
		FONT1.cXL = append(FONT1.cXL, mChar(t[i], "Rubik-Medium.ttf", 4))
	}
	FONT1.hS = FONT1.cS[0].r.H
	FONT1.hM = FONT1.cM[0].r.H
	FONT1.hL = FONT1.cL[0].r.H
	FONT1.hXL = FONT1.cXL[0].r.H
}
func mChar(t, fontPath string, siz1234 int) CHAR {
	c := CHAR{}
	var surf *sdl.Surface
	var f *ttf.Font
	switch siz1234 {
	case 1:
		f, _ = ttf.OpenFont(fontPath, txtS)
	case 2:
		f, _ = ttf.OpenFont(fontPath, txtM)
	case 3:
		f, _ = ttf.OpenFont(fontPath, txtL)
	case 4:
		f, _ = ttf.OpenFont(fontPath, txtXL)
	}
	surf, _ = f.RenderTextBlended(t, WHITE())
	c.tex, _ = RND.CreateTextureFromSurface(surf)
	c.tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	c.r = sdl.FRect{0, 0, float32(c.tex.W), float32(c.tex.H)}
	c.t = t
	f.Close()
	surf.Destroy()
	return c
}
