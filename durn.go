package main

import (
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

	router.HandleFunc("/", helloWorld)

	return // named return
}

func createServer(r *mux.Router) (srv *http.Server) {
	srv = &http.Server {
		Handler:      r,
		Addr:         "127.0.0.1:8000",
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

func helloWorld(res http.ResponseWriter, req *http.Request) {
	log.Println(fmt.Sprintf("%s accessed hello world", req.RemoteAddr))
	_, _ = res.Write([]byte("Hello World"))
}