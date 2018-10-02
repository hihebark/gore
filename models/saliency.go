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
	//intensityFeatures()
	return core.RGBChannel(imgsrc, "blue")
	//return core.GaussianBlur(imgsrc, kernel, radius)
}

func intensityFeatures() {
	fmt.Printf("- Extracting intensity features from image\n")
	fmt.Printf("- Intensity of Black color %v\n", core.Intensity(core.RGBAtoRGB(color.RGBA{255, 255, 255, 255})))
	fmt.Printf("- Intensity of White color %v\n", core.Intensity(core.RGBAtoRGB(color.RGBA{0, 0, 0, 255})))
}
func colourFeature() {
	fmt.Printf("- Extracting colour feature form image\n")
}
func orientationFeatures() {
	fmt.Printf("- Extracting orientation features from image\n")
}
func rgbyCondition(rgby float64) float64 {
	if rgby < 0 {
		return .0
	}
	return rgby
}
