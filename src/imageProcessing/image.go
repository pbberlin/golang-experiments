package imageProcessing

import (
	"image"
	"image/color"
)

type pixelIterator func(x int, y int, pixel color.Color)
type pixelBlockIterator func(x int, y int, block *image.RGBA)

func EachPixel(picture image.Image, action pixelIterator) {
	bounds := picture.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			action(x, y, picture.At(x, y))
		}
	}
}

func EachBlock(radius int, picture *image.RGBA, action pixelBlockIterator) {
	bounds := picture.Bounds()

	for y := bounds.Min.Y + radius; y < bounds.Max.Y-radius; y++ {
		for x := bounds.Min.X + radius; x < bounds.Max.X-radius; x++ {
			blockLimits := image.Rect(x-radius, y-radius, x+radius, y+radius)
			action(x, y, picture.SubImage(blockLimits).(*image.RGBA))
		}
	}
}

func AverageColor(picture *image.RGBA) color.Color {
	var sumR, sumG, sumB, sumA uint32
	size := picture.Bounds().Size()
	pixelCount := uint32(size.X * size.Y)

	EachPixel(picture, func(x, y int, pixel color.Color) {
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
