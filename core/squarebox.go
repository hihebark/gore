package core

import (
	"image"
	"image/draw"
)

type squarebox struct {
	x0, y0, x1, y1 int
}

func drawSquareBox (s squarebox, im draw.Image) error{
	draw.Draw(im,
		image.Rectangle{
			Min: image.Point{
				X:s.x0,
				Y:s.y0,
			},
			Max: image.Point{
				X:s.x1,
				Y:s.y1,
			},
		},
		image.Transparent,
		image.ZP,
		draw.Src)
	return nil
}
