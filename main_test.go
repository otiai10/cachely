package cachely

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var dummy *httptest.Server

func TestMain(m *testing.M) {

	mux := http.NewServeMux()
	mux.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("X-CACHELY-TEST", "xxx")
		// w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader(http.StatusPaymentRequired)
		fmt.Fprint(w, "bad request to foo")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

	})
	dummy = httptest.NewServer(mux)

	os.Exit(m.Run())
}
