package core

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"
)

func drawsquare(img image.Image, r image.Rectangle, density int, c color.Color) image.Image {
	nimg := image.NewRGBA(img.Bounds())
	draw.Draw(nimg, nimg.Bounds(), img, image.ZP, draw.Src)
	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			if x == r.Min.X || x == r.Max.X || y == r.Min.Y || y == r.Max.Y {
				for i := density; i >= 0; i-- {
					nimg.Set(x-i, y-i, c)
					nimg.Set(x+i, y+i, c)
				}
			}
		}
	}
	return nimg
}
