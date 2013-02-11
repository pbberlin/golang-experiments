package main

import (
	"os"
	"image"
	"image/png"
	"imageProcessing"
)

func getInputImage(filename string) *image.RGBA {
	inputFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	picture, err := png.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	return picture.(*image.RGBA)
}

func getOutputImage(blurRadius int, size image.Point) *image.RGBA {
	topX, topY := blurRadius, blurRadius
	bottomX, bottomY := size.X-blurRadius, size.Y-blurRadius

	return image.NewRGBA(image.Rect(topX, topY, bottomX, bottomY))
}

func writeProcessedPicture(picture image.Image, filename string) {
	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	png.Encode(outputFile, picture)
}

func main() {
	picture := getInputImage("lenna.png")
	blurRadius := 3
	pictureSize := picture.Bounds().Size()
	processedPicture := getOutputImage(blurRadius, pictureSize)

	imageProcessing.EachBlock(blurRadius, picture, func(x, y int, block *image.RGBA) {
		processedPicture.Set(x, y, imageProcessing.AverageColor(block))
	})

	writeProcessedPicture(processedPicture, "lenna_processed.png")
}
