package model

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hihebark/gore/core"
)

//Salience output salience image.
func Salience(imgsrc image.Image, kernel, radius int) image.Image {
	fmt.Printf("+ Calculating Salience map: \n")
	return gaussianBlur(imgsrc, kernel, radius)
}

func gaussianBlur(imgsrc image.Image, kernel, radius int) image.Image {
	bounds := imgsrc.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	imgdst := image.NewRGBA(bounds)
	rs := math.Ceil(float64(2.57) * float64(radius))
	l := maxY * maxX
	fmt.Printf("+ There is %d cells\n", l)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			//rgba := imgsrc.At(x, y).RGBA()
			var r, g, b, a uint32 = 0, 0, 0, 0
			var c float64 = 1

			for ky := -kernel; ky < kernel; ky++ {
				for kx := -kernel; kx <= kernel; kx++ {
					wght := core.Gaussian(kx, ky, rs)
					c += wght
					kr, kg, kb, ka := imgsrc.At(kx, ky).RGBA()
					r += kr * uint32(wght)
					g += kg * uint32(wght)
					b += kb * uint32(wght)
					a += ka * uint32(wght)
					//c++
					//_ := core.Gaussian(x, y, float64(5)) // to end this
				}
			}
			imgdst.Set(x, y, color.RGBA{uint8(r / uint32(c)), uint8(g / uint32(c)), uint8(b / uint32(c)), uint8(a / uint32(c))})
			fmt.Printf("- Processing with %10d cell\r", (maxX*maxY)-l)
			l--
		}
	}
	fmt.Printf("\n")
	return imgdst
}
