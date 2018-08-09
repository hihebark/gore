package core

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
)

func gradientVector(x, y float64) (float64, float64) {
	// http://mccormickml.com/2013/05/07/gradient-vectors/
	// magnitude   := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	// orientation := (math.Atan2(x, y) * 180 / math.Pi ) % 360

	magnitude := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	orientation := math.Mod((math.Atan2(x, y) * 180 / math.Pi), float64(360))
	return magnitude, orientation
}
func drawLine(p image.Point, angle, length float64, img image.Image) image.Image {
	mask, dst := image.NewRGBA(img.Bounds()), image.NewRGBA(img.Bounds())
	x2 := math.Round(float64(p.X) + (length * math.Cos(angle)))
	y2 := math.Round(float64(p.Y) + (length * math.Sin(angle)))
	slop := (x2 - float64(p.X)) / (y2 - float64(p.Y))
	b := int(float64(p.Y) - slop*float64(p.X))
	for x := p.X; x <= int(x2); x++ {
		mask.Set(x, int(slop*float64(x))+b, color.White)
	}
	draw.Draw(dst, img.Bounds(), img, image.ZP, draw.Src)
	draw.Draw(dst, mask.Bounds(), mask, image.ZP, draw.Over)
	return dst
}

func scaleImage(img image.Image, size int) image.Image {
	rect := image.Rect(0, 0, img.Bounds().Max.X/size, img.Bounds().Max.Y/size)
	dstimg := image.NewGray(rect)
	draw.ApproxBiLinear.Scale(dstimg, rect, img, img.Bounds(), draw.Over, nil)
	return dstimg
}
func newImage(r image.Rectangle, c color.Color) image.Image {
	nimg := image.NewRGBA(r)
	draw.Draw(nimg, nimg.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	return nimg
}
func drawCellinImage(imgs image.Image, imgd *image.RGBA, p image.Point) *image.RGBA {
	//nimgd := image.NewRGBA(imgd.Bounds())
	//draw.Copy(nimgd, image.ZP, imgd, nimgd.Bounds(), draw.Src, nil)
	//	draw.Copy(imgd, p, imgs, imgs.Bounds(), draw.Over, nil)
	//draw.Draw(nimgd, nimgd.Bounds(), imgd, image.ZP, draw.Src)
	//r := image.Rect(p.X+imgs.Bounds().Min.X, p.Y+imgs.Bounds().Min.Y, imgd.Bounds().Max.X, imgd.Bounds().Max.Y)
	draw.Draw(imgd, image.Rect(p.X, p.Y, imgs.Bounds().Max.X, imgs.Bounds().Max.Y), imgs, p, draw.Over)
	return imgd
}
