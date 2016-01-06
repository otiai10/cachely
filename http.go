package cachely

import (
	"net/http"
	"time"
)

// Expire ...
var Expire = 10 * time.Second

// Get ...
func Get(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return new(http.Response), err
	}

	if res, err := store.Get(req); res != nil && err == nil {
		return res, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	clone, err := Clone(res)
	if err == nil {
		err = store.Set(req, clone, time.Now().Add(Expire))
	}

	return res, err
}
