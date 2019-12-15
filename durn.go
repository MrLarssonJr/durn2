package main

import (
	"durn2/config"
	"durn2/handler"
	"durn2/middleware"
	"durn2/model"
	"durn2/view"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

import "github.com/gorilla/mux"

func setupLog() {
	log.SetFlags(log.Ldate | log.Ltime)

	debugMode := config.Default.GetMust("DEBUG_MODE")

	if b, err := strconv.ParseBool(debugMode); err != nil {
		log.Fatalf("Config value DEBUG_MODE not a boolean (%s)", debugMode)
	} else if b {
		log.SetFlags(log.Flags() | log.Llongfile)
	}

	log.SetOutput(os.Stdout)
}

func setupDB() *gorm.DB {
	databaseUrl := config.Default.GetMust("DATABASE_URL")
	db, err := gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Error opening connection to database - %v", err)
	}

	db.AutoMigrate(
		&model.Election{},
		&model.Candidate{},
		&model.Voter{},
		&model.Vote{},
		&model.VoteEntry{},
	)

	return db
}

func createViewFactory() view.Factory {
	return view.NewFactory(
		config.Default.GetMust("DEFAULT_SITE_TITLE"),
		config.Default.GetMust("WEB_TEMPLATE_PATH"),
	)
}

func createRouter(viewFactory view.Factory, db *gorm.DB) (router *mux.Router) {
	stylePath := config.Default.GetMust("WEB_STYLE_PATH")

	router = mux.NewRouter()
	router.Use(middleware.Log.Access)

	router.PathPrefix("/css/").
		Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(stylePath))))

	helloWorld := handler.NewHelloWorld(viewFactory)
	router.Handle("/", helloWorld)

	return // named return
}

func createServer(r *mux.Router) (srv *http.Server) {
	port := config.Default.GetMust("PORT")

	srv = &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return // named return
}

func main() {
	setupLog()

	db := setupDB()
	defer db.Close()

	vf := createViewFactory()

	r := createRouter(vf, db)

	srv := createServer(r)

	log.Printf("Starting web server on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
