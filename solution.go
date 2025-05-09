package main

import (
	"os"
	"strings"
)

type HTMLTag struct {
	OpenTag  string
	CloseTag string
}

const _breakTag = "<br />"

var _paragraphTag = HTMLTag{
	OpenTag:  "<p>",
	CloseTag: "</p>",
}

var mdHeadersMap = map[string]HTMLTag{
	"#": {
		OpenTag:  "<h1>",
		CloseTag: "</h1>",
	},
	"##": {
		OpenTag:  "<h2>",
		CloseTag: "</h2>",
	},
}

var mdParagraphsMap = map[string]HTMLTag{
	"*": {
		OpenTag:  "<em>",
		CloseTag: "</em>",
	},
	"**": {
		OpenTag:  "<strong>",
		CloseTag: "</strong>",
	},
}

func main() {
	testStr := `
# for headers

## for second headers

**bold** for bold

*italic* for italic

Paragraphs for plain text (wrapped in tags)
`

	writeToFile([]byte(parse(testStr)))
}

func parse(str string) string {
	var out strings.Builder
	lines := strings.Split(strings.TrimSpace(str), "\n")

	for _, line := range lines {
		out.WriteString(parseLine(line).String())
	}
	return out.String()
}

func parseLine(line string) (out *strings.Builder) {
	out = &strings.Builder{}

	if len(line) == 0 {
		return out
	}

	words := strings.Split(line, " ")
	var headerTag *HTMLTag
	for i, word := range words {
		if i == 0 {
			headerTag = parseHeaders(word)
			if headerTag != nil {
				out.WriteString(headerTag.OpenTag)
				continue
			} else {
				out.WriteString(_paragraphTag.OpenTag)
			}
		}
		out.WriteString(parseWord(word).String())
		out.WriteByte(byte(' '))
	}

	if headerTag != nil {
		out.WriteString(headerTag.CloseTag)
	} else {
		out.WriteString(_paragraphTag.CloseTag)
	}
	headerTag = nil
	out.WriteString(_breakTag)
	return out
}

func parseHeaders(word string) *HTMLTag {
	if tag, ok := mdHeadersMap[word]; ok {
		return &tag
	}
	return nil
}

func parseWord(word string) (out *strings.Builder) {
	out = &strings.Builder{}
	inStrong := false
	inEm := false

	for i := 0; i < len(word); i++ {
		if word[i] == '*' {
			if i+1 < len(word) && word[i+1] == '*' { // strong
				parseStrong(out, inStrong)
				inStrong = !inStrong
				i++
				continue
			} else { // em
				parseEm(out, inEm)
				inEm = !inEm
				continue
			}
		}
		out.WriteByte(word[i])
	}
	return out
}

func parseStrong(str *strings.Builder, inStrong bool) {
	if inStrong {
		str.WriteString(mdParagraphsMap["**"].CloseTag)
	} else {
		str.WriteString(mdParagraphsMap["**"].OpenTag)
	}
}

func parseEm(str *strings.Builder, inEm bool) {
	if inEm {
		str.WriteString(mdParagraphsMap["*"].CloseTag)
	} else {
		str.WriteString(mdParagraphsMap["*"].OpenTag)
	}
}

func writeToFile(data []byte) error {
	return os.WriteFile("./out.html", data, 0644)
}
