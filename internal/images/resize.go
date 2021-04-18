package images

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"

	"golang.org/x/image/draw"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// tmpJpegPath will generate an available path for a temporary file.
// UUID Version 4 practically will never generate duplicates:
// https://en.wikipedia.org/wiki/Universally_unique_identifier.
func tmpJpegPath() string {
	return fmt.Sprintf(
		"%s/%s%s",
		cacheDirPath,
		uuid.New().String(),
		jpegExtension,
	)
}

// openImage takes a file path and tries to open and decode
// it as an image.
func openImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// aspectRatio of rectangle.
func aspectRatio(rect image.Rectangle) float64 {
	if rect.Max.Y < 1 || rect.Max.X < 1 {
		return float64(0)
	}
	return float64(rect.Max.X) / float64(rect.Max.Y)
}

// sourceRect calculates the source image rectangle to be used in resizing.
// The purpose of this is to retain the aspect ratio.
func sourceRect(orinalRect image.Rectangle, destAspectRatio float64) image.Rectangle {
	x0 := 0
	y0 := 0
	x1 := orinalRect.Max.X
	y1 := orinalRect.Max.Y

	if aspectRatio(orinalRect) > destAspectRatio {
		// We want wider.
		totalX := int(math.Round(float64(y1) * destAspectRatio))
		cutPerSide := (x1 - totalX) / 2
		x0 += cutPerSide
		x1 -= cutPerSide
	} else {
		// We want narrower.
		totalY := int(math.Round(float64(x1) / destAspectRatio))
		cutPerSide := (y1 - totalY) / 2
		y0 += cutPerSide
		y1 -= cutPerSide
	}

	return image.Rect(x0, y0, x1, y1)
}

// resize takes a source image and tries to resize it.
// The file path to the resized image is returned.
func resize(ck cacheKey) (string, error) {
	tmpPath := tmpJpegPath()
	destPath := ck.jpegPath()
	log.Debugf("Resizing %s and outputting it to %s", ck.sourceImage(), destPath)

	// Open the source image.
	sourceImage, err := openImage(ck.sourceImage())
	if err != nil {
		return "", err
	}

	// Create a new resized image.
	rect := image.Rect(0, 0, ck.width, ck.height)
	img := image.NewNRGBA(rect)
	draw.BiLinear.Scale(
		img,
		rect,
		sourceImage,
		sourceRect(sourceImage.Bounds(), aspectRatio(rect)),
		draw.Over,
		nil,
	)

	// Write image to temporary file.
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return "", err
	}
	err = jpeg.Encode(tmpFile, img, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
	tmpFile.Close()
	if err != nil {
		return "", err
	}

	// Move the temporary file to destination path.
	err = os.Rename(tmpPath, destPath)
	if err != nil {
		return "", err
	}

	// Return the destination path.
	return destPath, nil
}
