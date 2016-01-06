package cachely

import (
	"net/http"
	"time"

	"github.com/otiai10/cachely/implements"
)

var store Store

// Store ...
type Store interface {
	Get(*http.Request) (*http.Response, error)
	Set(*http.Request, *http.Response, time.Time) error
	Flush() error
}

func init() {
	store = implements.DefaultStore{}
}

// Flush ...
func Flush() error {
	return store.Flush()
}
