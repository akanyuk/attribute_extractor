package parser

import (
	"fmt"
	"image"
	"image/png"
	"io"
)

func Parse(reader io.Reader) ([][]string, error) {
	img, err := png.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("png decode error: %s", err)
	}

	rows := make([][]string, 0)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y += 8 {
		row := make([]string, 0)
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x += 8 {
			attribute, err := getAttribute(img, x, y)
			if err != nil {
				return nil, err
			}

			row = append(row, fmt.Sprintf("#%02x", attribute))
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func getAttribute(img image.Image, x int, y int) (int, error) {
	firstAttribute, err := rgbaToPixel(img.At(x, y).RGBA()).toAttribute()
	if err != nil {
		return 0, err
	}

	for y1 := y; y1 < y+8 && y1 < img.Bounds().Max.Y; y1++ {
		for x1 := x; x1 < x+8 && x1 < img.Bounds().Max.X; x1++ {
			attr, err := rgbaToPixel(img.At(x1, y1).RGBA()).toAttribute()
			if err != nil {
				return 0, err
			}

			if attr == firstAttribute {
				continue
			}

			return mixAttributes(firstAttribute, attr), nil
		}
	}

	return mixAttributes(firstAttribute, Attribute{}), nil
}

func mixAttributes(ink Attribute, paper Attribute) int {
	var bright int
	if ink.Bright || paper.Bright {
		bright = 0x40
	}

	return paper.Color<<3 + ink.Color + bright
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
