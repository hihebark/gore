package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"github.com/hihebark/gore/core"
	"github.com/hihebark/gore/models"
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
		defer img.Close()
		if err != nil {
			fmt.Printf("image:os.Open path:%s\n", *path)
		}

		info, _ := img.Stat()
		name := strings.Split(info.Name(), ".")[0]
		imgdec, format, err := image.Decode(img)
		if err != nil {
			fmt.Printf("error while decoding image: %v\n", err)
			panic("Decode")
		}
		i := core.NewImageInfo(format, name, imgdec.Bounds(), 2, 17)
		switch *models {
		case "hog":
			gray := i.Grayscale(imgdec)
			imghog := model.HogVect(gray, i)
			i.Save("hog", imghog)
		case "sal":
			imgsal := model.Salience(imgdec, 3, 1)
			i.Save("sal", imgsal)
		}
	case *path == "":
		rgb := core.RGBAtoRGB(color.RGBA{255, 0, 0, 128})
		xyz := core.RGBtoXYZ(rgb)
		lab := core.XYZtoCieLAB(xyz)
		fmt.Printf("rgb: %v\nxyz: %v\nlab:%v\n", rgb, xyz, lab)
	default:
		flag.PrintDefaults()
		os.Exit(2)
	}
	//i.Save("blur", model.Salience(imgdec, 3, 1))
	/*
		gray := i.Grayscale(imgdec)
		imghog := model.HogVect(gray, i)
		i.Save("hog", imghog)
	*/
}
