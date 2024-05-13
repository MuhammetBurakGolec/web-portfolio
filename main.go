package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	staticDir := "./"
	fs := http.FileServer(http.Dir(staticDir))

	http.Handle("/", fs)

	addr := ":8080"
	fmt.Printf("Started at Port%s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
