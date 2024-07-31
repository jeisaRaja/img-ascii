package ascii

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func ImageToASCII(file io.Reader, config *ConfigASCII) {
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error while decoding image: ", err.Error())
		return
	}
	max := img.Bounds().Max
	scaleX, scaleY := config.scaleX, config.scaleY
	for y := 0; y < max.Y; y += scaleX {
		for x := 0; x < max.X; x += scaleY {
			c := AveragePixel(img, x, y, scaleX, scaleY, config)
			fmt.Print(grayToASCII(c, config))
		}
		fmt.Println()
	}
}

func Grayscale(color color.Color, cfg *ConfigASCII) int {
	r, g, b, _ := color.RGBA()

	r = uint32(float32(r) * cfg.brightness)
	g = uint32(float32(g) * cfg.brightness)
	b = uint32(float32(b) * cfg.brightness)

	if r > 65536 {
		r = 65536
	}
	if g > 65536 {
		g = 65536
	}
	if b > 65536 {
		b = 65536
	}

	return int(0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b))
}

func AveragePixel(img image.Image, x, y, w, h int, cfg *ConfigASCII) int {
	cnt, sum, max := 0, 0, img.Bounds().Max
	for i := x; i < x+w && i < max.X; i++ {
		for j := y; j < y+h && j < max.Y; j++ {
			sum += Grayscale(img.At(i, j), cfg)
			cnt++
		}
	}
	return sum / cnt
}

func grayToASCII(gray int, cfg *ConfigASCII) string {
	levels := len(cfg.letterSet)
	return string(cfg.letterSet[(gray*int(levels-1))/65536])
}

func clamp(x float64) uint8 {
	v := int64(x + 0.5)
	if v > 255 {
		return 255
	}
	if v > 0 {
		return uint8(v)
	}
	return 0
}
