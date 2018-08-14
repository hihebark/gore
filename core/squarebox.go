package core

import (
	"fmt"
	"image"
	"image/color"
	"os"

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
func drawsquareI(img image.Image, p image.Point) image.Image {
	maski, err := os.Open("data/squarebox.png")
	defer maski.Close()
	if err != nil {
		fmt.Printf("image:drawsquareI:os.Open\n")
	}
	mask, _ := decode(maski)
	dstimg := image.NewRGBA(img.Bounds())
	draw.Draw(dstimg, img.Bounds(), img, image.ZP, draw.Src)
	draw.Draw(dstimg, mask.Bounds(), mask, image.ZP, draw.Over)
	return dstimg
}
