package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"time"
	"log"
)

const (
	MarkdownType = "markdown"
	PriceTitle   = "Prices List"
)

type MarkdownMsg struct {
	MsgType  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", spotHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func spotHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	mg := &MarkdownMsg{
		MsgType: MarkdownType,
		Markdown: Markdown{
			Title: PriceTitle,
			Text:  vars["region"],
		},
	}
	jsonBytes, err := json.Marshal(mg)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(jsonBytes)
}
