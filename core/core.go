package core

import "github.com/russross/blackfriday"

func ParseRequest(input []byte) []byte {
	extensions := buildExtensions()
	render := getRender()
	output := blackfriday.Markdown(input, render, extensions)
	return output
}

func buildExtensions() int {
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	return extensions
}
