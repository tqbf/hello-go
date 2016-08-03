package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("assets")))
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
