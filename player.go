package main

import (
	"math"

	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	pl PLAYER
)

type PLAYER struct {
	r          sdl.FRect
	cnt, v2    sdl.FPoint
	anm        ANIM
	siz        float32
	lr, moving bool
	zi         int

	velX, velY, maxSpd, accel, frict float32
}

// MARK: DRAW
func dPL() {
	//SHADOW
	r := pl.r
	r.Y += r.H / 2
	dIMRec(ETCIM[0], mRecSmallerOffset(r, UN/8))
	//PL IMG
	if pl.lr {
		if pl.moving {
			pl.anm = dAnimRecLoopColor(pl.anm, pl.r, MAGENTA(), true)
		} else {
			pl.anm = dAnimRecLoopColorPart(pl.anm, pl.r, MAGENTA(), true, 2, 4)
		}
	} else {
		if pl.moving {
			pl.anm = dAnimRecLoopColor(pl.anm, pl.r, MAGENTA(), false)
		} else {
			pl.anm = dAnimRecLoopColorPart(pl.anm, pl.r, MAGENTA(), false, 2, 4)
		}
	}

	if dBUG {
		dRecFillAlpha(pl.r, MAGENTA(), 80)
		dRecLine(pl.r, MAGENTA())
		dRecFillCnt(pl.v2, 8, 8, ORANGE())
	}
}

// MARK: CHECK
func cPLmove(x, y float32) bool {
	canmove := true
	v := pl.cnt
	v.X += x * pl.maxSpd
	v.Y += y * pl.maxSpd
	v.Y += pl.siz / 2
	if cPointIsoGridWalls(v, LEV) {
		canmove = false
	}

	return canmove
}

// MARK: UP MOVE
func movePL() {
	//PLAYER MOVEMENT

	var dx, dy float32
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_UP] || keys[sdl.SCANCODE_W] {
		dy -= 1
	}
	if keys[sdl.SCANCODE_DOWN] || keys[sdl.SCANCODE_S] {
		dy += 1
	}
	if keys[sdl.SCANCODE_LEFT] || keys[sdl.SCANCODE_A] {
		pl.lr = true
		dx -= 1
	}
	if keys[sdl.SCANCODE_RIGHT] || keys[sdl.SCANCODE_D] {
		pl.lr = false
		dx += 1
	}

	// --- Normalize movement ---
	pl.moving = false
	if dx != 0 || dy != 0 {
		pl.moving = true
		len := float32(math.Sqrt(float64(dx*dx + dy*dy)))
		dx /= len
		dy /= len
		if cPLmove(dx, dy) {
			pl.cnt.X += dx * pl.maxSpd
			pl.cnt.Y += dy * pl.maxSpd

			//ZINDEX
			for i := range LEV {
				if cPointIsoRec(pl.v2, LEV[i]) {
					pl.zi = LEV[i].zi
					break
				}
			}
		}
	}

}
func uPL() {
	pl.r = mRecCnt(pl.cnt, pl.siz, pl.siz)
	pl.v2 = pl.cnt
	pl.v2.Y += pl.r.H / 2

}

// MARK: MAKE
func mPL() {
	rcnt, _ := fIsoRnum(LEV, 86)
	pl = PLAYER{}
	pl.cnt = rcnt.cnt
	pl.siz = UN + UN/2
	pl.r = mRecCnt(pl.cnt, pl.siz, pl.siz)
	pl.maxSpd = UN / 5
	pl.anm = mAnimIMSheet(PLIM, 20)
	pl.v2 = pl.cnt
	pl.v2.Y += pl.r.H / 2
}
