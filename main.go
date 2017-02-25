package main

import (
	"fmt"

	"github.com/russross/blackfriday"
)

func main() {
	input := "# Header\ncontent\n#HEader\n*a*"

	inputInBytes := []byte(input)
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	render := blackfriday.HtmlRenderer(htmlFlags, "", "")
	output := blackfriday.Markdown(inputInBytes, render, extensions)

	fmt.Println(string(output))
}
