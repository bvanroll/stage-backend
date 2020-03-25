package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/err/500", errorInternal)
	http.HandleFunc("/err/timeout", errorTimeout)
	http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"serverName\":\"backend-server\",\"success:false\",\"version\": \"experimental\"}")
}

func errorInternal(w http.ResponseWriter, r *http.Request) {
	panic(500)
}

func errorTimeout(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "this was delayed by 2 seconds")
}