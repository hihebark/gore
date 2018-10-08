package model

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hihebark/gore/core"
)

//Salience output salience image.
func Salience(imgsrc image.Image, kernel, radius int) image.Image {
	fmt.Printf("+ Calculating Salience map: \n")
	return intensityFeatures(imgsrc)
	//return core.RGBChannel(imgsrc, "blue")
	//return core.GaussianBlur(imgsrc, kernel, radius)
}

func intensityFeatures(imgsrc image.Image) image.Image {
	fmt.Printf("- Extracting intensity features from image\n")
	maxX, maxY := imgsrc.Bounds().Max.X, imgsrc.Bounds().Max.Y
	imgdst := image.NewRGBA(imgsrc.Bounds())
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			r, g, b, a := imgsrc.At(x, y).RGBA()
			ri, gi, bi := core.Intensityrgb(core.RGBAtoRGB(color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}))
			imgdst.Set(x, y, color.RGBA{ri, gi, bi, uint8(a)})
		}
	}
	return imgdst
	//fmt.Printf("- Intensity of Black color %v\n", core.Intensity(core.RGBAtoRGB(color.RGBA{255, 255, 255, 255})))
	//fmt.Printf("- Intensity of White color %v\n", core.Intensity(core.RGBAtoRGB(color.RGBA{0, 0, 0, 255})))
}
func colourFeature() {
	fmt.Printf("- Extracting colour feature form image\n")
}
func orientationFeatures() {
	fmt.Printf("- Extracting orientation features from image\n")
}
