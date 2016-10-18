package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

var otto *lazyPlayer
var crashes int
var obstacle rectangle

func main() {
	_ = loadLevel("res/levels/level1.xml")

	otto = newLazyPlayer(20, 40)

	obstacle = rectangle{100, 100, 200, 200}
	ebiten.Run(update, 640, 640, 1, "Your game's title")

}

func update(screen *ebiten.Image) error {

	otto.update()
	otto.checkCollision(obstacle)
	if otto.rect.overlaps(obstacle) {
		fmt.Println("Still!")
	}

	obstacle.drawFilled(screen, color.RGBA{0xff, 0x00, 0x00, 0xff})
	otto.draw(screen)

	return nil
}
