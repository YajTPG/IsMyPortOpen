package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func StartServer(port int, host string) {
	fmt.Printf("Starting server on http://%s:%d\n", host, port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(time.Now().String(), " ", r.Host, "\n")
		fmt.Fprintf(w, "Looks like port %d is open!", port)

	})
	http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
}
