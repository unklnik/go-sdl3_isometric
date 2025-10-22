package main

import (
	"math/rand/v2"

	"github.com/Zyko0/go-sdl3/sdl"
)

// UTILS
func RUINT8COL(min, max int) uint8 {
	if max > 255 {
		max = 255
	}
	if min < 0 {
		min = 0
	}
	return uint8(min + rand.IntN(max-min))
}
func COL2RGBA(c sdl.Color) (uint8, uint8, uint8, uint8) {
	return c.R, c.G, c.B, c.A
}
func COL2F32COL(c sdl.Color) sdl.FColor {
	return sdl.FColor{float32(c.R), float32(c.G), float32(c.B), float32(c.A)}
}
func COLALPHA(c sdl.Color, a uint8) sdl.Color {
	c.A = a
	return c
}
func COL(c sdl.Color) {
	RND.SetDrawColor(COL2RGBA(c))
}

// COLORS
func RED() sdl.Color {
	return sdl.Color{255, 0, 0, 255}
}
func GREEN() sdl.Color {
	return sdl.Color{0, 255, 0, 255}
}
func DARKGREEN() sdl.Color {
	return sdl.Color{1, 50, 32, 255}
}
func BLUE() sdl.Color {
	return sdl.Color{0, 0, 255, 255}
}
func MAGENTA() sdl.Color {
	return sdl.Color{255, 0, 144, 255}
}
func ORANGE() sdl.Color {
	return sdl.Color{255, 165, 0, 255}
}
func WHITE() sdl.Color {
	return sdl.Color{255, 255, 255, 255}
}
func BLACK() sdl.Color {
	return sdl.Color{0, 0, 0, 255}
}
