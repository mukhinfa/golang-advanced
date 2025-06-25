package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	digit := rand.Intn(6) + 1
	w.Write([]byte(strconv.Itoa(digit)))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	router.HandleFunc("/", handler)
	server.ListenAndServe()
}
