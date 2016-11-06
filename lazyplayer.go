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
	grounded        bool
}

func newLazyPlayer(width float64, heigth float64) *lazyPlayer {
	lp := lazyPlayer{rectangle{70, 70, width, heigth}, 0, 0, false, false}
	return &lp
}

func (lp *lazyPlayer) update() {
	lp.xSpeed = 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) && !lp.lastTimeJumpKey && lp.grounded {
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

	var leftCollision, rightCollision, upCollision, downCollision bool

	lp.rect.x += lp.xSpeed
	for _, collider := range colliders {
		if lp.rect.overlaps(collider) {
			if lp.xSpeed > 0 {
				lp.rect.x = collider.x - lp.rect.w
				rightCollision = true
			} else if lp.xSpeed < 0 {
				lp.rect.x = collider.x + collider.w
				leftCollision = true
			}
		}
	}

	lp.rect.y += lp.ySpeed
	for _, collider := range colliders {
		if lp.rect.overlaps(collider) {
			if lp.ySpeed > 0 {
				lp.rect.y = collider.y - lp.rect.h
				downCollision = true
			} else if lp.ySpeed < 0 {
				lp.rect.y = collider.y + collider.h
				upCollision = true
			}
		}
	}

	if leftCollision || rightCollision {
		lp.xSpeed = 0
	}
	if upCollision || downCollision {
		lp.ySpeed = 0
	}

	if downCollision {
		lp.grounded = true
	} else {
		lp.grounded = false
	}
}
