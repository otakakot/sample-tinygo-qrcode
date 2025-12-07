package main

import (
	"image/color"

	"github.com/skip2/go-qrcode"
	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"
)

func main() {
	display := initdisplay.InitDisplay()

	width, height := display.Size()
	if width < 320 || height < 240 {
		display.SetRotation(ili9341.Rotation270)
	}

	qr, err := qrcode.New("https://x.com/otakakot", qrcode.Low)
	if err != nil {
		panic(err)
	}

	bm := qr.Bitmap()

	qrSize := len(bm)
	blockSize := 8

	displayWidth := int16(qrSize * blockSize)
	displayHeight := int16(qrSize * blockSize)

	offsetX := int16((320 - int(displayWidth)) / 2)
	offsetY := int16((240 - int(displayHeight)) / 2)

	white := color.RGBA{255, 255, 255, 255}
	display.FillRectangle(offsetX, offsetY, displayWidth, displayHeight, white)

	black := color.RGBA{0, 0, 0, 255}
	for y := 0; y < qrSize; y++ {
		for x := 0; x < qrSize; x++ {
			if bm[y][x] {
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
