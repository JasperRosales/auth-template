package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	user "github.com/JasperRosales/auth-template/internal/model"
)

type api struct {
	addr string
}

func (api *api) Hello(w http.ResponseWriter, r *http.Request) {
	user := user.New().SetId(1).SetEmail("j@gmail.com").SetPassword("password").Build()

	w.Write([]byte("Hello, World!\n" + fmt.Sprintf("User: %+v", user)))
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

	err := http.ListenAndServe(s.Addr, mux)
	log.Fatal(err)
}
