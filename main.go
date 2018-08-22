package main

import (
	"log"
	"net/http"

	"github.com/nombiezinja/notifi-server/config"
	"github.com/nombiezinja/notifi-server/db"
	"github.com/nombiezinja/notifi-server/handlers"
)

func main() {

	//register static routes
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/connected/", http.StripPrefix("/connected/", http.FileServer(http.Dir("static/connected/"))))

	//register handler routes
	http.HandleFunc("/test", handlers.TestHandler)
	db.Connect()
	//log and start server
	log.Println("running server on ", config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(config.AppConfig.Port, nil))

}
