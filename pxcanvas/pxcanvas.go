package pxcanvas

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PXCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows*pxCanvas.PxSize + int(pxCanvas.CanvasOffset.y))
	return image.Rect(x0, y0, x1, y1)
}
