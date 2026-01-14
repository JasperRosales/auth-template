package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}

func main() {
    godotenv.Load()
    
    http.HandleFunc("/", Hello)

    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT environment variable not set") 
    }

    addr := fmt.Sprintf(":%s", port)
    log.Println("Server starting on", addr)

    err := http.ListenAndServe(addr, nil)
    log.Fatal(err) 
}
