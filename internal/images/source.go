package images

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// cacheKey for cached files is based on the directory, image width,
// image height and the current time. Cache key renews itself every
// hour, so for example in k8s horizontal scaling, the pods should
// use same cache keys and thus they would be using the same source
// images as the cache key also acts as a rand seed when selecting
// the source image.
func cacheKey(directory string, width int, height int) int64 {
	dirNum, ok := dirRandSeeds[directory]
	if !ok {
		log.Fatalf("invalid directory '%s' given to cacheKey", directory)
	}

	t := time.Now()
	i64, err := strconv.ParseInt(
		fmt.Sprintf("%s%02d%04d%04d", t.Format("06010215"), dirNum, width, height),
		10,
		64,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("cacheKey: %d", i64)

	return i64
}

// sourceImage will return the path of the source image.
// The source image is chosen based on the cache key.
func sourceImage(directory string, width int, height int) string {
	images, ok := dirImages[directory]
	if !ok {
		log.Fatalf("no images under %s directory", directory)
	}

	// Use cachekey as the rand seed.
	rand.Seed(cacheKey(directory, width, height))
	return images[rand.Intn(len(images))]
}

// hostToDirectory will map a host to image sub directory.
func hostToDirectory(host string) string {
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

func Debug(host string, width int, height int) string {
	return sourceImage(hostToDirectory(host), width, height)
}
