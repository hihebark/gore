package saliency

import (
	"image"
	"image/color"

	"github.com/hihebark/gore/core"
	"github.com/hihebark/gore/log"
)

//Salience output salience image.
func Salience(imgsrc image.Image, kernel, radius int, i *core.ImageInfo) []core.Img {
	log.Inf("Calculating Salience map:")
	//_ := i.Grayscale(imgsrc)
	return []core.Img{core.Img{intensityFeatures(imgsrc), "intensity"}}
	//return core.RGBChannel(imgsrc, "blue")
	//return core.GaussianBlur(imgsrc, kernel, radius)
}

func intensityFeatures(imgsrc image.Image) image.Image {
	log.Inf("Extracting intensity features from image")
	maxX, maxY := imgsrc.Bounds().Max.X, imgsrc.Bounds().Max.Y
	imgdst := image.NewGray(imgsrc.Bounds())
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			r, g, b, _ := imgsrc.At(x, y).RGBA()
			imgdst.SetGray(x, y, color.Gray{uint8((r + g + b) / 3)})
		}
	}
	return imgdst
	//fmt.Printf("- Intensity of Black color %v\n", core.Intensity(core.RGBAtoRGB(color.RGBA{255, 255, 255, 255})))
	//fmt.Printf("- Intensity of White color %v\n", core.Intensity(core.RGBAtoRGB(color.RGBA{0, 0, 0, 255})))
}
func colourFeature() {
	log.Inf("Extracting colour feature form image")
}
func orientationFeatures() {
	log.Inf("Extracting orientation features from image")
}
