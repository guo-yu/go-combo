package main

import "os"

var testsJs = []string{
	"./static/js/a.js",
	"./static/js/b.js",
}

var testCss = []string{
	"./static/css/a.css",
	"./static/css/b.css",
}

func main() {
	// Combo test static: js
	comboJs, err := Combo(testsJs)
	// Combo test static: css
	comboCss, err2 := Combo(testCss)
	// Combo test static: js to dist file
	err3 := ComboTo("./static/js/dist.js", testsJs)
}
