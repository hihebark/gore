package main

import (
	"flag"
	"fmt"
	"github.com/hihebark/gore/core"
)

/*******************************************************************************
 For more check the link --v
 https://medium.com/@ageitgey/machine-learning-is-fun-part-4-modern-face-recognition-with-deep-learning-c3cffc121d78
*******************************************************************************/
var img *string

func init() {
	img = flag.String("img", "", "Image")
}

func main() {
	fmt.Printf(" ┏ ┳ ┓\n ┣ o ┫\n ┗ ┻ ┛\n")
	fmt.Println(" Gore - 0.0.0")
	if *img != "" {
		core.Start("data/harold.png")
		//core.Start("data/RMS.png")
	}
}
