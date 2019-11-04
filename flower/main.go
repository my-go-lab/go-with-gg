package main

import (
	"github.com/fogleman/gg"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(1.0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		// RotateAbout(angle, x, y float64)
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		// DrawEllipse(x, y, rx, ry float64)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("out.png")
}
