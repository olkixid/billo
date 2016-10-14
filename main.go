package main

import (
	//	"fmt"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	/*
		ta := loadTexAtlas("res/someTiles.xml")
		_ = ta

		fmt.Println(ta.subTexRects)
	*/
	createPixel()
	ebiten.Run(update, 320, 640, 1, "Your game's title")

}

func update(screen *ebiten.Image) error {

	rect := rectangle{10, 30, 10, 30}

	rect.drawFilledTo(screen)
	return nil
}
