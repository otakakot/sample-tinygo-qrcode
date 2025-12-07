package main

import (
	"image/color"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"
)

func main() {
	display := initdisplay.InitDisplay()

	width, height := display.Size()
	if width < 320 || height < 240 {
		display.SetRotation(ili9341.Rotation270)
	}

	enc := qrcode.NewQRCodeWriter()
	bm, _ := enc.Encode("https://x.com/otakakot", gozxing.BarcodeFormat_QR_CODE, 240, 240, nil)

	qrWidth := bm.GetWidth()
	qrHeight := bm.GetHeight()
	blockSize := 8

	displayWidth := int16(qrWidth * blockSize)
	displayHeight := int16(qrHeight * blockSize)

	offsetX := int16((320 - int(displayWidth)) / 2)
	offsetY := int16((240 - int(displayHeight)) / 2)

	white := color.RGBA{255, 255, 255, 255}
	display.FillRectangle(offsetX, offsetY, displayWidth, displayHeight, white)

	black := color.RGBA{0, 0, 0, 255}
	for y := 0; y < qrHeight; y++ {
		for x := 0; x < qrWidth; x++ {
			if bm.Get(x, y) {
				display.FillRectangle(
					offsetX+int16(x*blockSize),
					offsetY+int16(y*blockSize),
					int16(blockSize),
					int16(blockSize),
					black,
				)
			}
		}
	}
}
