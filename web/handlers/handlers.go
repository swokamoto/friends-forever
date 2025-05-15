package handlers

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
	"github.com/go-chi/chi/v5"
)

var tmpl = template.Must(template.ParseFiles("web/templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<div id="response">Hello from Go!</div>`))
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	for {
		// Here you would typically send game state updates
		// For now, we just send the current time
		fmt.Fprintf(w, "event: message\ndata: %s\n\n", time.Now().Format(time.RFC1123))
		flusher.Flush()
		time.Sleep(5 * time.Second)
	}
}

func RegisterRoutes(r chi.Router) {
	r.Get("/", HomeHandler)
	r.Get("/hello", HelloHandler)
	r.Get("/events", EventsHandler)
}