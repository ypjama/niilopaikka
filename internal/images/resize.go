package images

import (
	log "github.com/sirupsen/logrus"
)

// resize takes a source image and tries to resize it.
// The file path to the resized image is returned.
// TODO: resize and output to tmp -file.
// TODO: mv tmp -file to actual cache path.
func resize(ck cacheKey) (string, error) {
	tmpPath := ck.jpegPath(true)
	path := ck.jpegPath(false)

	log.Debugf("Resizing %s and outputting it to %s", ck.sourceImage(), tmpPath)

	return path, nil
}
