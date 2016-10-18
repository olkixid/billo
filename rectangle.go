package main

import (
	//	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

func init() {
	pixel, _ = ebiten.NewImage(1, 1, ebiten.FilterLinear)
}

type rectangle struct {
	x float32
	y float32
	w float32
	h float32
}

func (r rectangle) overlaps(otherRect rectangle) bool {
	if r.x+r.w > otherRect.x && otherRect.x+otherRect.w > r.x {
		if r.y+r.h > otherRect.y && otherRect.y+otherRect.h > r.y {
			return true
		}
	}
	return false
}

func (r rectangle) intersect(otherRect rectangle) rectangle {
	ir := rectangle{}

	w1 := r.x + r.w - otherRect.x
	w2 := otherRect.x + otherRect.w - r.x
	if w1 < w2 {
		ir.w = w1
	} else {
		ir.w = w2
	}
	if ir.w <= 0 {
		return rectangle{}
	}

	h1 := r.y + r.h - otherRect.y
	h2 := otherRect.y + otherRect.h - r.y
	if h1 < h2 {
		ir.h = h1
	} else {
		ir.h = h2
	}
	if ir.h <= 0 {
		return rectangle{}
	}

	if r.x > otherRect.x {
		ir.x = r.x
	} else {
		ir.x = otherRect.x
	}

	if r.y > otherRect.y {
		ir.y = r.y
	} else {
		ir.y = otherRect.y
	}
	return ir
}

//drawing stuff

var pixel *ebiten.Image

type rectangleFillImageParts struct {
	dstx0 int
	dsty0 int
	dstx1 int
	dsty1 int
}

func (rectangleFillImageParts) Len() int {
	return 1
}
func (rectangleFillImageParts) Src(i int) (x0, y0, x1, y1 int) {
	return 0, 0, 1, 1
}
func (rip rectangleFillImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	return rip.dstx0, rip.dsty0, rip.dstx1, rip.dsty1
}

func (r rectangle) drawFilled(target *ebiten.Image, clr color.Color) {
	pixel.Fill(clr)
	rip := rectangleFillImageParts{int(r.x), int(r.y), int(r.x + r.w), int(r.y + r.h)}
	opt := ebiten.DrawImageOptions{}
	opt.ImageParts = rip
	target.DrawImage(pixel, &opt)
}
