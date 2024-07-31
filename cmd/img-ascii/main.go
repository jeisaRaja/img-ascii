package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jeisaraja/img-ascii/pkg/interactive"
	"github.com/jeisaraja/img-ascii/pkg/normal"
)

func main() {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	_mode := f.String("i", "n", "interactive mode")

  err := f.Parse(os.Args[1:])
  if err != nil {
    fmt.Println("error in main, ", err.Error())
    return
  }
  
	switch *_mode {
	case "n":
		normal.Start(f.Args())
	case "i":
		interactive.Start(os.Stdin, os.Stdout)
	}
}
