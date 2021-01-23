package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func test2(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	_, _ = fmt.Fprintf(w, "ok")
}

func main() {
	http.HandleFunc("/test", test2)
	err := http.ListenAndServe(":9069", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}