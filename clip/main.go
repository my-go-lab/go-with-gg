package main

/**
func (r *Context) Clip
  Clip updates the clipping region by *intersecting* the current clipping region
  with the current path as it would be filled by dc.Fill().
  The path is cleared after this operation.
*/

import (
	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(350, 500, 300)
	dc.Clip()
	dc.DrawCircle(650, 500, 300)
	dc.Clip()
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}
