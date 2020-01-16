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
	"image/color"
	"os"
)

const MinimumArea = 3000

type CounterRecognizer struct {
	readedImage gocv.Mat
}

func (c *CounterRecognizer) Process() *gocv.Mat {
	gray := c.gray(c.readedImage)
	defer gray.Close()

	blurred := c.blur(gray)
	defer blurred.Close()

	edged := c.canny(blurred)

	c.findContours(edged)

	return edged
}

func (c *CounterRecognizer) findContours(img *gocv.Mat) {
	contours := gocv.FindContours(*img, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	dst := &c.readedImage
	for i, c := range contours {
		area := gocv.ContourArea(c)
		if area < MinimumArea {
			continue
		}

		statusColor := color.RGBA{255, 0, 0, 0}
		gocv.DrawContours(dst, contours, i, statusColor, 2)

		rect := gocv.BoundingRect(c)
		gocv.Rectangle(dst, rect, color.RGBA{0, 0, 255, 0}, 2)
	}

}

func (c *CounterRecognizer) blur(src *gocv.Mat) *gocv.Mat {
	result := gocv.NewMat()

	imagePoint := image.Point{5, 5}
	var borderType gocv.BorderType
	borderType = 0
	gocv.GaussianBlur(c.readedImage, &result, imagePoint, 0, 0, borderType)
	return &result
}

func (c *CounterRecognizer) gray(image gocv.Mat) *gocv.Mat {
	result := gocv.NewMat()
	gocv.CvtColor(image, &result, gocv.ColorRGBToGray)
	return &result
}

func (c *CounterRecognizer) canny(image *gocv.Mat) *gocv.Mat {
	result := gocv.NewMat()
	gocv.Canny(*image, &result, 50, 200)
	return &result
}

func (c *CounterRecognizer) saveToFile(image *gocv.Mat) {
	gocv.IMWrite("out.jpg", *image)
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

	recognizer := CounterRecognizer{img}
	res := recognizer.Process()
	defer res.Close()

	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
