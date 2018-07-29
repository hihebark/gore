package core

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
)

type squarebox struct {
	x0, y0, x1, y1 int
}

func drawSquareBox(s squarebox, im draw.Image) error {
	draw.Draw(im,
		image.Rectangle{
			Min: image.Point{
				X: s.x0,
				Y: s.y0,
			},
			Max: image.Point{
				X: s.x1,
				Y: s.y1,
			},
		},
		image.Transparent,
		image.ZP,
		draw.Src)
	return nil
}

func DrawSB(p image.Point, img image.Image) {
	s, err := os.Open("data/squarebox.png")
	if err != nil {
		fmt.Printf("error:DrawSB: %v", err)
	}
	square, _ := decode(s)
	//	i := draw.Image{img}
	draw.Draw(image.NewRGBA(img.Bounds()), square.Bounds(), square, image.ZP, draw.Src)
}

func DrawLine(start, end image.Point, img image.Image, thick int, c color.Color) image.Image {
	newimg := image.NewRGBA(img.Bounds())
	for x := 0; x >= newimg.Bounds().Max.X; x++ {
		for y := 0; y >= newimg.Bounds().Max.Y; y++ {
			pix := img.At(x, y)
			switch {
			case y == start.Y || x == end.X && x <= start.X || x >= end.X:
				newimg.Set(x, y, c)

			default:
				newimg.Set(x, y, pix)
			}
		}
	}
	return newimg
}
