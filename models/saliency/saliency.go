package saliency

import (
	"image"
	_ "image/color"

	"github.com/hihebark/gore/core"
	"github.com/hihebark/gore/log"
)

//Salience output salience image.
func Salience(imgsrc image.Image, kernel, radius int, i *core.ImageInfo) []core.Img {
	log.Inf("Calculating Salience map:")
	imggray := i.Grayscale(imgsrc)
	return []core.Img{core.Img{core.RGBChannel(imggray, "red"), "redgray"}}
	//return core.RGBChannel(imgsrc, "blue")
	//return core.GaussianBlur(imgsrc, kernel, radius)
}

func colourFeature() {
	log.Inf("Extracting colour feature form image")
}
func orientationFeatures() {
	log.Inf("Extracting orientation features from image")
}
