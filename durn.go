package main

import (
	"durn2/config"
	"durn2/handler"
	"durn2/middleware"
	"durn2/view"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

import "github.com/gorilla/mux"

func setupLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	debugMode := config.Default.GetMust("DEBUG_MODE")

	if b, err := strconv.ParseBool(debugMode); err != nil {
		log.Fatalf("Config value DEBUG_MODE not a boolean (%s)", debugMode)
	} else if b {
		log.SetFlags(log.Flags() | log.Llongfile)
	}

	log.SetOutput(os.Stdout)
}

func createViewFactory() view.Factory {
	return view.NewFactory(
		config.Default.GetMust("DEFAULT_SITE_TITLE"),
		config.Default.GetMust("WEB_TEMPLATE_PATH"),
	)
}

func createRouter(viewFactory view.Factory) (router *mux.Router) {
	stylePath := config.Default.GetMust("WEB_STYLE_PATH")

	router = mux.NewRouter()
	router.Use(middleware.Log.Access)

	router.PathPrefix("/css/").
		Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(stylePath))))

	helloWorld := handler.NewHelloWorld()
	router.Handle("/", helloWorld)

	return // named return
}

func createServer(r *mux.Router) (srv *http.Server) {
	port := config.Default.GetMust("WEB_PORT")

	srv = &http.Server {
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return // named return
}

func main() {
	setupLog()
	vf := createViewFactory()
	r := createRouter(vf)
	srv := createServer(r)

	log.Printf("Starting web server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}