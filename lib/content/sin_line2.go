package content

import (
	"image/color"
	"math"
	"sudare_contents/lib"

	"github.com/llgcode/draw2d/draw2dimg"
)

type ContentSinLine2 struct {
	data  *lib.Cylinder
	count float64
}

func NewContentSinLine2() CylinderContent {
	c := new(ContentSinLine2)
	c.data = lib.NewCylinder()
	return c
}

func (c *ContentSinLine2) GetFrame() []byte {
	c.count += 0.1
	c.data.Render(func(i int, gc *draw2dimg.GraphicContext) {

		unit := (2 * math.Pi) / lib.CylinderCount
		center := float64(lib.CylinderHeight / 2)
		sin := math.Sin((float64(i)*unit + c.count) * 1.2)
		sin2 := math.Sin(-(float64(i)*unit + c.count) * 2)

		wave := sin * 20

		gc.SetStrokeColor(color.RGBA{0xcc, 0x00, 0xff, 0xff})
		gc.SetLineWidth(1)

		gc.MoveTo(0, center) // should always be called first for a new path
		gc.LineTo(15*math.Abs(sin2)+7, center+wave)
		gc.Close()
		gc.FillStroke()
	})
	return c.data.GetData()
}
