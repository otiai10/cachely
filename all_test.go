package cachely

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

func TestGet(t *testing.T) {

	expire = 600 * time.Millisecond

	res, err := Get(dummy.URL + "/foo")
	Expect(t, err).ToBe(nil)
	Expect(t, res.StatusCode).ToBe(http.StatusPaymentRequired)

	d1, _ := http.ParseTime(res.Header.Get("Date"))

	time.Sleep(400 * time.Millisecond)
	Flush()

	res, err = Get(dummy.URL + "/foo")
	Expect(t, err).ToBe(nil)
	Expect(t, res.StatusCode).ToBe(http.StatusPaymentRequired)

	d2, _ := http.ParseTime(res.Header.Get("Date"))

	Expect(t, d1.Unix()).ToBe(d2.Unix())

	time.Sleep(400 * time.Millisecond)
	Flush()

	res, err = Get(dummy.URL + "/foo")
	Expect(t, err).ToBe(nil)
	Expect(t, res.StatusCode).ToBe(http.StatusPaymentRequired)

	d3, _ := http.ParseTime(res.Header.Get("Date"))

	Expect(t, d1.Unix()).Not().ToBe(d3.Unix())

	Because(t, "it supplies copy of cached response", func(t *testing.T) {
		r1, _ := Get(dummy.URL + "/foo")
		b1, _ := ioutil.ReadAll(r1.Body)
		r2, _ := Get(dummy.URL + "/foo")
		b2, _ := ioutil.ReadAll(r2.Body)
		r3, _ := Get(dummy.URL + "/foo")
		b3, _ := ioutil.ReadAll(r3.Body)
		Expect(t, string(b1)).ToBe(string(b2))
		Expect(t, string(b2)).ToBe(string(b3))
		Expect(t, string(b3)).ToBe(string(b1))
	})
}
