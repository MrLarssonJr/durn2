package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

import "github.com/gorilla/mux"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.LUTC)
	log.SetOutput(os.Stdout)

	router := mux.NewRouter()

	router.HandleFunc("/", helloWorld)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func helloWorld(res http.ResponseWriter, req *http.Request) {
	log.Println(fmt.Sprintf("%s accessed hello world", req.RemoteAddr))
	_, _ = res.Write([]byte("Hello World"))
}