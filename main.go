package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := make(map[string]interface{})
		res["Method"] = r.Method
		res["URL"] = r.URL.EscapedPath()
		res["Proto"] = r.Proto
		res["Header"] = r.Header
		res["Body"] = r.Body
		res["ContentLength"] = r.ContentLength
		res["TransferEncoding"] = r.TransferEncoding
		res["Host"] = r.Host
		res["Form"] = r.Form
		res["PostForm"] = r.PostForm
		res["MultiPartForm"] = r.MultipartForm
		res["Trailer"] = r.Trailer
		res["RemoteAddr"] = r.RemoteAddr
		res["RequestURI"] = r.RequestURI
		res["TLS"] = r.TLS
		res["RequestedAt"] = time.Now()

		resJSON, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(resJSON))
		}
	})

	http.ListenAndServe(":8080", nil)
}
