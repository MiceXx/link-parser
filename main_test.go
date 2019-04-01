package linkparser

import (
	"reflect"
	"strings"
	"testing"
)

var html1 = `
<html>
	<body>
		<h1>abc</h1>
		<a href="/page">New link</a>
	</body>
</html>
`
var html2 = `
<html>
	<body>
		<h1>abc</h1>
		<a href="/page">New link</a>
		<a href="/page2">New link2</a>
	</body>
</html>
`
var html3 = `
<html>
	<body>
		<h1>abc</h1>
	</body>
</html>
`
var html4 = `
<html>
	<body>
		<h1>abc</h1>
		<div>
			<div>
				<a href="/page">New link</a>
			<div>
		</div>
		<div>
			<a href="/page2">New link2</a>
			<div>
				<a href="/page3">New link3</a>
			<div>
		</div>
	</body>
</html>
`

var html5 = `
<html>
	<body>
		<h1>abc</h1>
		<div>
				<a href="/page">New<span>hello</span> link</a>
		</div>
	</body>
</html>
`

var html6 = `
<html>
	<body>
		<h1>abc</h1>
		<div>
				<a href="/page">New link<!-- Comment --></a>
		</div>
	</body>
</html>
`

func TestParse(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Link
	}{
		{html1, []Link{Link{Href: "/page", Text: "New link"}}},
		{html2, []Link{Link{Href: "/page", Text: "New link"}, Link{Href: "/page2", Text: "New link2"}}},
		{html3, nil},
		{html4, []Link{Link{Href: "/page", Text: "New link"}, Link{Href: "/page2", Text: "New link2"}, Link{Href: "/page3", Text: "New link3"}}},
		{html5, []Link{Link{Href: "/page", Text: "Newhello link"}}},
		{html6, []Link{Link{Href: "/page", Text: "New link"}}},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			actual, _ := Parse(reader)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %s; expected %s", actual, tc.expected)
			}
		})
	}
}
