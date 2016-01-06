cachely
========

Cachely execute http request, with the same interface of http package.

Install
========

`go get github.com/otiai10/cachely`

Just replace `http.Get` with `cachely.Get`, that's all
```go
import "github.com/otiai10/cachely"

// res, err := http.Get(url)
res, err := cachely.Get(url)
```

Example
========

```go
cachely.Expires(200 * time.MilliSecond)

res1, _ := cachely.Get(url)
// New response

res2, _ := cachely.Get(url)
// The same response as res1

time.Sleep(400 * time.MilliSecond)

res3, _ := cachely.Get(url)
// New response, because it has been expired.

// You can flush expired records explicitly
cachely.Flush()

res4, _ := cachely.Get(url)
// The same response as res3, because it's not expired
```
