package core

import (
	"image"
	"image/color"

	"golang.org/x/image/draw"
)

//DrawSquare to draw squareline in a given image.
func DrawSquare(img image.Image, r image.Rectangle, density int, c color.Color) image.Image {
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

//DrawLine draw a line in image.
func DrawLine(p image.Point, angle, length float64, imgsrc image.Image, c color.Color) *image.RGBA {
	bound := imgsrc.Bounds()
	dstimg, mask := image.NewRGBA(bound), image.NewRGBA(bound)
	x1 := xFromAngle(p.X, length, angle)
	y1 := yFromAngle(p.Y, lenght, angle)
	x2 := xFromAngle(p.X, length, angle+180)
	a := (x1 - float64(p.X)) / (y1 - float64(p.Y))
	b := int(float64(p.Y) - a*float64(p.X))
	s, e := x2, x1
	if x1 < 0 {
		s, e = x1, x2
	}
	for x := int(s); x <= int(e); x++ {
		mask.Set(x, int(a*float64(x))+b, c)
	}
	draw.Draw(dstimg, bound, imgsrc, bound.Min, draw.Src)
	draw.Draw(dstimg, bound, mask, bound.Min, draw.Over)
	return dstimg
}
func drawsquareI(imgsrc, mask image.Image, p image.Point) image.Image {
	dstimg := image.NewRGBA(imgsrc.Bounds())
	draw.Draw(dstimg, imgsrc.Bounds(), imgsrc, image.ZP, draw.Src)
	draw.Draw(dstimg, mask.Bounds(), mask, image.ZP, draw.Over)
	return dstimg
}
