package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

var pixel *ebiten.Image

func createPixel() {
	fmt.Println("alive")
	pixel, _ = ebiten.NewImage(1, 1, ebiten.FilterLinear)
	pixel.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})

	fmt.Println("alive")
}

type rectangle struct {
	x float32
	y float32
	w float32
	h float32
}

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

func (r rectangle) drawFilledTo(target *ebiten.Image) {
	rip := rectangleFillImageParts{int(r.x), int(r.y), int(r.x + r.w), int(r.y + r.h)}
	opt := ebiten.DrawImageOptions{}
	opt.ImageParts = rip
	target.DrawImage(pixel, &opt)
}
