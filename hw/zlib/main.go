package main

import (
	"compress/flate"
	"compress/zlib"
	"io"
	"net/http"
	"strings"
)

type zlibWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w zlibWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за zlib-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>")
}

func deflateHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает deflate-сжатие
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			next.ServeHTTP(w, r)
			return
		}
		flatew, err := zlib.NewWriterLevel(w, flate.BestCompression)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer flatew.Close()

		w.Header().Set("Content-Encoding", "deflate")
		next.ServeHTTP(zlibWriter{ResponseWriter: w, Writer: flatew}, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandle)
	http.ListenAndServe(":3000", deflateHandle(mux))
}
