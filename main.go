package main

import (
	//	"fmt"
	"github.com/hajimehoshi/ebiten"
	//	"image/color"
)

var otto *lazyPlayer
var lv *level

var globalTexAtlas *texAtlas

func init() {
	globalTexAtlas = loadTexAtlas("res/someTiles.xml")
}

func main() {
	lv = loadLevel("res/levels/level1.xml")

	otto = newLazyPlayer(20, 40)

	ebiten.Run(update, 640, 640, 1, "Your game's title")

}

func update(screen *ebiten.Image) error {

	otto.update()
	otto.checkCollisions(lv)

	lv.drawTo(screen)
	otto.draw(screen)

	return nil
}
