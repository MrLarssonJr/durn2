package main

import (
	"durn2/config"
	"durn2/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

import "github.com/gorilla/mux"

func setupLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.LUTC)
	log.SetOutput(os.Stdout)
}

func createRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.Use(middleware.Log.Access)

	router.HandleFunc("/", helloWorld)

	return // named return
}

func createServer(r *mux.Router) (srv *http.Server) {
	var port string
	if val, exist := config.Default.Get("WEB_PORT"); exist {
		port = val
	} else {
		log.Fatalf("Unable to retrive SERVER_PORT from default config.")
	}

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
	r := createRouter()
	srv := createServer(r)

	log.Printf("Starting web server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func helloWorld(res http.ResponseWriter, _ *http.Request) {
	_, _ = res.Write([]byte("Hello World"))
}