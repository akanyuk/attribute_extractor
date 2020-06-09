package parser

import "fmt"

var attributeMap = map[Pixel]Attribute{
	Pixel{R: 0, G: 0, B: 0, A: 255}:       {Color: 0},
	Pixel{R: 0, G: 0, B: 192, A: 255}:     {Color: 1},
	Pixel{R: 0, G: 192, B: 0, A: 255}:     {Color: 2},
	Pixel{R: 192, G: 0, B: 192, A: 255}:   {Color: 3},
	Pixel{R: 192, G: 0, B: 0, A: 255}:     {Color: 4},
	Pixel{R: 0, G: 192, B: 192, A: 255}:   {Color: 5},
	Pixel{R: 192, G: 192, B: 0, A: 255}:   {Color: 6},
	Pixel{R: 192, G: 192, B: 192, A: 255}: {Color: 7},
	Pixel{R: 0, G: 0, B: 255, A: 255}:     {Color: 1, Bright: true},
	Pixel{R: 0, G: 255, B: 0, A: 255}:     {Color: 2, Bright: true},
	Pixel{R: 255, G: 0, B: 255, A: 255}:   {Color: 3, Bright: true},
	Pixel{R: 255, G: 0, B: 0, A: 255}:     {Color: 4, Bright: true},
	Pixel{R: 0, G: 255, B: 255, A: 255}:   {Color: 5, Bright: true},
	Pixel{R: 255, G: 255, B: 0, A: 255}:   {Color: 6, Bright: true},
	Pixel{R: 255, G: 255, B: 255, A: 255}: {Color: 7, Bright: true},
}

type Pixel struct {
	R int
	G int
	B int
	A int
}

type Attribute struct {
	Bright bool
	Color  int
}

func (p Pixel) toAttribute() (Attribute, error) {
	attribute, ok := attributeMap[p]
	if !ok {
		return Attribute{}, fmt.Errorf("pixel not defined: %v", p)
	}

	return attribute, nil
}
