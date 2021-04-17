package images

import (
	"embed"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// filesystem for assets.
var filesystem *embed.FS

// dirImages keeps track of which source images we
// have in each subfolder.
var dirImages map[string][]string

// dirRandSeeds maps each sub directory to an integer value.
// These integers are then used in the cacheKey function.
var dirRandSeeds map[string]int

// cacheDirPath is the path to our cached files.
var cacheDirPath string

// SetFS will set the embed filesystem for images.
func SetFS(f *embed.FS, tmpDirPath string) {
	filesystem = f
	cacheDirPath = tmpDirPath
	dirImages = map[string][]string{}
	dirRandSeeds = map[string]int{}

	// Map out which images we have in our arsenal.
	imagesDir := "assets/images"
	imagesDirEntry, err := filesystem.ReadDir(imagesDir)
	if err != nil {
		log.Fatal(err)
	}
	for i, d := range imagesDirEntry {
		if !d.IsDir() {
			continue
		}

		dirImages[d.Name()] = []string{}
		dirRandSeeds[d.Name()] = i

		subDirEntry, err := filesystem.ReadDir(fmt.Sprintf("%s/%s", imagesDir, d.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range subDirEntry {
			if f.IsDir() || !f.Type().IsRegular() {
				continue
			}

			// Image must be jpeg or png.
			imagePath := fmt.Sprintf("%s/%s/%s", imagesDir, d.Name(), f.Name())
			bb, err := filesystem.ReadFile(imagePath)
			if err != nil {
				log.Fatal(err)
			}
			ct := http.DetectContentType(bb)
			if ct != "image/jpeg" && ct != "image/png" {
				log.Fatalf(
					"%s was not valid image. Accepted types are %s and %s",
					f.Name(),
					`"image/jpeg"`,
					`"image/png"`,
				)
			}

			// Looks like we've got a valid source image.
			// Let's add that to the map.
			dirImages[d.Name()] = append(dirImages[d.Name()], imagePath)
		}

		if len(dirImages[d.Name()]) < 1 {
			log.Fatalf("image directory '%s' did not have any images", d.Name())
		}

		log.Debugf(
			"image directory %s has %d valid source images",
			d.Name(),
			len(dirImages[d.Name()]),
		)
	}
}
