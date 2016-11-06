package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type lazyPlayer struct {
	rect   rectangle
	xSpeed float64
	ySpeed float64

	lastTimeJumpKey bool
}

func newLazyPlayer(width float64, heigth float64) *lazyPlayer {
	lp := lazyPlayer{rectangle{70, 70, width, heigth}, 0, 0, false}
	return &lp
}

func (lp *lazyPlayer) update() {
	lp.xSpeed = 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) && !lp.lastTimeJumpKey {
		lp.ySpeed += -11
	}
	lp.lastTimeJumpKey = ebiten.IsKeyPressed(ebiten.KeyUp)

	/*
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			lp.ySpeed += 8
		}
	*/
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		lp.xSpeed += -5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		lp.xSpeed += 5
	}

	lp.ySpeed += 0.5
}

func (lp *lazyPlayer) draw(target *ebiten.Image) {
	lp.rect.drawFilled(target, color.RGBA{0x00, 0xff, 0x00, 0xff})
}

func (lp *lazyPlayer) checkCollisions(lv *level) {
	checkRect := lp.rect
	checkRect.x += lp.xSpeed
	checkRect.y += lp.ySpeed

	colliders := lv.getOverlappingRects(checkRect)

	newXSpeed, newYSpeed := lp.xSpeed, lp.ySpeed

	lp.rect.x += lp.xSpeed
	for _, collider := range colliders {
		if lp.rect.overlaps(collider) {
			newXSpeed = 0
			if lp.xSpeed > 0 {
				lp.rect.x = collider.x - lp.rect.w
			} else if lp.xSpeed < 0 {
				lp.rect.x = collider.x + collider.w
			}
		}
	}

	lp.rect.y += lp.ySpeed
	for _, collider := range colliders {
		if lp.rect.overlaps(collider) {
			newYSpeed = 0
			if lp.ySpeed > 0 {
				lp.rect.y = collider.y - lp.rect.h
			} else if lp.ySpeed < 0 {
				lp.rect.y = collider.y + collider.h
			}
		}
	}

	lp.xSpeed, lp.ySpeed = newXSpeed, newYSpeed
}
