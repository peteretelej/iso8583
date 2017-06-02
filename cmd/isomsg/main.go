package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/peteretelej/iso8583"
)

var (
	bitmap = flag.String("bitmap", "", "convert bitmap to bits")
)

func main() {
	flag.Parse()

	switch {
	case *bitmap != "":
		out, err := iso8583.BitmapToBinary(*bitmap)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
		return
	}
}
