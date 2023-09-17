package main

import (
	"embed"
	"flag"
	"fmt"
	"hmcalister/GolangWebAppTemplate/backend/api"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Path to embedded frontend. Made a constant so changes to embedded path and the go:embed directive are all in the same place!
const EMBED_PATH = "frontend/build"

//go:embed all:frontend/build
var content embed.FS

var (
	// Flag to indicate development mode, rather than build mode.
	developmentFlag *bool

	// Flag to allow CORS requests, typically only turned on during development
	allowCORSFlag *bool

	// Port number to serve on
	// port *int

	// The frontend filesystem, after processing away any subdirectory shenanigans.
	frontendFileSystem fs.FS
)

func init() {
	developmentFlag = flag.Bool("developmentServer", false, "Set program to run in development mode. Adds verbosity, debug statements, logging, etc.")
	allowCORSFlag = flag.Bool("allowCORS", false, "Allows CORS requests to server. WARNING: Unsafe for use in production, only use in development!")
	flag.Parse()

	// Set release mode if dev flag not set (default)
	if !*developmentFlag {
		gin.SetMode(gin.ReleaseMode)
	}

	// Get the frontend filesystem and strip out the subdirectory paths
	rootFS := fs.FS(content)
	frontendFileSystem, _ = fs.Sub(rootFS, EMBED_PATH)
}

func main() {
	// GIN Router and Middleware ----------------------------------------------

	router := gin.Default()
	if *allowCORSFlag {
		router.Use(cors.Default())
	}

	// Routes------------------------------------------------------------------

	router.GET("/api/ping", api.PongResponse)

	// Routing the static files------------------------------------------------
	// Do not touch unless you know what you are doing!

	// To avoid conflicting routes in Gin we walk over the embedded filesystem and statically serve all files
	frontendHTTPFileSystem := http.FS(frontendFileSystem)
	fs.WalkDir(frontendFileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			router.StaticFileFS(path, path, frontendHTTPFileSystem)
		}
		return nil
	})
	// Gin will happily serve "index.html" on "/" but immediately redirects this to "/" again. On Loop. Forever.
	// So we instead serve "./" from the HTTP file system
	router.StaticFileFS("/", "./", frontendHTTPFileSystem)

	// ------------------------------------------------------------------------
	// Print serving address if in build mode, so enduser can simply click on link
	fmt.Println("Serving on http://localhost:8080/")
	router.Run(":8080")
}
