package main

import (
	"log"

	"github.com/fogleman/gg"
)

func main() {
	im, err := gg.LoadImage("building.jpg")
	if err != nil {
		log.Fatal(err)
	}

	const S = 512
	dc := gg.NewContext(S, S)
	dc.DrawRoundedRectangle(0, 0, S, S, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	dc.SavePNG("out.png")
}
