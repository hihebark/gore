package img

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hihebark/gore/log"
)

//Image image information.
type Image struct {
	sync.RWMutex
	Interface
	Sync   sync.WaitGroup
	Ext    string
	Name   string
	Img    image.Image
	Bounds image.Rectangle
}

// Interface
type Interface interface {
	Save(f string) error
	Scale() image.Image
	Grayscale() (image.Image, error)
	Read() ([]string, error)
}

//NewImageInfo return ImageInfo struct.
func NewImage(name string, bounds image.Rectangle) *Image {
	file := strings.Split(name, filepath.Ext(name))
	return &Image{
		Sync:   sync.WaitGroup{},
		Ext:    file[1],
		Name:   file[0],
		Bounds: bounds,
	}
}

// Grayscale make an image gray.
func (i *Image) Grayscale() image.Image {
	log.Inf("Grascaling %s.%s ...", i.Img, i.Ext)
	if i.img.ColorModel() == color.GrayModel {
		return i.Img
	}
	w, h := i.Bounds.Max.X, i.Bounds.Max.Y
	gray := image.NewGray(i.Img)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, color.GrayModel.Convert(i.Img.At(x, y)))
		}
	}
	return gray
}

// Save save image to output file.
func (i *Image) Save(folder string) error {
	outputName := fmt.Sprintf("%s/%s.gore.%s", folder, i.Name, i.Ext)
	out, err := os.Create(outputName)
	if err != nil {
		log.Err("Error While creating image %v...", err)
		defer out.Close()
		return err
	}
	switch i.Ext {
	case "png":
		png.Encode(out, i.Img)
	case "jpg", "jpeg":
		jpeg.Encode(out, i.Img, nil)
	}
	log.Inf("Saving %s", outputName)
	return nil
}

// Scale re-scale image t given size.
func (i *Image) Scale(size int) image.Image {
	log.Inf("+ Scale image into %d", size)
	bounds := i.Bounds
	rect := image.Rect(0, 0, int(bounds.Max.X/size), int(bounds.Max.Y/size))
	dstimg := image.NewRGBA(rect)
	draw.ApproxBiLinear.Scale(dstimg, rect, i.Img, i.Bounds(), draw.Over, nil)
	return dstimg
}

// Read pixel by pixel and return RGBA array.
func (i *Image) Read() color.RGBA {
	maxX, maxY := i.Bounds.Max.X, i.Bounds.Max.Y
	log.Inf("Reading %s with %d*%d ...", i.Name, maxX, maxY)
	pix := make([][]color.RGBA, maxX*maxY)
	for y := 0; y <= maxY; y++ {
		row := make([]color.RGBA, maxY)
		for x := 0; x <= maxX; x++ {
			row[i] = i.Img.At(x, y).RGBA()
		}
		pix[y] = row
	}
	return pix
}
