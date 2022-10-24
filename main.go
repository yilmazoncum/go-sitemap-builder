package main

import (
	"fmt"
	"strings"

	link "main/linkParser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, error := link.Parse(r)
	if error != nil {
		panic(error)
	}
	fmt.Printf("%+v\n", links)
}
