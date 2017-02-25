package core

import "github.com/russross/blackfriday"

func getRender() blackfriday.Renderer {
	flags := buildFlags()
	render := blackfriday.HtmlRenderer(flags, "", "")
	return render
}

func buildFlags() int {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_XHTML
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	return htmlFlags
}
