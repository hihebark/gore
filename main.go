package main

import (
	"fmt"
	"os"
	"github.com/hihebark/reGognition/core"
)
func main() {
	fmt.Println("reGognition - 0.0.0")
	//file, err := os.Open("data/image.jpg")
	file, err := os.Open("data/im.png")
	//file, err := os.Open("data/red.png")
	//file, err := os.Open("data/blue.png")
	//file, err := os.Open("data/green.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	//pixels, err := core.GetPixels(file)
	core.GetRGBA(file)

//	if err != nil {
//		fmt.Println("Error: Image could not be decoded")
//		os.Exit(1)
//	}
//	fmt.Printf("pixel: %v\n", pixels)
}
