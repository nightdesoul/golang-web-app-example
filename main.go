package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["url"] = r.URL.Path
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	w.Write(b)
}
