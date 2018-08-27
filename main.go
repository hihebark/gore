package main

import (
	"flag"
	"fmt"
	"github.com/hihebark/gore/core"
)

var img *string

func init() {
	img = flag.String("img", "", "Path to the image.")
}

func main() {
	fmt.Printf(" ┏ ┳ ┓\n ┣ o ┫\n ┗ ┻ ┛\n")
	fmt.Println(" Gore - 0.0.0")
	flag.Parse()
	if *img != "" {
		core.Start(*img)
	}
}
