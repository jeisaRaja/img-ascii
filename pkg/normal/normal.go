package normal

import (
	"flag"
	"fmt"
	"os"

	"github.com/jeisaraja/img-ascii/pkg/ascii"
)

func Start(args []string) {
	normalFlag := flag.NewFlagSet("normal", flag.ExitOnError)

	brightness := normalFlag.Float64("b", 1.0, "set the brightness")
	letterSet := normalFlag.Int("l", 1, "choose letter set (1,2,3)")
	sx := normalFlag.Int("sx", 1, "set the scale x")
	sy := normalFlag.Int("sy", 1, "set the scale y")

	err := normalFlag.Parse(args)
	if err != nil {
		panic(err.Error())
	}

  fmt.Println(args)
	filepath := normalFlag.Arg(len(normalFlag.Args()) - 1)
	if filepath == "" {
		fmt.Println("No filepath provided!")
		return
	}

	openFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(normalFlag.Args())
		fmt.Println("Error while opening file : ", err.Error())
		return
	}

	config := ascii.DefaultConfig()

	if *brightness != 1.0 {
		config.Brightness(float32(*brightness))
	}

	if *letterSet != 1 {
		config.LetterSet(*letterSet)
	}

	if *sx != 1.0 {
		config.ScaleX(*sx)
	}

	if *sy != 1.0 {
		config.ScaleY(*sy)
	}

	ascii.ImageToASCII(openFile, config)
}
