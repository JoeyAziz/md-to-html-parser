package main

import (
	"os"

	"github.com/joeyaziz/md-to-html-parser/markdown"
)

func main() {
	testStr := `
# for headers

## for second headers

**bold** for bold

*italic* for italic

Paragraphs for plain text (wrapped in tags)
`

	writeToFile([]byte(markdown.Parse(testStr)))
}

func writeToFile(data []byte) error {
	return os.WriteFile("./out.html", data, 0644)
}
