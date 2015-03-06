package paysimple

import "net/http"

// backend is a customizable backend for the API. http.Client implements
// the backend interface by default.
type backend interface {
	Do(req *http.Request) (resp *http.Response, err error)
}
