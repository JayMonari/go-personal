package main

import (
	"fmt"
	"link"
	"strings"
)

const exHTML = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	ll, err := link.Parse(strings.NewReader(exHTML))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", ll)
}
