// What it does:
//
// This example uses the Window class to open an image file, and then display
// the image in a Window class.
//
// How to run:
//
// 		go run main.go in/counter.jpg

//
// +build example

package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"os"
)

type CounterRecognizer struct {
	readedImage gocv.Mat
}

func (c *CounterRecognizer) Process() {
	gray := c.gray(&c.readedImage)
	blured := c.blur(&gray)
}

func (c *CounterRecognizer) blur(src *gocv.Mat) gocv.Mat {
	var result gocv.Mat
	imagePoint := image.Point{5, 5}
	var borderType gocv.BorderType
	borderType = 0
	gocv.GaussianBlur(c.readedImage, &result, imagePoint, 0, 0, borderType)
	return result
}

func (c *CounterRecognizer) gray(image *gocv.Mat) gocv.Mat {
	var result gocv.Mat
	gocv.CvtColor(*image, &result, gocv.ColorRGBToGray)
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowimage [imgfile]")
		return
	}

	filename := os.Args[1]
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}
	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
