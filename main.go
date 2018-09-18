package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"github.com/hihebark/gore/core"
	_ "github.com/hihebark/gore/models"
)

const maxsizex int = 600

var (
	path  *string
	model *string
)

func init() {
	path = flag.String("p", "", "Path to the image.")
	model = flag.String("model", "", "Model to use hog, sal, ...")

}

func main() {
	fmt.Printf("  ┏ ┳ ┓\n  ┣ o ┫\n  ┗ ┻ ┛\n")
	fmt.Println("  Gore - 0.0.1")
	flag.Parse()
	switch {
	case *path != "" && *model != "":
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
		if imgdec.Bounds().Max.X > maxsizex {
			imgdec = i.Scale(imgdec)
		}
		gray := i.Grayscale(imgdec)
		i.Save("gray", gray)
		break
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
