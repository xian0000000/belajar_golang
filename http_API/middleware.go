package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Data map[string]any

var data = Data{"nama": "xina", "text": ""}

func Give_data(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Autharization")

	if token != "secret" {
		http.Error(w, "GETTTT OUTTTTTT", 401)
		return
	}

	var req Data

	ip, port, _ := net.SplitHostPort(r.RemoteAddr)

	json.NewDecoder(r.Body).Decode(&req)

	for key, val := range req {
		data[key] = val
	}

	fmt.Println("pesan", data["text"])

	json.NewEncoder(w).Encode(data)
	fmt.Println(ip, port)

	fmt.Println(time.Since(start))
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, "client req ke : ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", logging(http.HandlerFunc(Give_data)))

	http.ListenAndServe(":8080", mux)
}
