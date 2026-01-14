package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type api struct{
    addr string
}


func (api *api) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}


func main() {
	godotenv.Load()

    port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

    api := &api{
        addr: fmt.Sprintf(":%s", port),
    }

    mux := http.NewServeMux()

	s := &http.Server{
		Addr:           api.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	mux.HandleFunc("GET /hello", api.Hello)

	
	log.Println("Server starting on", s.Addr)

	err := http.ListenAndServe(s.Addr, nil)
	log.Fatal(err)
}
