package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Add the request string
		url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
		fmt.Fprintf(w, "Request: %s\n", url)

		// Add the host
		fmt.Fprintf(w, "Host: %s\n", r.Host)

		fmt.Fprintf(w, "\n# EnvVars:\n")
		var keys []string
		for _, pair := range os.Environ() {
			keys = append(keys, pair)
		}
		sort.Strings(keys)
		for _, pair := range keys {
			fmt.Fprintf(w, "%s\n", pair)
		}

		fmt.Fprintf(w, "\n# Headers:\n")
		// Loop through headers
		var headers []string
		for key, _ := range r.Header {
			headers = append(headers, key)
		}
		sort.Strings(headers)
		for _, key := range headers {
			for _, h := range r.Header[key] {
				name := strings.ToLower(key)
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}

		// If this is a POST, add post data
		if r.Method == "POST" {
			r.ParseForm()
			fmt.Fprintf(w, "\n# Post:\n")
			fmt.Fprintf(w, r.Form.Encode())
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
