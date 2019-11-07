# Invert Mask

![](./out.png)

```go
package main

import (
  "github.com/fogleman/gg"
)

func main() {

  const S = 1024

  dc := gg.NewContext(S, S)
  dc.DrawCircle(S/2, S/2, 384)
  dc.Clip()
  dc.InvertMask()
  dc.DrawRectangle(0, 0, S, S)
  dc.SetRGB(0, 0, 0)
  dc.Fill()
  dc.SavePNG("out.png")
}
```