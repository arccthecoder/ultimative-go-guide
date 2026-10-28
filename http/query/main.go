package main

import (
	"fmt"
	"net/http"
)

// localhost:9091/default?foo=x&boo=y

func handler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("foo param:", fooParam)
	fmt.Println("boo param:", booParam)
}

func main() {
	http.HandleFunc("/default", handler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("failed to start HTTP server")
	}

}
