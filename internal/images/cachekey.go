package images

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// cacheKey identifies which files should be used.
//
// The seed in cacheKey is based on the image directory, image width,
// image height and the current time. The seed renews itself every
// hour.
type cacheKey struct {
	directory string
	width     int
	height    int
	seed      int64
}

// newCacheKey returns a cacheKey struct.
func newCacheKey(directory string, width int, height int) cacheKey {
	dirNum, ok := dirRandSeeds[directory]
	if !ok {
		log.Fatalf("invalid directory '%s' given to newCacheKey", directory)
	}

	if width > 9999 || height > 9999 {
		log.Fatal("too large width or height")
	}

	// Calculate rand seed based on multiple variables.
	t := time.Now()
	seed, err := strconv.ParseInt(
		fmt.Sprintf("%s%02d%04d%04d", t.Format("06010215"), dirNum, width, height),
		10,
		64,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("cacheKey seed: %d", seed)

	return cacheKey{
		directory: directory,
		width:     width,
		height:    height,
		seed:      seed,
	}
}

// jpegPath tells where the cached jpeg file should be.
func (ck cacheKey) jpegPath(isTmp bool) string {
	if isTmp {
		return fmt.Sprintf(
			"%s/%s%d%s",
			cacheDirPath,
			tmpFilePrefix,
			ck.seed,
			jpegExtension,
		)
	}
	return fmt.Sprintf("%s/%d%s", cacheDirPath, ck.seed, jpegExtension)
}

// sourceImage returns the path of the source image.
// The source image should change every hour for each resolution.
func (ck cacheKey) sourceImage() string {
	images, ok := dirImages[ck.directory]
	if !ok {
		log.Fatalf("no images under %s directory", ck.directory)
	}

	// With cacheKey seed we get the same result every time.
	rand.Seed(ck.seed)
	return images[rand.Intn(len(images))]
}

// cachedJpeg will return file path to cached file if it exists.
func (ck cacheKey) cachedJpeg() (string, bool) {
	path := ck.jpegPath(false)
	file, err := os.Stat(path)

	if err != nil {
		if !os.IsNotExist(err) {
			log.Error(err)
		}
		return "", false
	}
	if file.IsDir() {
		return "", false
	}
	return path, true
}
