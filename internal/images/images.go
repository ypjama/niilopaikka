package images

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	jpegExtension = ".jpg"
)

// HostToDirectory will map a host to image sub directory.
func HostToDirectory(host string) string {
	fallbackIsValid := false
	fallback := "niilopaikka"

	for n := range dirRandSeeds {
		if n == fallback {
			fallbackIsValid = true
		}

		if strings.Contains(strings.ToLower(host), n) {
			return n
		}
	}

	if fallbackIsValid {
		return fallback
	}

	log.Fatal("could not map host to image directory")
	return ""
}

// Image returns a file path to a resized image.
func Image(host string, width int, height int) (string, error) {
	directory := HostToDirectory(host)
	ck := newCacheKey(directory, width, height)

	// Check if image is cached.
	path, ok := ck.cachedJpeg()
	if ok {
		log.Debugf("cached image %s found", path)
		return path, nil
	}

	// Resize the source image and return the output path.
	path, err := resize(ck)
	if err != nil {
		return "", err
	}

	return path, nil
}
