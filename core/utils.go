package core

import (
	"golang.org/x/image/draw"
	"image"
	"math"
)

func gradientVector(x, y float64) (float64, float64) {
	magnitude := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	orientation := math.Mod((math.Atan2(x, y) * 180 / math.Pi), float64(360))
	return magnitude, orientation
}
func drawLine(start, end image.Point, angle, length float64) {
	// x2 := x1 + (length * math.Cos(angle))
	// y2 := y1 + (length * math.Sin(angle))

}

func scaleImage(img image.Image, size int) image.Image {
	rect := image.Rect(0, 0, img.Bounds().Max.X/size, img.Bounds().Max.Y/size)
	dstimg := image.NewRGBA(rect)
	draw.ApproxBiLinear.Scale(dstimg, rect, img, img.Bounds(), draw.Over, nil)
	return dstimg
}
