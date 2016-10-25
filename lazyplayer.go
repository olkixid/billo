package main

import (
	"container/list"
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type lazyPlayer struct {
	rect   rectangle
	xSpeed float32
	ySpeed float32

	collidingRects *list.List
}

func newLazyPlayer(width float32, heigth float32) *lazyPlayer {
	lp := lazyPlayer{rectangle{70, 70, width, heigth}, 0, 0, list.New()}
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
}

func (lp *lazyPlayer) draw(target *ebiten.Image) {
	lp.rect.drawFilled(target, color.RGBA{0x00, 0xff, 0x00, 0xff})
}

func (lp *lazyPlayer) checkCollision(rect rectangle) {
	checkRect := lp.rect
	checkRect.x += lp.xSpeed
	checkRect.y += lp.ySpeed

	if checkRect.overlaps(rect) {
		lp.collidingRects.PushBack(rect)
	}
}

func (lp *lazyPlayer) fixPositions() {
	if lp.collidingRects.Front() == nil {
		lp.rect.x += lp.xSpeed
		lp.rect.y += lp.ySpeed
		return
	}

	lp.rect.x += lp.xSpeed
	elm := lp.collidingRects.Front()
	for elm != nil {
		var currentRect rectangle = elm.Value.(rectangle)
		if lp.rect.overlaps(currentRect) {
			ir := lp.rect.intersect(currentRect)
			if lp.xSpeed > 0 {
				lp.rect.x -= ir.w
			} else if lp.xSpeed < 0 {
				lp.rect.x += ir.w
			}
		}

		elm = elm.Next()
	}

	lp.rect.y += lp.ySpeed
	elm = lp.collidingRects.Front()
	for elm != nil {
		var currentRect rectangle = elm.Value.(rectangle)
		if lp.rect.overlaps(currentRect) {
			ir := lp.rect.intersect(currentRect)
			if lp.ySpeed > 0 {
				lp.rect.y -= ir.h
			} else if lp.ySpeed < 0 {
				lp.rect.y += ir.h
			}
		}

		elm = elm.Next()
	}

	lp.collidingRects.Init()
}
