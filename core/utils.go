package core

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
)

func gradientVector(x, y float64) (float64, float64) {
	magnitude := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	orientation := math.Mod((math.Atan2(x, y) * 180 / math.Pi), float64(360))
	return magnitude, orientation
}
func drawLine(p image.Point, angle, length float64, img image.Image) image.Image {
	nimg := image.NewRGBA(img.Bounds())
	x2 := math.Round(float64(p.X) + (length * math.Cos(angle)))
	y2 := math.Round(float64(p.Y) + (length * math.Sin(angle)))
	slop := (x2 - float64(p.X)) / (y2 - float64(p.Y))
	b := int(float64(p.Y) - slop*float64(p.X))
	for x := p.X; x <= int(x2); x++ {
		nimg.Set(x, int(slop*float64(x))+b, color.White)
	}
	return nimg
}

func scaleImage(img image.Image, size int) image.Image {
	rect := image.Rect(0, 0, img.Bounds().Max.X/size, img.Bounds().Max.Y/size)
	dstimg := image.NewRGBA(rect)
	draw.ApproxBiLinear.Scale(dstimg, rect, img, img.Bounds(), draw.Over, nil)
	return dstimg
}
