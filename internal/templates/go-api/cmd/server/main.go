package main

import (
	"log"
	"net/http"
)

func main() {
	port := "{{.Port}}"

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Printf("{{.ProjectName}} running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
