package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/hihebark/gore/core"
	"github.com/hihebark/gore/models"
)

const maxsizex int = 600

var (
	path  *string
	scale *int
)

func init() {
	path = flag.String("p", "", "Path to the image.")
	scale = flag.Int("s", 2, "Scale image into the given s.")
}

func main() {
	fmt.Printf("  ┏ ┳ ┓\n  ┣ o ┫\n  ┗ ┻ ┛\n")
	fmt.Println("  Gore - 0.0.1")
	flag.Parse()
	if *path == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}
	img, err := os.Open(*path)
	defer img.Close()
	if err != nil {
		fmt.Printf("image:os.Open path:%s\n", path)
	}

	info, _ := img.Stat()
	name := strings.Split(info.Name(), ".")[0]
	imgdec, format, err := image.Decode(img)
	if err != nil {
		fmt.Printf("error while decoding image: %v\n", err)
		panic("Decode")
	}
	i := core.NewImageInfo(format, name, imgdec.Bounds(), *scale, 17)
	if imgdec.Bounds().Max.X > maxsizex {
		imgdec = i.Scale(imgdec)
	}
	i.Save("blur", model.Salience(imgdec, 3, 1))
	/*
		gray := i.Grayscale(imgdec)
		imghog := model.HogVect(gray, i)
		i.Save("hog", imghog)
		if err != nil {
			fmt.Printf("image:os.Open path:%s\n", path)
		}
		imgmodel, _, _ := image.Decode(imgm)
		points := core.DetectFace(imghog, imgmodel)
		fmt.Printf("Points: %v\n", points)
		c := color.RGBA{255, 255, 0, 255}
		for _, p := range points {
			imgdec = core.DrawSquare(imgdec, p.Rect, 1, c)
		+= imgsrc.At(kx, ky).RGBA()
		i.Save("Square", imgdec)
	*/
	i.Wg.Wait()

}
