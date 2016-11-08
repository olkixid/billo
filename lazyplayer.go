package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
)

type lazyPlayer struct {
	rect   rectangle
	xSpeed float64
	ySpeed float64

	lastTimeJumpKey bool
	grounded        bool
	anim            animation
}

func newLazyPlayer(width float64, heigth float64) *lazyPlayer {
	lp := lazyPlayer{rectangle{70, 70, 0, 0}, 0, 0, false, false, animation{}}
	lp.anim.init()
	lp.rect.w, lp.rect.h = float64(lp.anim.running[0].Dx()), float64(lp.anim.running[0].Dy())
	fmt.Println(lp.rect)
	return &lp
}

func (lp *lazyPlayer) update() {
	lp.xSpeed = 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) && !lp.lastTimeJumpKey && lp.grounded {
		lp.ySpeed += -12
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

func (lp *lazyPlayer) checkCollisions(lv *level) {
	checkRect := lp.rect

	if lp.xSpeed < 0 {
		checkRect.x -= -lp.xSpeed
		checkRect.w += -lp.xSpeed
	} else {
		checkRect.w += lp.xSpeed
	}

	if lp.ySpeed < 0 {
		checkRect.y -= -lp.ySpeed
		checkRect.h += -lp.ySpeed
	} else {
		checkRect.h += lp.ySpeed
	}

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

//Drawing Stuff
func (lp *lazyPlayer) draw(target *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = lp
	err := target.DrawImage(globalTexAtlas.img, op)
	if err != nil {
		fmt.Printf("Drawing error: %v\n", err)
	}
}

func (lazyPlayer) Len() int {
	return 1
}

func (lp *lazyPlayer) Dst(i int) (x0, y0, x1, y1 int) {
	r := lp.rect
	if lp.xSpeed < 0 {
		r.x += r.w
		r.w = -r.w
	}
	return int(r.x), int(r.y), int(r.x + r.w), int(r.y + r.h)
}

func (lp *lazyPlayer) Src(i int) (x0, y0, x1, y1 int) {
	running := false
	if lp.xSpeed != 0 {
		running = true
	}
	imgR := lp.anim.getImgRect(running)
	return imgR.Min.X, imgR.Min.Y, imgR.Max.X, imgR.Max.Y
}

type animation struct {
	running             [7]image.Rectangle
	currentRunningState int
	front               image.Rectangle
}

func (an *animation) init() {
	an.front = globalTexAtlas.subTexRects["p3_front"]
	an.running[0] = globalTexAtlas.subTexRects["p3_walk01"]
	an.running[1] = globalTexAtlas.subTexRects["p3_walk02"]
	an.running[2] = globalTexAtlas.subTexRects["p3_walk03"]
	an.running[3] = globalTexAtlas.subTexRects["p3_walk04"]
	an.running[4] = globalTexAtlas.subTexRects["p3_walk05"]
	an.running[5] = globalTexAtlas.subTexRects["p3_walk06"]
	an.running[6] = globalTexAtlas.subTexRects["p3_walk07"]
}

func (an *animation) getImgRect(running bool) image.Rectangle {
	if !running {
		an.currentRunningState = -1
		return an.front
	}
	an.currentRunningState++
	if an.currentRunningState >= len(an.running) {
		an.currentRunningState = 0
	}
	return an.running[an.currentRunningState]
}
