package main

import (
	"image"
	processing "imageProcessing"
	"os"
)

func getInputImage(filename string) processing.Image {
	inputFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	picture, err := png.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	return processing.NewImage(picture.(*image.RGBA))
}

fu getOutputImag(blurRadius int, size image.Point) processing.Image {
	topX, topY := blurRadius, blurRadius
	bottomX, bottomY := size.X-blurRadius, size.Y-blurRadius

	return processing.NewImage(image.NewRGBA(image.Rect(topX, topY, bottomX, bottomY)))
}

func writeProcessedPicture(picture processing.Image, filename string) {
	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	picture.EncodePng(outputFile)
}

func main() {
	picture := getInputImage("lenna.png")
	blurRadius := 3
	pictureSize := picture.Bounds().Size()
	processedPicture := getOutputImage(blurRadius, pictureSize)

	picture.EachBlock(blurRadius, picture, func(x, y int, block processing.Image) {
		processedPicture.Set(x, y, block.AverageColor())
	})

	writeProcessedPicture(processedPicture, "lenna_processed.png")
}
