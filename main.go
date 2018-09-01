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

var path *string

func init() {
	path = flag.String("p", "", "Path to the image.")
}

func main() {
	fmt.Printf(" ┏ ┳ ┓\n ┣ o ┫\n ┗ ┻ ┛\n")
	fmt.Println(" Gore - 0.0.0")
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
	i := core.NewImageInfo(format, name, imgdec.Bounds(), 2, 17)
	gray := i.Grayscale(imgdec)
	//imginf.saveI("SquareBox", drawsquareI(gray, image.Pt(200, 50)))
	i.Save("hog", model.HogVect(gray, i))
	i.Wg.Wait()

}
