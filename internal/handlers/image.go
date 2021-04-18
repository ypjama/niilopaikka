package handlers

import (
	"net/http"
	"niilopaikka/internal/images"
	"os"
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
		BadRequestWithDescription(w, r, `Liian iso.`)
		return
	}

	// TODO: resize image
	path, err := images.Image(r.Host, width, height)
	if err != nil {
		log.Error(err)
		InternalServerError(w, r)
		return
	}

	bb, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		InternalServerError(w, r)
		return
	}

	w.Header().Add("content-type", "image/jpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(bb)
}
