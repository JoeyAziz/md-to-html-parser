package markdown

import (
	"testing"
)

const colorRed = "\033[0;31m"

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Header and paragraph",
			input: `# Heading

This is a **bold** and *italic* test.`,
			expected: `<h1>Heading</h1><br /><p>This is a <strong>bold</strong> and <em>italic</em> test.</p>`,
		},
		{
			name:     "Subheader only",
			input:    `## Subheading`,
			expected: `<h2>Subheading</h2>`,
		},
		{
			name:     "Plain paragraph",
			input:    `Just a simple paragraph.`,
			expected: `<p>Just a simple paragraph.</p>`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Parse(tc.input)
			if result != tc.expected {
				t.Errorf("\nExpected:\n%q\n%sGot:\n%q", tc.expected, colorRed, result)
			}
		})
	}
}
