package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"github.com/hihebark/gore/core"
	"github.com/hihebark/gore/log"
	"github.com/hihebark/gore/models/hog"
	"github.com/hihebark/gore/models/saliency"
)

const maxsizex int = 600

var (
	path   *string
	models *string
)

func init() {
	path = flag.String("p", "", "Path to the image.")
	models = flag.String("model", "", "Model to use hog, sal, ...")

}

func main() {
	fmt.Printf("  ┏ ┳ ┓\n  ┣ o ┫\n  ┗ ┻ ┛\n")
	fmt.Println("  Gore - 0.0.1")
	flag.Parse()
	switch {
	case *path != "" && *models != "":
		img, err := os.Open(*path)
		if err != nil {
			log.Err("image:os.Open path: %v", *path)
			defer img.Close()
		}
		info, _ := img.Stat()
		name := strings.Split(info.Name(), ".")[0]
		imgdec, format, err := image.Decode(img)
		if err != nil {
			log.Err("error while decoding image: %v", err)
		}
		i := core.NewImageInfo(format, name, imgdec.Bounds(), 2, 17)
		switch *models {
		case "hog":
			gray := i.Grayscale(imgdec)
			imghog := hog.HogVect(gray, i)
			i.Save("hog", imghog)
		case "sal":
			imgs := saliency.Salience(imgdec, 3, 1)
			for _, v := range imgs {
				i.Save(fmt.Sprintf("sal-%s", v.Name), v.Image)
			}
		case "gabor":
			log.Inf("Gabor filter: %v", core.GaborFilterKernel(7.0, 0.0, 90.0, 9, imgdec))
		}
	case *path == "":
		c := color.RGBA{255, 0, 0, 255}
		rgb := core.RGBAtoRGB(c)
		xyz := core.RGBtoXYZ(rgb)
		lab := core.XYZtoCieLAB(xyz)
		log.Inf("rgb: %v \txyz: %v \tlab: %v", rgb, xyz, lab)
		log.Inf("rgby: %v", core.RGBtoRGBY(core.RGBAtoRGB(color.RGBA{255, 255, 100, 255})))
		log.Inf("Intensity: %v", core.Intensity(core.RGBAtoRGB(c)))
		log.Inf("Gabor: %v", core.Gabor(10, 10, 3.14))
	default:
		flag.PrintDefaults()
		os.Exit(2)
	}
}
