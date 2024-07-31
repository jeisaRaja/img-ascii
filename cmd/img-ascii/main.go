package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jeisaraja/img-ascii/pkg/ascii"
	"github.com/jeisaraja/img-ascii/pkg/interactive"
	"github.com/jeisaraja/img-ascii/pkg/normal"
)

func main() {

	config := ascii.DefaultConfig()

	_interactive := flag.Bool("i", true, "interactive mode")
	_filepath := flag.String("f", "", "provide filepath to the image")
	_letterSet := flag.Int("l", 1, "select letter set")
	_brightness := flag.Float64("b", 1.0, "set brightness")
	_scaleX := flag.Int("sx", 20, "set scale x")
	_scaleY := flag.Int("sy", 10, "set scale y")

	flag.Parse()

	config.LetterSet(*_letterSet)
	config.Brightness(float32(*_brightness))
	config.ScaleX(*_scaleX)
	config.ScaleY(*_scaleY)
	config.Filepath = *_filepath

	switch *_interactive {
	case true:
		if *_filepath == "" {
			fmt.Println("Please provide filepath using -f")
			return
		}
		normal.Start(config)
	case false:
		interactive.Start(os.Stdin, os.Stdout)
	}
}
