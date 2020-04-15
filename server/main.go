package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "success to connect a server")
}

func main() {
	muxHTTP := http.NewServeMux()
	muxHTTP.HandleFunc("/", handler)

	ServerHTTP := &http.Server{
		Addr:           ":8080",
		Handler:        muxHTTP,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := ServerHTTP.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
