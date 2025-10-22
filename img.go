package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
)

type ANIM struct {
	ims         []IM
	dFrame, fps int
	timer       time.Time
}

type IM struct {
	r   sdl.FRect
	tex *sdl.Texture
}

// MARK:DRAW ANIM
func dAnimFrame(anm ANIM) ANIM {
	if anm.timer.IsZero() {
		anm.timer = time.Now()
	}
	if time.Since(anm.timer) >= time.Second/time.Duration(anm.fps) {
		anm.dFrame++
		if anm.dFrame >= len(anm.ims) {
			anm.dFrame = 0
		}
		anm.timer = time.Now()
	}
	return anm
}
func dAnimFramePart(anm ANIM, numframes, fps int) ANIM {
	if anm.timer.IsZero() {
		anm.timer = time.Now()
	}
	if time.Since(anm.timer) >= time.Second/time.Duration(fps) {
		anm.dFrame++
		if anm.dFrame >= numframes {
			anm.dFrame = 0
		}
		anm.timer = time.Now()
	}
	return anm
}

func dAnimRecLoopPart(anm ANIM, r sdl.FRect, flip bool, numframes, fps int) ANIM {
	if flip {
		dIMRecFlip(anm.ims[anm.dFrame], r)
	} else {
		dIMRec(anm.ims[anm.dFrame], r)
	}
	anm = dAnimFramePart(anm, numframes, fps)
	return anm
}
func dAnimRecLoopColorPart(anm ANIM, r sdl.FRect, c sdl.Color, flip bool, numframes, fps int) ANIM {
	if flip {
		dIMRecFlipColor(anm.ims[anm.dFrame], r, c)
	} else {
		dIMRecColor(anm.ims[anm.dFrame], r, c)
	}
	anm = dAnimFramePart(anm, numframes, fps)
	return anm
}
func dAnimRecLoop(anm ANIM, r sdl.FRect, flip bool) ANIM {
	if flip {
		dIMRecFlip(anm.ims[anm.dFrame], r)
	} else {
		dIMRec(anm.ims[anm.dFrame], r)
	}
	anm = dAnimFrame(anm)
	return anm
}
func dAnimRecLoopColor(anm ANIM, r sdl.FRect, c sdl.Color, flip bool) ANIM {
	if flip {
		dIMRecFlipColor(anm.ims[anm.dFrame], r, c)
	} else {
		dIMRecColor(anm.ims[anm.dFrame], r, c)
	}

	anm = dAnimFrame(anm)
	return anm
}

// MARK: DRAW
func dIMSheetXY(im []IM, x, y, siz float32) {
	if dBUG {
		if x < dbgRecW {
			x = dbgRecW + 4
		}
	}
	ox := x
	for i := range im {
		dIMRec(im[i], sdl.FRect{x, y, im[i].r.W * siz, im[i].r.H * siz})
		if dBUG {
			dTxtXY(fmt.Sprint(i), x, y+im[i].r.H*siz, 1)
		}
		x += im[i].r.W * siz
		if x+im[i].r.W*siz > float32(SCRW) {
			x = ox
			y += im[i].r.H * siz
			if dBUG {
				y += FONT1.hS
			}
		}
	}
}
func dIMXY(im IM, x, y float32) {
	RND.RenderTexture(im.tex, &im.r, &sdl.FRect{x, y, im.r.W, im.r.H})
}
func dIMRec(im IM, r sdl.FRect) {
	RND.RenderTexture(im.tex, &im.r, &r)
}
func dIMRecFlip(im IM, r sdl.FRect) {
	RND.RenderTextureRotated(im.tex, &im.r, &r, 0, qRecCNT(r), sdl.FLIP_HORIZONTAL)
}
func dIMRecColor(im IM, r sdl.FRect, c sdl.Color) {
	im.tex = texColor(im.tex, c)
	RND.RenderTexture(im.tex, &im.r, &r)
	im.tex = texRevertColor(im.tex)
}
func dIMRecFlipColor(im IM, r sdl.FRect, c sdl.Color) {
	im.tex = texColor(im.tex, c)
	RND.RenderTextureRotated(im.tex, &im.r, &r, 0, qRecCNT(r), sdl.FLIP_HORIZONTAL)
	im.tex = texRevertColor(im.tex)
}

// MARK: MAKE
func mAnimIMSheet(ims []IM, fps int) ANIM {
	anm := ANIM{}
	anm.ims = ims
	anm.fps = fps
	return anm
}
func mIMSheetXY1RowNum(path string, x, y, w, h float32, num int) []IM {
	var ims []IM
	surf, _ := img.Load(path)
	tex, _ := RND.CreateTextureFromSurface(surf)
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	surf.Destroy()
	for num > 0 {
		im := IM{}
		im.tex = tex
		im.r = sdl.FRect{float32(x), float32(y), float32(w), float32(h)}
		ims = append(ims, im)
		x += w
		num--
	}
	return ims
}
func mIMSheetXYMultiRowTexSize(path string, x, y, w, h int32) []IM {
	var ims []IM
	surf, _ := img.Load(path)
	tex, _ := RND.CreateTextureFromSurface(surf)
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	surf.Destroy()
	numW := tex.W / w
	numH := tex.H / h
	a := numW * numH
	c := 0
	ox := x
	for a > 0 {
		im := IM{}
		im.tex = tex
		im.r = sdl.FRect{float32(x), float32(y), float32(w), float32(h)}
		ims = append(ims, im)
		x += w
		c++
		a--
		if c == int(numW) {
			c = 0
			x = ox
			y += h
		}
	}

	return ims
}
func mIMSheetPathFiles(path string) []IM {
	var im []IM
	s := fPNG(path)
	for i := range s {
		im = append(im, mIMPath(s[i]))
	}
	return im
}
func mIMPath(path string) IM {
	var im IM
	surf, _ := img.Load(path)
	im.tex, _ = RND.CreateTextureFromSurface(surf)
	im.tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	im.r = sdl.FRect{0, 0, float32(im.tex.W), float32(im.tex.H)}
	surf.Destroy()
	return im
}

// MARK: UTILS
func fPNG(path string) []string {
	baseDir, _ := os.Getwd()
	searchDir := filepath.Join(baseDir, path)

	var files []string
	filepath.WalkDir(searchDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// ignore errors silently
			return nil
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".png") {
			rel, err := filepath.Rel(baseDir, path)
			if err == nil {
				files = append(files, rel)
			}
		}
		return nil
	})

	return files
}
