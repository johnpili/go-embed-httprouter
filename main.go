package main

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed static/*
var staticFS embed.FS

func main() {
	f := fs.FS(staticFS)
	v, _ := fs.Sub(f, "static")

	router := httprouter.New()
	router.ServeFiles(fmt.Sprintf("%s", "/static/*filepath"), http.FS(v))
	router.HandlerFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/img/cat1.jpg", http.StatusTemporaryRedirect)
	})

	port := "8080"
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	log.Printf("Server running at http://localhost:%s/\n", port)
	log.Fatal(httpServer.ListenAndServe())
}
