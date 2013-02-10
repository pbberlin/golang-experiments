package main

import (
	"fmt"
	"image"
	"image/png"
	"image/color"
	"os"
)

type pixelIterator func(x int, y int, pixel color.Color)
type pixelBlockIterator func(x int, y int, block *image.RGBA)

func eachPixel(picture image.Image, action pixelIterator) {
	bounds := picture.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			action(x, y, picture.At(x, y))
		}
	}
}

func eachBlock(radius int, picture *image.RGBA, action pixelBlockIterator) {
	bounds := picture.Bounds()

	for y := bounds.Min.Y + radius; y < bounds.Max.Y - radius; y++ {
		for x := bounds.Min.X + radius; x < bounds.Max.X - radius; x++ {
			blockLimits := image.Rect(x - radius, y - radius, x + radius, y + radius);
			action(x, y, picture.SubImage(blockLimits).(*image.RGBA))
		}
	}
}

func averageColor(picture *image.RGBA) (color.Color) {
	var sumR, sumG, sumB, sumA uint32
	size       := picture.Bounds().Size()
	pixelCount := uint32(size.X * size.Y)

	eachPixel(picture, func (x, y int, pixel color.Color) {
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
		255,
	}
}

func getInputImage(filename string) (*image.RGBA) {
	inputFile, err := os.Open(filename)
	if err != nil { panic(err) }
	defer inputFile.Close()

	picture, err := png.Decode(inputFile)
	if err != nil { panic(err) }

	return picture.(*image.RGBA)
}

func getOutputImage(blurRadius int, size image.Point) (*image.RGBA) {
	topX, topY       := blurRadius, blurRadius
	bottomX, bottomY := size.X - blurRadius, size.Y - blurRadius

	return image.NewRGBA(image.Rect(topX, topY, bottomX, bottomY))
}

func writeProcessedPicture(picture image.Image, filename string) {
	outputFile, err := os.Create(filename)
	if err != nil { panic(err) }
	defer outputFile.Close()

	png.Encode(outputFile, picture)
}

func main() {
	picture := getInputImage("lenna.png")
	blurRadius       := 3
	pictureSize      := picture.Bounds().Size()
	processedPicture := getOutputImage(blurRadius, pictureSize)

	eachBlock(blurRadius, picture, func (x, y int, block *image.RGBA) {
		processedPicture.Set(x, y, averageColor(block))
	});

	writeProcessedPicture(processedPicture, "lenna_processed.png")

	// To avoid annoying "fmt not used" error
	fmt.Println("")
}
