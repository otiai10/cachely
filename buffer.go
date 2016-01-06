package cachely

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Buffer is a dummy response.Body
type Buffer struct {
	bytes.Buffer
}

// Close ...
func (buf *Buffer) Close() error {
	return nil
}

// Clone ...
func Clone(original *http.Response) (*http.Response, error) {

	clone := *original

	b, err := ioutil.ReadAll(original.Body)
	if err != nil {
		return nil, err
	}

	original.Body = &Buffer{*bytes.NewBuffer(b)}
	clone.Body = &Buffer{*bytes.NewBuffer(b)}

	return &clone, nil
}
