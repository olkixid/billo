package main

import (
	"encoding/xml"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type block struct {
	srcRect image.Rectangle
	dstRect rectangle
}

type level struct {
	grid   [][]*block
	blocks []*block
}

func (l *level) getOverlappingRects(r rectangle) []rectangle {
	overlapping := make([]rectangle, 0, 8)
	for _, block := range l.blocks {
		if block.dstRect.overlaps(r) {
			overlapping = append(overlapping, block.dstRect)
		}
	}
	return overlapping
}

func (l *level) drawTo(target *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ImageParts = l
	err := target.DrawImage(globalTexAtlas.img, op)
	if err != nil {
		fmt.Printf("Drawing error: %v\n", err)
	}
}

func (l *level) Len() int {
	return len(l.blocks)
}

func (l *level) Dst(i int) (x0, y0, x1, y1 int) {
	r := l.blocks[i].dstRect
	return int(r.x), int(r.y), int(r.x + r.w), int(r.y + r.h)
}

func (l *level) Src(i int) (x0, y0, x1, y1 int) {
	r := l.blocks[i].srcRect
	return r.Min.X, r.Min.Y, r.Max.X, r.Max.Y
}

func loadLevel(fileName string) *level {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("An error occurred while opening Level XML: %v\n", err)
		return nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("An error occurred while reading Level XML: %v\n", err)
		return nil
	}

	xmlRoot := XMLLevel{}
	err = xml.Unmarshal(data, &xmlRoot)
	if err != nil {
		fmt.Printf("An error occurred while parsing Level XML: %v\n", err)
		return nil
	}
	//fmt.Println(xmlRoot.Grid.Rows)
	srcRects := make(map[int]image.Rectangle)
	for _, e := range xmlRoot.Enum {
		id, err := strconv.Atoi(e.TileId)
		if err != nil {
			fmt.Printf("An error occurred while parsing TileId (Level XML): %v\n", err)
			fmt.Println("Tile will be skipped!")
		}

		srcRects[id] = globalTexAtlas.subTexRects[e.TileName]
	}

	//fmt.Println(srcRects)

	var idGrid [][]int
	for r, row := range xmlRoot.Grid.Rows {
		idGrid = append(idGrid, []int{})
		splittedRow := strings.Split(row, ";")
		fmt.Println(splittedRow[0])
		for _, element := range splittedRow {
			tileId, _ := strconv.Atoi(element)
			idGrid[r] = append(idGrid[r], tileId)
		}
	}

	//fmt.Println(idGrid)

	var maxRowLen int
	for _, row := range idGrid {
		if len(row) > maxRowLen {
			maxRowLen = len(row)
		}
	}

	l := level{}
	l.grid = make([][]*block, len(idGrid))
	for i := range l.grid {
		l.grid[i] = make([]*block, maxRowLen)
	}

	for r, row := range idGrid {
		for c, tileId := range row {
			sr := srcRects[tileId]
			if sr != image.ZR {
				freshBlock := block{}
				freshBlock.srcRect = sr
				freshBlock.dstRect = rectangle{float64(c * 70), float64(r * 70), 70, 70}
				l.grid[r][c] = &freshBlock
				l.blocks = append(l.blocks, &freshBlock)
			}
		}
	}

	return &l
}

type XMLLevel struct {
	XMLName xml.Name      `xml:"Level"`
	Enum    []XMLTileEnum `xml:"TileEnum"`
	Grid    XMLGrid       `xml:"Grid"`
}

type XMLTileEnum struct {
	XMLName  xml.Name `xml:"TileEnum"`
	TileName string   `xml:"tileName,attr"`
	TileId   string   `xml:"tileId,attr"`
}

type XMLGrid struct {
	XMLName xml.Name `xml:"Grid"`
	Rows    []string `xml:"Row"`
}
