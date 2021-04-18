package handlers

import (
	"fmt"
	"net/http"
	"niilopaikka/internal/images"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// ImageHandler serves a placeholder image with the requested dimensions.
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.WithFields(log.Fields{
		"vars": vars,
	}).Debug("ImageHandler")

	// Parse width and height.
	widthStr, ok := vars["width"]
	if !ok {
		BadRequest(w, r)
		return
	}
	heightStr, ok := vars["height"]
	if !ok {
		BadRequest(w, r)
		return
	}
	var width, height int
	var err error
	if width, err = strconv.Atoi(widthStr); err != nil {
		BadRequest(w, r)
		return
	}
	if height, err = strconv.Atoi(heightStr); err != nil {
		BadRequest(w, r)
		return
	}
	if width < 1 || width > 3500 || height < 1 || height > 3500 {
		BadRequest(w, r)
		return
	}

	// TODO: serve generated image if it exits and is not too old.
	// TODO: resize image
	// TODO: save resized image to generated folder
	// TODO: serve resized image

	path, err := images.Image(r.Host, width, height)
	if err != nil {
		log.Error(err)
		InternalServerError(w, r)
		return
	}

	fmt.Fprint(w, "TODO: "+path)
}
