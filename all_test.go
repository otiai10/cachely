package cachely

import (
	"net/http"
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

func TestGet(t *testing.T) {

	Expire = 600 * time.Millisecond

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
}
