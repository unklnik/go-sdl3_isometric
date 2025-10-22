package main

import (
	"math/rand/v2"
	"sort"

	"github.com/Zyko0/go-sdl3/sdl"
)

// MARK: REC
func qRecCNT(r sdl.FRect) *sdl.FPoint {
	return &sdl.FPoint{r.X + r.W/2, r.Y + r.H/2}
}

// MARK: TEX
func texRevertColor(t *sdl.Texture) *sdl.Texture {
	t.SetColorMod(255, 255, 255)
	return t
}
func texColor(t *sdl.Texture, c sdl.Color) *sdl.Texture {
	t.SetColorMod(c.R, c.G, c.B)
	return t
}

// MARK: IM
func remIMfromSheet(im []IM, pos []int) []IM {
	if len(pos) == 0 {
		return im
	}
	sort.Ints(pos)
	write := 0
	skip := 0
	for read := 0; read < len(im); read++ {
		if skip < len(pos) && read == pos[skip] {
			skip++
			continue
		}
		im[write] = im[read]
		write++
	}
	return im[:write]
}

// MARK: GRID
func gridSORT(r []isoR) []isoR {
	sort.Slice(r, func(i, j int) bool { return r[i].zi > r[j].zi })
	return r
}
func fIsoRnum(r []isoR, num int) (isoR, int) {
	num2 := 0
	r2 := isoR{}
	for i := range r {
		if r[i].num == num {
			num2 = 1
			r2 = r[i]
		}
	}
	return r2, num2
}

// MARK: RENDER
func rTex(rSource, rDest sdl.FRect, tex *sdl.Texture) {
	RND.RenderTexture(tex, &rSource, &rDest)
}

// MARK: COLLISIONS
func cPointIsoGridWalls(p sdl.FPoint, r []isoR) bool {
	collides := false
	for i := range r {
		if r[i].wall {
			if cPointIsoRec(p, r[i]) {
				collides = true
				break
			}
		}
	}
	return collides
}
func cPointIsoRec(p sdl.FPoint, r isoR) bool {
	collides := false
	if cPointTri(p, r.t[0]) || cPointTri(p, r.t[1]) {
		collides = true
	}
	return collides
}
func cIsoRecRec(ir isoR, r sdl.FRect) bool {
	collides := false
	for i := range ir.p {
		if cPointRec(ir.p[i], r) {
			collides = true
			break
		}
	}
	return collides
}
func cPointRec(p sdl.FPoint, r sdl.FRect) bool {
	return p.X >= r.X &&
		p.X <= r.X+r.W &&
		p.Y >= r.Y &&
		p.Y <= r.Y+r.H
}
func cPointTri(p sdl.FPoint, t tri) bool {
	a := t.v[0].Position
	b := t.v[1].Position
	c := t.v[2].Position
	v0 := sdl.FPoint{X: c.X - a.X, Y: c.Y - a.Y}
	v1p := sdl.FPoint{X: b.X - a.X, Y: b.Y - a.Y}
	v2p := sdl.FPoint{X: p.X - a.X, Y: p.Y - a.Y}
	dot00 := v0.X*v0.X + v0.Y*v0.Y
	dot01 := v0.X*v1p.X + v0.Y*v1p.Y
	dot02 := v0.X*v2p.X + v0.Y*v2p.Y
	dot11 := v1p.X*v1p.X + v1p.Y*v1p.Y
	dot12 := v1p.X*v2p.X + v1p.Y*v2p.Y
	denom := dot00*dot11 - dot01*dot01
	if denom == 0 {
		return false
	}
	invDenom := 1.0 / denom
	u := (dot11*dot02 - dot01*dot12) * invDenom
	v := (dot00*dot12 - dot01*dot02) * invDenom
	return (u >= 0) && (v >= 0) && (u+v <= 1)
}

// MARK: RANDOM NUMBERS
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RF32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
