package main

import (
	"log"
	"net/http"
	"os"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/int128/jira-to-slack/handlers"
)

func router() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", &handlers.Index{}).Methods("GET")
	r.Handle("/", gh.ContentTypeHandler(&handlers.Webhook{}, "application/json")).Methods("POST")

	m := http.NewServeMux()
	m.Handle("/", gh.LoggingHandler(os.Stdout, r))
	m.Handle("/healthz", &handlers.Healthz{})
	return m
}

func main() {
	addr := ":3000"
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, router()); err != nil {
		log.Fatalf("Error while listening on %s: %s", addr, err)
	}
}
