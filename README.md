cachely
========

Cachely execute http request, with the same interface of http package.

Example
========


```go
cachely.Expires(600 * time.MilliSecond)

res1, _ := cachely.Get("http://google.com")

res2, _ := cachely.Get("http://google.com")

res3, _ := cachely.Get("http://google.com")
```
