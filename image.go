package rice

import (
	"image"

	"github.com/qeesung/image2ascii/convert"
)

func Image2ASCIIString(img image.Image) string {
	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	// Create the image converter
	converter := convert.NewImageConverter()
	return converter.Image2ASCIIString(img, &convertOptions)
}
