package status_writer

import "net/http"

// StatusWriter type adds a `Status` property to Go's `http.ResponseWriter` type.
type StatusWriter struct {
	http.ResponseWriter
	Status int
}

// New shows setup for a middleware function. This could be a logging
// function or similar whereby you require access to the status code from
// `http.ResponseWriter`
//
//			func MyMiddleware(h http.HandleFunc) http.HandlerFunc {
//					return func(w http.ResponseWriter, r *http.Request) {
//	       			sw := status_writer.New(w) // Here we override http.ResponseWriter's `WriteHeader` function
//						h(sw, r)
//	       			statusCode := sw.Status // Now work with the status code
//	   			}
//			}
func New(w http.ResponseWriter) *StatusWriter {
	return &StatusWriter{
		ResponseWriter: w,
		Status:         http.StatusOK,
	}
}

// WriteHeader override the http.ResponseWriter's `WriteHeader` method
func (w *StatusWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}
