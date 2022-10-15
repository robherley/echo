package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(reqDump)
	})

	port := "8080"
	if eport := os.Getenv("PORT"); eport != "" {
		port = eport
	}

	http.ListenAndServe(net.JoinHostPort("0.0.0.0", port), nil)
}
