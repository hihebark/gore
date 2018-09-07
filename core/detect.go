package core

import (
	"fmt"
	"image"
	//"image/color"
	"sort"
)

type points struct {
	percent float64
	point   image.Point
}
type bypercent []points

func (b bypercent) Len() int           { return len(b) }
func (b bypercent) Less(i, j int) bool { return b[i].percent > b[j].percent }
func (b bypercent) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

//DetectFace if found.
func DetectFace(imgsrc, imgmodel image.Image) []points {
	fmt.Printf("+ Detecting for face(s)...\n")
	bmodel, bsrc := imgmodel.Bounds(), imgsrc.Bounds()
	pts := []points{} //make(map[image.Point]float64)
	for ys := bsrc.Min.Y; ys < bsrc.Max.Y; ys++ {
		for xs := bsrc.Min.X; xs < bsrc.Max.X; xs++ {
			equal := 0
			for ym := bmodel.Min.Y; ym < bmodel.Max.Y; ym++ {
				for xm := bmodel.Min.X; xm < bmodel.Max.X; xm++ {
					if imgsrc.At(xs+xm, ys+ym) == imgmodel.At(xm, ym) {
						equal++
					}
				}
			}
			pts = append(pts, points{float64(equal * 100 / (bmodel.Max.X * bmodel.Max.Y)), image.Pt(xs, ys)})
			//pts[image.Pt(xs, ys)] = float64(equal * 100 / (bmodel.Max.X * bmodel.Max.Y))
		}
	}
	sort.Sort(bypercent(pts))
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].percent > pts[j].percent
	})
	return pts[:10]
}
