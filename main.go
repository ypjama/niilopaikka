package main

import (
	"embed"
	"net/http"
	"niilopaikka/internal/handlers"
	"niilopaikka/internal/images"
	"os"
	"strconv"
	"strings"

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

	// Handlers need the HTML templates.
	handlers.ParseTemplates(&webfs)

	// Set asset file system for the images package.
	images.SetFS(&assetfs)

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
