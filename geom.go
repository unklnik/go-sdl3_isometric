package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/sdl"
)

var ()

type isoR struct {
	p       []sdl.FPoint
	t       []tri
	cnt     sdl.FPoint
	zi, num int
	im      IM
	rD      sdl.FRect
	wall    bool
}
type isoG struct {
	r []isoR
}
type tri struct {
	v []sdl.Vertex
}

// MARK: REC
func dRecFillAlpha(r sdl.FRect, c sdl.Color, a uint8) {
	c = COLALPHA(c, a)
	COL(c)
	RND.RenderFillRect(&r)
}
func dRecLine(r sdl.FRect, c sdl.Color) {
	COL(c)
	RND.RenderRect(&r)
}
func mRecCnt(cnt sdl.FPoint, w, h float32) sdl.FRect {
	return sdl.FRect{cnt.X - w/2, cnt.Y - h/2, w, h}
}
func mRecSmallerOffset(r sdl.FRect, offset float32) sdl.FRect {
	return sdl.FRect{r.X + offset, r.Y + offset, r.W - offset*2, r.H - offset*2}
}
func dRecFillCnt(cnt sdl.FPoint, w, h float32, c sdl.Color) {
	COL(c)
	RND.RenderFillRect(&sdl.FRect{cnt.X - w/2, cnt.Y - h/2, w, h})
}

// MARK: TRIANGLES
func dTri(t tri, c sdl.Color) {
	t = cTri(t, c)
	RND.RenderGeometry(nil, t.v, nil)
}
func cTri(t tri, c sdl.Color) tri {
	for i := range t.v {
		t.v[i].Color = COL2F32COL(c)
	}
	return t
}
func mTri(p []sdl.FPoint) tri {
	t := tri{}
	v := sdl.Vertex{}
	v.Position = p[0]
	t.v = append(t.v, v)
	v.Position = p[1]
	t.v = append(t.v, v)
	v.Position = p[2]
	t.v = append(t.v, v)
	return t
}

// MARK: ISOREC
func mIsoRzi(p sdl.FPoint, w float32, zi int) isoR {
	r := mIsoR(p, w)
	r.zi = zi
	return r
}
func mIsoR(p sdl.FPoint, w float32) isoR {
	r := isoR{}
	r.p = append(r.p, p)
	p.X -= w
	p.Y -= w / 2
	r.p = append(r.p, p)
	p.X += w
	p.Y -= w / 2
	r.p = append(r.p, p)
	p.X += w
	p.Y += w / 2
	r.p = append(r.p, p)
	t := mTri([]sdl.FPoint{r.p[0], r.p[1], r.p[2]})
	r.t = append(r.t, t)
	t = mTri([]sdl.FPoint{r.p[0], r.p[2], r.p[3]})
	r.t = append(r.t, t)
	r.cnt = r.p[0]
	r.cnt.Y -= w / 2
	r.rD = sdl.FRect{r.p[1].X, r.p[0].Y - w*2, w * 2, w * 2}
	return r
}
func dIsoRline(r isoR, c sdl.Color) {
	dLINEp(r.p[0], r.p[1], c)
	dLINEp(r.p[1], r.p[2], c)
	dLINEp(r.p[2], r.p[3], c)
	dLINEp(r.p[0], r.p[3], c)
}
func dIsoRfill(r isoR, c sdl.Color) {
	dTri(r.t[0], c)
	dTri(r.t[1], c)
}

// MARK: GRID DRAW
func dGridWalls(r []isoR) {
	playDrawn := false
	for i := range r {
		if !playDrawn && r[i].zi == pl.zi {
			dPL()
			playDrawn = true
		}
		if r[i].wall {
			dIMRec(wallBlok, r[i].rD)
		}
	}
}
func dGridMouse(r []isoR, c sdl.Color) {
	for i := range r {
		if cPointTri(MS, r[i].t[0]) || cPointTri(MS, r[i].t[1]) {
			dIsoRfill(r[i], c)
		}
	}
}
func dGridIM(r []isoR) {
	for i := range r {
		dIMRec(r[i].im, r[i].rD)
	}
}
func dGridIMoffsetY(r []isoR) {
	for i := range r {
		r2 := r[i].rD
		r2.Y += r[i].rD.H / 2
		dIMRec(r[i].im, r2)
	}
}
func dGrid(r []isoR, cLine, cFill sdl.Color) {
	for i := range r {
		dIsoRfill(r[i], cFill)
		dIsoRline(r[i], cLine)
	}
	if dBUG {
		for i := range r {
			if dBUG {
				RND.SetDrawColor(COL2RGBA(cLine))
				RND.RenderPoint(r[i].cnt.X, r[i].cnt.Y)
				dTxtXY("z"+fmt.Sprint(r[i].zi), r[i].cnt.X, r[i].cnt.Y, 1)
			}
		}
	}
}
func dGridLines(r []isoR, cLine sdl.Color) {
	for i := range r {
		dIsoRline(r[i], cLine)
	}
	if dBUG {
		for i := range r {
			if dBUG {
				RND.SetDrawColor(COL2RGBA(cLine))
				RND.RenderPoint(r[i].cnt.X, r[i].cnt.Y)
				//dTxtXYcolor("z"+fmt.Sprint(r[i].zi), r[i].cnt.X, r[i].cnt.Y, 1, WHITE())
				//dTxtXYcolor("n"+fmt.Sprint(r[i].num), r[i].cnt.X, r[i].cnt.Y-FONT1.hS, 1, WHITE())
			}
		}
	}
}

// MARK: GRID MAKE
func mGridInnerWalls(r []isoR) []isoR {

	num2 := RINT(5, 15)
	for num2 > 0 {
		choose := RINT(0, 11)
		//choose = 9
		switch choose {
		case 10: //LINE SPACE DOWN
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum+(levW*2)].wall = true
					r[bloknum-(levW*2)].wall = true
					break
				}
			}
		case 9: //LINE SPACE RIGHT
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-2].wall = true
					r[bloknum-4].wall = true
					break
				}
			}
		case 8: //5 BLOK CROSS
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum-1].wall = true
					r[(bloknum-2)-levW].wall = true
					r[(bloknum-1)-levW].wall = true
					r[(bloknum-1)-levW*2].wall = true
					r[bloknum-levW].wall = true
					break
				}
			}
		case 7: //4 BLOK CROSS
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-2].wall = true
					r[bloknum-levW*2].wall = true
					r[(bloknum-2)-levW*2].wall = true
					break
				}
			}
		case 6: //RANDOM 1 BLOK
			num := RINT(3, 10)
			for num > 0 {
				for {
					bloknum := RINT(levW*2, len(LEV)-levW*2)
					if cPointIsoRec(r[bloknum].p[0], inner) {
						r[bloknum].wall = true
						num--
						break
					}
				}
			}
		case 5: //C UP
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-2].wall = true
					r[bloknum-levW].wall = true
					r[bloknum-(levW+2)].wall = true
					r[bloknum-levW*2].wall = true
					r[bloknum-((levW*2)+1)].wall = true
					r[bloknum-((levW*2)+2)].wall = true
					break
				}
			}
		case 4: //C LEFT
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-1].wall = true
					r[bloknum-2].wall = true
					r[bloknum-(levW+2)].wall = true
					r[bloknum-levW*2].wall = true
					r[bloknum-((levW*2)+1)].wall = true
					r[bloknum-((levW*2)+2)].wall = true
					break
				}
			}
		case 3: //C DOWN
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-1].wall = true
					r[bloknum-2].wall = true
					r[bloknum-levW].wall = true
					r[bloknum-(levW+2)].wall = true
					r[bloknum-levW*2].wall = true
					r[bloknum-((levW*2)+2)].wall = true
					break
				}
			}
		case 2: //C RIGHT
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-1].wall = true
					r[bloknum-2].wall = true
					r[bloknum-levW].wall = true
					r[bloknum-levW*2].wall = true
					r[bloknum-((levW*2)+1)].wall = true
					r[bloknum-((levW*2)+2)].wall = true
					break
				}
			}
		case 1: //9 BLOK
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-1].wall = true
					r[bloknum-2].wall = true
					r[bloknum-levW].wall = true
					r[bloknum-(levW+1)].wall = true
					r[bloknum-(levW+2)].wall = true
					r[bloknum-levW*2].wall = true
					r[bloknum-((levW*2)+1)].wall = true
					r[bloknum-((levW*2)+2)].wall = true
					break
				}
			}
		case 0: //4 BLOK
			for {
				bloknum := RINT(levW*2, len(LEV)-levW*2)
				if cPointIsoRec(r[bloknum].p[0], inner) {
					r[bloknum].wall = true
					r[bloknum-1].wall = true
					r[bloknum-levW].wall = true
					r[bloknum-(levW+1)].wall = true
					break
				}
			}
		}
		num2--
	}

	return r
}
func mGridWallsEdge(r []isoR) []isoR {
	for i := range r {
		if i < levW {
			r[i].wall = true
		} else if i >= len(r)-levW {
			r[i].wall = true
		}
		if r[i].num%levW == 0 {
			r[i].wall = true
		} else if (r[i].num+1)%levW == 0 {
			r[i].wall = true
		}
	}
	return r
}
func mGridIM(r []isoR, im IM) []isoR {
	for i := range r {
		r[i].im = im
	}

	return r
}
func mGridCNT(w float32, numW, numH int) []isoR {
	var r []isoR
	p := CNT
	h := w * float32(numH)
	p.Y += h / 2
	a := numW * numH
	c := 0
	c2 := 0
	op := p
	zi := 0
	ozi := zi
	for a > 0 {
		ir := mIsoRzi(p, w, zi)
		ir.num = c2
		r = append(r, ir)
		p.X -= w
		p.Y -= w / 2
		c++
		c2++
		zi++
		a--
		if c == numW {
			zi = ozi
			zi++
			ozi = zi
			c = 0
			p = op
			p.X += w
			p.Y -= w / 2
			op = p
		}
	}

	return r
}

// MARK: LINES
func dLINEp(p1, p2 sdl.FPoint, c sdl.Color) {
	RND.SetDrawColor(COL2RGBA(c))
	RND.RenderLine(p1.X, p1.Y, p2.X, p2.Y)
}
