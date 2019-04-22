package main

import (
	"io"
	"net/http"
	"fmt"
)

// Respond to call on endpoint
func respond(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Bruce from your Go sample application 03 running in Microclimate! (edited on GitHub")
	fmt.Printf("main.go: Go app endpoint response\n")
}

func main() {
	http.HandleFunc("/", respond)
	fmt.Printf("main.go: Go app listening on port 8000\n")
	http.ListenAndServe(":8000", nil)
}
