package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"niilopaikka/internal/handlers"
	"niilopaikka/internal/images"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

//go:embed web/*
var webfs embed.FS

//go:embed assets/*
var assetfs embed.FS

// TODO: ads (adsense).
// TODO: google analytics.
func main() {
	// Logger.
	log.SetOutput(os.Stdout)
	logFormatter, ok := os.LookupEnv("LOG_FORMATTER")
	if ok && strings.ToUpper(logFormatter) == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	ll, _ := os.LookupEnv("LOG_LEVEL")
	if logLevel, err := log.ParseLevel(ll); err == nil {
		log.SetLevel(logLevel)
	}

	// Cache directory.
	cacheDirPath, err := ioutil.TempDir("", "niilopiilo")
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("cache directory: %s", cacheDirPath)
	defer os.RemoveAll(cacheDirPath)

	// Start collecting garbage in the background.
	go garbage(cacheDirPath)

	// Handlers need the HTML templates.
	handlers.ParseTemplates(&webfs)

	// Set file system and cache directory for the images package.
	images.SetFS(&assetfs, cacheDirPath)

	// Port.
	var defaultPort = "8080"
	var port string
	if port, ok = os.LookupEnv("PORT"); !ok {
		port = defaultPort
	}
	portInt, err := strconv.Atoi(port)
	if err != nil || portInt < 1 {
		port = defaultPort
	}

	// Routes.
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	r.HandleFunc("/{width:[0-9]+}/{height:[0-9]+}", handlers.ImageHandler).Methods("GET")
	log.Debugf("Listening port %s", port)
	log.Fatal(
		http.ListenAndServe(":"+port, r),
	)
}

func garbage(dir string) {
	log.Debugf("garbage dir: %s", dir)

	for range time.Tick(time.Minute * 5) {
		log.Debugf("Collecting garbage from %s", dir)

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		// Threshold is one hour.
		threshold := time.Now().Add(-time.Hour * 1).Unix()
		for _, file := range files {
			if file.IsDir() || threshold < file.ModTime().Unix() {
				continue
			}

			// Too old, remove it!
			path := fmt.Sprintf("%s/%s", dir, file.Name())
			log.Debugf("Removing file %s", path)
			os.Remove(path)
		}
	}
}
