package cachely

import (
	"net/http"
	"time"
)

var expire = 10 * time.Second

// Expires ...
func Expires(dur time.Duration) {
	expire = dur
}

// Get ...
func Get(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return new(http.Response), err
	}

	res, err := store.Get(req)
	if res != nil && err == nil {
		return Clone(res)
	}

	res, err = http.Get(url)
	if err != nil {
		return res, err
	}

	clone, err := Clone(res)
	if err == nil {
		err = store.Set(req, clone, time.Now().Add(expire))
	}

	return res, err
}
