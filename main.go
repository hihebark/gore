package main

import (
	"fmt"
	"os"
	"github.com/hihebark/gore/core"
)

//https://medium.com/@ageitgey/machine-learning-is-fun-part-4-modern-face-recognition-with-deep-learning-c3cffc121d78

func main() {
	fmt.Println("Gore - 0.0.0")
	//file, err := os.Open("data/image.jpg")
	//file, err := os.Open("data/im.png")
	//file, err := os.Open("data/red.png")
	//file, err := os.Open("data/blue.png")
	//file, err := os.Open("data/green.png")
	file, err := os.Open("data/RMS.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	core.MakeItGray(file)
//	pixels, err := core.GetPixels(file)
//	//core.GetRGBA(file)

//	if err != nil {
//		fmt.Println("Error: Image could not be decoded")
//		os.Exit(1)
//	}
//	fmt.Printf("pixel: %v\n", pixels)
}
