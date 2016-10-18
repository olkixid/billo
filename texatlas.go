package main

import (
	"encoding/xml"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type texAtlas struct {
	img         *ebiten.Image
	subTexRects map[string]image.Rectangle
}

func loadTexAtlas(fileName string) *texAtlas {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("An error occurred while opening Texure Atlas XML: %v\n", err)
		return nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("An error occurred while reading Texure Atlas XML: %v\n", err)
		return nil
	}

	xmlRoot := XMLTexAtlas{}
	err = xml.Unmarshal(data, &xmlRoot)
	if err != nil {
		fmt.Printf("An error occurred while parsing Texure Atlas XML: %v\n", err)
		return nil
	}

	path := filepath.Dir(fileName)

	img, _, err := ebitenutil.NewImageFromFile(filepath.Join(path, xmlRoot.ImagePath), ebiten.FilterLinear)
	if err != nil {
		fmt.Printf("An error occurred while loading Texure Atlas Image File: %v\n", err)
		return nil
	}

	ta := texAtlas{img, make(map[string]image.Rectangle)}

	for _, s := range xmlRoot.Sprites {
		x, xErr := strconv.Atoi(s.XPos)
		y, yErr := strconv.Atoi(s.YPos)
		w, wErr := strconv.Atoi(s.Width)
		h, hErr := strconv.Atoi(s.Height)

		if xErr != nil || yErr != nil || wErr != nil || hErr != nil {
			fmt.Printf("An error occurred while atoiing coordinates(Texure Atlas Sprites). \n")
		}

		rect := image.Rect(x, y, x+w, y+h)

		if rect.In(img.Bounds()) && len(s.Name) > 0 {
			ta.subTexRects[s.Name] = rect
		}
	}

	fmt.Println("banana")
	return &ta
}

type XMLTexAtlas struct {
	XMLName   xml.Name            `xml:"TextureAtlas"`
	ImagePath string              `xml:"imagePath,attr"`
	Width     string              `xml:"width,attr"`
	Height    string              `xml:"height,attr"`
	Sprites   []XMLTexAtlasSprite `xml:"sprite"`
}

type XMLTexAtlasSprite struct {
	XMLName xml.Name `xml:"sprite"`
	Name    string   `xml:"n,attr"`
	XPos    string   `xml:"x,attr"`
	YPos    string   `xml:"y,attr"`
	Width   string   `xml:"w,attr"`
	Height  string   `xml:"h,attr"`
}
