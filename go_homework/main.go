package main

import (
	"fmt"
	"os"
	"strings"

	"bytes"

	"golang.org/x/net/html"
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

func visit(n *html.Node, words, pics *int) { // if same type in a row just add type at the end of last var
	//if its an element node, what tag does it have?

	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data)) //dont forget to de-reference pointer

	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
		fmt.Println("Attributes:", n.Attr[0].Key + ": " + n.Attr[0].Val)
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		visit(child, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int
	visit(doc, &words, &pics) //dont forget to pass address to pointer parameter

	return words, pics
}

func main() {

	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	fmt.Println(raw)

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)

}
