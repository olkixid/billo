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
	x float64
	y float64
	w float64
	h float64
}

func (r rectangle) overlaps(otherRect rectangle) bool {
	if r.x+r.w > otherRect.x && otherRect.x+otherRect.w > r.x {
		if r.y+r.h > otherRect.y && otherRect.y+otherRect.h > r.y {
			return true
		}
	}
	return false
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
