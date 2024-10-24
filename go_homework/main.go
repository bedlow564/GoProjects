package main

import (
	"fmt"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My first Heading</h1>
	  <p>My first paragraph</p>
	  <p>HTML images are defined with the img tag:</p>
	  <img src="xxx.jpg" width="104" height="142">
  </body>
</html>
`

func main() {
	fmt.Println(raw)
}