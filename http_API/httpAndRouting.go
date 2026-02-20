// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// func hallo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "hallo xina")

// }

// func user(w http.ResponseWriter, r *http.Request) {
// 	nama := r.URL.Query().Get("nama")
// 	fmt.Fprintln(w, "hallo", nama)
// }

// type Data struct {
// 	Nama  string `json:"nama"`
// 	Class string `json:"class"`
// 	Power int    `json:"power"`
// }

// func ambil_data_json(w http.ResponseWriter, r *http.Request) {
// 	var user Data

// 	json.NewDecoder(r.Body).Decode(&user)

// 	fmt.Println(user)

// 	w.Header().Set("Content-Type", "application/json")

// 	res := map[string]any{
// 		"status": "202",
// 		"pesan1": fmt.Sprintf("hallo %s", user.Nama),
// 		"pesan2": user,
// 	}
// 	json.NewEncoder(w).Encode(res)

// }

// func main() {
// 	http.HandleFunc("/hallo/", user)
// 	http.HandleFunc("/data", ambil_data_json)
// 	http.ListenAndServe(":8000", nil)
// }
