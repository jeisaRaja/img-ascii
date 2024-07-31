package normal

import (
	"fmt"
	"os"

	"github.com/jeisaraja/img-ascii/pkg/ascii"
)

func Start(config *ascii.ConfigASCII) {

	openFile, err := os.Open(config.Filepath)
	if err != nil {
		fmt.Println("error while opening image, ", err.Error())
		return
	}

	ascii.ImageToASCII(openFile, config)
}
