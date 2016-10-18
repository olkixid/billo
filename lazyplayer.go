package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type lazyPlayer struct {
	rect   rectangle
	xSpeed float32
	ySpeed float32
}

func newLazyPlayer(width float32, heigth float32) *lazyPlayer {
	lp := lazyPlayer{rectangle{0, 0, width, heigth}, 0, 0}
	return &lp
}

func (lp *lazyPlayer) update() {
	lp.xSpeed = 0
	lp.ySpeed = 0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		lp.ySpeed += -8
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		lp.ySpeed += 8
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		lp.xSpeed += -8
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		lp.xSpeed += 8
	}
	lp.rect.x += lp.xSpeed
	lp.rect.y += lp.ySpeed
}

func (lp *lazyPlayer) draw(target *ebiten.Image) {
	lp.rect.drawFilled(target, color.RGBA{0x00, 0xff, 0x00, 0xff})
}

func (lp *lazyPlayer) checkCollision(rect rectangle) {
	if lp.rect.overlaps(rect) {

		ir := lp.rect.intersect(rect)

		if ir.w < ir.h {
			lpMiddleX := lp.rect.x + lp.rect.w/2
			irMiddleX := ir.x + ir.w/2
			if lpMiddleX > irMiddleX {
				lp.rect.x += ir.w
			} else {
				lp.rect.x -= ir.w
			}
		} else {
			lpMiddleY := lp.rect.y + lp.rect.h/2
			irMiddleY := ir.y + ir.h/2
			if lpMiddleY > irMiddleY {
				lp.rect.y += ir.h
			} else {
				lp.rect.y -= ir.h
			}
		}
	}
}
