# linkparser

__Usage:__ 
```go get github.com/MiceXx/linkparser```

```
parser.Parse(r io.Reader)
```

Returns a slice of Links 
```
type Link struct {
	Href string
	Text string
}
```
