package implements

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// DefaultStore ...
type DefaultStore map[string]*record

const sep = ";"

type record struct {
	resp   *http.Response
	expire time.Time
}

var (
	store       = DefaultStore{}
	centralLock = new(sync.Mutex)
)

// Flush ...
func (s DefaultStore) Flush() error {
	now := time.Now()
	centralLock.Lock()
	for key, r := range s {
		if r.expire.Before(now) {
			delete(s, key)
		}
	}
	centralLock.Unlock()
	return nil
}

// Get ...
func (s DefaultStore) Get(req *http.Request) (*http.Response, error) {
	key, err := GenRequestUniqueKey(req)
	if err != nil {
		return nil, err
	}
	hit, ok := s[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	if hit.expire.Before(time.Now()) {
		delete(s, key)
		return nil, fmt.Errorf("already expired")
	}
	return hit.resp, nil
}

// Set ...
func (s DefaultStore) Set(req *http.Request, res *http.Response, expire time.Time) error {
	key, err := GenRequestUniqueKey(req)
	if err != nil {
		return err
	}
	// TODO: it must be expired
	s[key] = &record{
		resp:   res,
		expire: expire,
	}
	return nil
}

// GenRequestUniqueKey ...
// TODO: it must be unique also by Headers
func GenRequestUniqueKey(req *http.Request) (key string, err error) {
	body := []byte{}
	if req.Body != nil {
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return "", err
		}
	}
	return req.Method + sep + req.URL.String() + sep + string(body), nil
}
