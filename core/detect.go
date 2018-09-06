package core

import (
	"fmt"
	"image"
	//"image/color"
)

//DetectFace
func DetectFace(imgsrc, imgmodel image.Image) image.Rectangle {
	fmt.Printf("+ Detecting for face(s)...\n")
	bmodel, bsrc := imgmodel.Bounds(), imgsrc.Bounds()
	points := make(map[image.Point]float64)
	for ys := bsrc.Min.Y; ys < bsrc.Max.Y; ys++ {
		for xs := bsrc.Min.X; xs < bsrc.Max.X; xs++ {
			equal := 0
			for ym := bmodel.Min.Y; ym < bmodel.Max.Y; ym++ {
				for xm := bmodel.Min.X; xm < bmodel.Max.X; xm++ {
					if imgsrc.At(xs, ys) == imgmodel.At(xm, ym) {
						equal++
					}
				}
			}
			points[image.Pt(xs, ys)] = float64(equal * 100 / (bmodel.Max.X * bmodel.Max.Y))
		}
	}
	return image.Rect(0, 0, 0, 0)
}
