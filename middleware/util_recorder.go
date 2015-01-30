package middleware

import "net/http"

// responseRecorder is a type of ResponseWriter that captures
// the status code written to it and also the size of the body
// written in the response. A status code does not have
// to be written, however, in which case 200 must be assumed.
// It is best to have the constructor initialize this type
// with that default status code.
type responseRecorder struct {
	http.ResponseWriter
	status int
	size   int
}

// newResponseRecorder makes and returns a new responseRecorder,
// which captures the HTTP Status code from the ResponseWriter
// and also the length of the response body written through it.
// Because a status is not set unless WriteHeader is called
// explicitly, this constructor initializes with a status code
// of 200 to cover the default case.
func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
		status:         http.StatusOK,
	}
}

// WriteHeader records the status code and calls the
// underlying ResponseWriter's WriteHeader method.
func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// Write is a wrapper that records the size of the body
// that gets written.
func (r *responseRecorder) Write(buf []byte) (int, error) {
	n, err := r.ResponseWriter.Write(buf)
	if err == nil {
		r.size += n
	}
	return n, err
}