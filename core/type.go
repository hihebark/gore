package core

import (
	"image"
	"sync"
)

//Img
type Img struct {
	Image image.Image
	Name  string
}

//ImageInfo image information.
type ImageInfo struct {
	Wg sync.WaitGroup
	sync.RWMutex
	Format   string
	Name     string
	Bounds   image.Rectangle
	Scalsize int
	Cellsize int
}

//NewImageInfo return ImageInfo struct.
func NewImageInfo(f, n string, b image.Rectangle, s, c int) *ImageInfo {
	return &ImageInfo{
		Wg:       sync.WaitGroup{},
		Format:   f,
		Name:     n,
		Bounds:   b,
		Scalsize: s,
		Cellsize: c,
	}
}
