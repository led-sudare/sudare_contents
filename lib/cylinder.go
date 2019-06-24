package lib

import (
	"image"
	"image/color"
	"sudare_contents/lib/util"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

type Cylinder struct {
	images  []*image.RGBA
	gcs     []*draw2dimg.GraphicContext
	rawData []byte
}

func NewCylinder() *Cylinder {

	c := new(Cylinder)
	c.images = make([]*image.RGBA, CylinderCount)
	c.gcs = make([]*draw2dimg.GraphicContext, CylinderCount)
	c.rawData = make([]byte, CylinderWidth*CylinderHeight*CylinderCount*2)
	util.ConcurrentEnum(0, CylinderCount, func(i int) {
		c.images[i] = image.NewRGBA(image.Rect(0, 0, CylinderWidth, CylinderHeight))
		c.gcs[i] = draw2dimg.NewGraphicContext(c.images[i])
	})
	return c
}

func (c *Cylinder) Render(draw func(int, *draw2dimg.GraphicContext)) {
	util.ConcurrentEnum(0, CylinderCount, func(i int) {
		gc := c.gcs[i]
		gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
		draw2dkit.Rectangle(gc, 0, 0, CylinderWidth, CylinderHeight)
		gc.Fill()
		draw(i, c.gcs[i])
	})
}

func (c *Cylinder) GetData() []byte {

	util.ConcurrentEnum(0, CylinderCount, func(i int) {
		for x := 0; x < CylinderWidth; x++ {
			for y := 0; y < CylinderHeight; y++ {
				idx565 := (i*CylinderHeight*CylinderWidth + CylinderWidth*y + x) * 2
				//				log.Info(idx565)
				r, g, b, _ := c.images[i].At(x, y).RGBA()
				c.rawData[idx565+0] = byte(r)&0xF8 + byte(g)>>5
				c.rawData[idx565+1] = (byte(g)<<2)&0xe0 + byte(b)>>3
			}
		}
	})
	return c.rawData
}
