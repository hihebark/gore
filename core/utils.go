package core

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
)

//GradientVector Computing The Gradient Image and return the magnitude and orientation.
func GradientVector(x, y float64) (float64, float64) {
	// http://mccormickml.com/2013/05/07/gradient-vectors/
	magnitude := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	orientation := math.Mod((math.Atan2(x, y) * 180 / math.Pi), float64(360))
	return magnitude, orientation
}

//DrawLine draw a line in image.
func DrawLine(p image.Point, angle, length float64, img image.Image, c color.Color) *image.RGBA {
	bounds := img.Bounds()
	mask, dstimg := image.NewRGBA(bounds), image.NewRGBA(bounds)
	x2 := math.Round(float64(p.X) + (length * math.Cos(angle)))
	y2 := math.Round(float64(p.Y) + (length * math.Sin(angle)))
	x2m := math.Round(float64(p.X) + (length * math.Cos(angle+180)))
	slop := (x2 - float64(p.X)) / (y2 - float64(p.Y))
	b := int(float64(p.Y) - slop*float64(p.X))
	for x := int(x2m); x <= int(x2); x++ {
		mask.Set(x, int(slop*float64(x))+b, c)
	}
	draw.Draw(dstimg, img.Bounds(), img, bounds.Min, draw.Src)
	draw.Draw(dstimg, mask.Bounds(), mask, bounds.Min, draw.Over)
	return dstimg
}
