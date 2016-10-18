package main

import (
	"encoding/xml"
	"fmt"
	"image"
	"io/ioutil"
	"os"
)

type block struct {
	srcImgRect image.Rectangle
	dstRect    rectangle
}

type level struct {
	blocks []block
	grid   [][]*block
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

	fmt.Println(xmlRoot)
	return &level{}
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
