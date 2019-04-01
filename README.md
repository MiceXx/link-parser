# link-parser

__Usage:__ 
```go get github.com/MiceXx/link-parser```

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
