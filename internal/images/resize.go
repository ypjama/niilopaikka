package images

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// resize takes a source image and tries to resize it.
// The file path to the resized image is returned.
// TODO: resize and output to tmp -file.
// TODO: mv tmp -file to actual cache path.
func resize(ck cacheKey) (string, error) {
	tmpPath := ck.jpegPath(uuid.New().String() + "-")
	destPath := ck.jpegPath("")
	log.Debugf("Resizing %s and outputting it to %s", ck.sourceImage(), destPath)

	sourceFile, err := os.Open(ck.sourceImage())
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return "", err
	}
	img := image.NewNRGBA(image.Rect(0, 0, ck.width, ck.height))
	err = jpeg.Encode(tmpFile, img, &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
	tmpFile.Close()

	if err != nil {
		return "", err
	}

	err = os.Rename(tmpPath, destPath)
	if err != nil {
		return "", err
	}

	return destPath, nil
}
