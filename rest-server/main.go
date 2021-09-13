package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	pb "stats"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// EnableGZIP will attempt to compress the response if the client has passed a
// header value for Accept-Encoding which allows gzip
func EnableGZIP(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn.ServeHTTP(gzr, r)
	})
}

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(pb.CovidCases)
		if err != nil {
			log.Printf("Unable to marshal data: %v", err)
		}

		rw.Header().Set("Content-Encoding", "gzip")
		rw.Write(data)
	})

	http.Handle("/covid.json", EnableGZIP(handler))

	log.Printf("Handlers have been registered")

	http.ListenAndServe(":8080", nil)
	log.Printf("Server has started")
}
