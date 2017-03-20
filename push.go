// +build go1.8

package gziphandler

import "net/http"

// Push initiates an HTTP/2 server push.
// Push returns ErrNotSupported if the client has disabled push or if push
// is not supported on the underlying connection.
func (w *GzipResponseWriter) Push(target string, opts *http.PushOptions) error {
	// Try to get the pusher interface.
	pusher, ok := w.ResponseWriter.(http.Pusher)
	if !ok || pusher == nil {
		return http.ErrNotSupported
	}

	if opts == nil {
		opts = &http.PushOptions{
			Header: http.Header{
				acceptEncoding: []string{"gzip"},
			},
		}
		return pusher.Push(target, opts)
	}

	if opts.Header == nil {
		opts.Header = http.Header{
			acceptEncoding: []string{"gzip"},
		}
		return pusher.Push(target, opts)
	}

	if encoding := opts.Header.Get(acceptEncoding); encoding == "" {
		opts.Header[acceptEncoding] = w.clientAcceptEncoding
		return pusher.Push(target, opts)
	}

	return pusher.Push(target, opts)
}
