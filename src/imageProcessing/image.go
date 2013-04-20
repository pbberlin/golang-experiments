package imageProcessing

import (
	"image"
	"image/png"
	"image/color"
)

type Image struct {
	bitmap *image.RGBA
}

type pixelIterator func(x int, y int, pixel color.Color)
type pixelBlockIterator func(x int, y int, block Image)

func NewImage(bitmap *image.RGBA) Image {
	return Image{bitmap}
}

func (self Image) Set(x, y int, pixel color.Color) {
	self.bitmap.Set(x, y, pixel)
}

func (self Image) EachPixel(action pixelIterator) {
	bounds := self.bitmap.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			action(x, y, self.bitmap.At(x, y))
		}
	}
}

func (self Image) EachBlock(radius int, action pixelBlockIterator) {
	bounds := self.bitmap.Bounds()

	for y := bounds.Min.Y + radius; y < bounds.Max.Y-radius; y++ {
		for x := bounds.Min.X + radius; x < bounds.Max.X-radius; x++ {
			blockLimits := image.Rect(x-radius, y-radius, x+radius, y+radius)
			action(x, y, Image{self.bitmap.SubImage(blockLimits).(*image.RGBA)})
		}
	}
}

func (self Image) AverageColor() color.Color {
	var sumR, sumG, sumB, sumA uint32
	size := self.bitmap.Bounds().Size()
	pixelCount := uint32(size.X * size.Y)

	self.EachPixel(func(x, y int, pixel color.Color) {
		r, g, b, a := pixel.RGBA()
		sumR += (r / 0x101)
		sumG += (g / 0x101)
		sumB += (b / 0x101)
		sumA += (a / 0x101)
	})

	return color.RGBA{
		uint8(sumR / pixelCount),
		uint8(sumG / pixelCount),
		uint8(sumB / pixelCount),
		uint8(sumA / pixelCount),
	}
}

func (self Image) EncodePng(file someName) {
	png.Encode(file, self)
}
