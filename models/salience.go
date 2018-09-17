package model

import (
	"fmt"
	"image"
	"image/color"
	//"math"

	"github.com/hihebark/gore/core"
)

//Salience output salience image.
func Salience(imgsrc image.Image, kernel, radius int) image.Image {
	fmt.Printf("+ Calculating Salience map: \n")
	//return blur(imgsrc, 1.0)
	return gaussianBlur(imgsrc, kernel, radius)
}
func blur(imgsrc image.Image, radius float64) image.Image {
	maxY, maxX := imgsrc.Bounds().Max.Y, imgsrc.Bounds().Max.X
	imgdst := image.NewRGBA(imgsrc.Bounds())
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			var r, g, b, a uint32 = 0, 0, 0, 0
			var count uint32 = 0

			for ky := -radius; ky < radius; ky++ {
				for kx := -radius; kx <= radius; kx++ {
					kr, kg, kb, ka := imgsrc.At(x+int(kx), y+int(ky)).RGBA()
					r += kr
					g += kg
					b += kb
					a += ka
					count++
				}
			}
			//r += ((uint32(radius)*2 + 1) * (uint32(radius)*2 + 1))
			//g += ((uint32(radius)*2 + 1) * (uint32(radius)*2 + 1))
			//b += ((uint32(radius)*2 + 1) * (uint32(radius)*2 + 1))
			//a += ((uint32(radius)*2 + 1) * (uint32(radius)*2 + 1))
			c := color.RGBA{uint8(r/count) + 1, uint8(g/count) + 1, uint8(b/count) + 1, uint8(a / count)}
			imgdst.Set(x, y, c)
		}
	}
	return imgdst
}

func gaussianBlur(imgsrc image.Image, kernel, radius int) image.Image {
	bounds := imgsrc.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	imgdst := image.NewRGBA64(bounds)
	l := maxY * maxX
	kernels := gaussianMap(kernel, float64(radius))
	fmt.Printf("+ There is %d cells\n", l)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			var r, g, b, a uint16

			for ky := -kernel; ky < kernel; ky++ {
				for kx := -kernel; kx < kernel; kx++ {
					kr, kg, kb, ka := imgsrc.At(x+kx, y+ky).RGBA()
					r += uint16(float64(kr) * kernels[kernel+kx][kernel+ky])
					g += uint16(float64(kg) * kernels[kernel+kx][kernel+ky])
					b += uint16(float64(kb) * kernels[kernel+kx][kernel+ky])
					a += uint16(float64(ka) * kernels[kernel+kx][kernel+ky])
				}
			}
			imgdst.SetRGBA64(x, y, color.RGBA64{r, g, b, a})
			fmt.Printf("- Processing with %5d cell\r", (maxX*maxY)-l)
			l--
		}
	}
	fmt.Printf("\n")
	return imgdst
}

func gaussianMap(ks int, sigma float64) [][]float64 {
	var sum float64 = 0.0
	l := ks*2 + 1
	kernel := make([][]float64, l)
	for i := 0; i < l; i++ {
		row := make([]float64, l)
		for j := 0; j < l; j++ {
			g := core.Gaussian(i, j, sigma)
			row[j] = g
			sum += g
		}
		kernel[i] = row
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			kernel[i][j] /= sum
		}
	}
	return kernel
}
