package content

import (
	"image/color"
	"math"
	"sudare_contents/lib"

	"github.com/llgcode/draw2d/draw2dimg"
)

type ContentSinLine struct {
	data  *lib.Cylinder
	count float64
}

func NewContentSinLine() CylinderContent {
	c := new(ContentSinLine)
	c.data = lib.NewCylinder()
	return c
}

func (c *ContentSinLine) GetFrame() []byte {
	c.count += 0.1
	c.data.Render(func(i int, gc *draw2dimg.GraphicContext) {

		unit := (2 * math.Pi) / lib.CylinderCount
		center := float64(lib.CylinderHeight / 2)

		wave := math.Sin(float64(i)*unit+c.count) * 20

		gc.SetStrokeColor(color.RGBA{0x00, 0xff, 0x00, 0xff})
		gc.SetLineWidth(1)

		gc.MoveTo(0, center) // should always be called first for a new path
		gc.LineTo(30, center+wave)
		gc.Close()
		gc.FillStroke()
	})
	return c.data.GetData()
}
