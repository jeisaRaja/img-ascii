package normal

import (
	"fmt"
	"os"

	"github.com/jeisaraja/img-ascii/pkg/ascii"
)

func NormalMode() {
	openFile, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Error while opening file: ", err.Error())
		return
	}

	ascii.ImageToASCII(openFile)
}
