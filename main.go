package main

import (
	"fmt"
	"image"
  _ "image/jpeg"
  _ "image/png"
	"image/color"
	"os"
)

// var asciiChars = "Ñ@#W$9876543210?!abc;:+=-,._           "
var asciiChars = "@@%#*+=-:.  "
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <image path>")
		return
	}

	openFile, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Error while opening file: ", err.Error())
		return
	}

	img, _, err := image.Decode(openFile)
	if err != nil {
		fmt.Println("Error while decoding image: ", err.Error())
		return
	}

	max := img.Bounds().Max
	scaleX, scaleY := 30, 12
	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := AveragePixel(img, x, y, scaleX, scaleY)
			fmt.Print(grayToASCII(c))
		}
		fmt.Println()
	}
}

func Grayscale(color color.Color) int {
	r, g, b, _ := color.RGBA()
	return int(0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b))
}

func AveragePixel(img image.Image, x, y, w, h int) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += Grayscale(img.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

func grayToASCII(gray int) string {
	levels := len(asciiChars)
	return string(asciiChars[(gray*int(levels-1))/65536])
}
