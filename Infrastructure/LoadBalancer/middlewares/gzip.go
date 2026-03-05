package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

func init() {
	Register("gzip", func(cfg map[string]any) (Middleware, error) {
		return Gzip, nil
	})
}

func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		grw := gzipResponseWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(grw, r)
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
