package ascii

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

// var asciiChars = "Ñ@#W$9876543210?!abc;:+=-,._           "
var asciiChars = "@@%#*+=-:.   "

//var asciiChars = "@&%QNW0gB#$DR8mHKAUbOGp4d9hPkqswE2]ayjz/?c*F)J7(Ltv1If{C}r;><=^,':.-` "

func ImageToASCII(file io.Reader) {
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error while decoding image: ", err.Error())
		return
	}
	max := img.Bounds().Max
	scaleX, scaleY := 4,2
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
