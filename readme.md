## Question:

Write a function that converts a simplified subset of Markdown to HTML. Your parser should support the following Markdown features:

```
# for <h1> headers

## for <h2> headers

**bold** for <strong>

*italic* for <em>

Paragraphs for plain text (wrapped in <p> tags)

Line breaks (empty lines separate paragraphs)

```

## Assumptions:

1. The input contains no HTML tags, only simplified Markdown.

2. There are no nested formatting (e.g., **bold *italic*** will not appear).

3. Headings always start at the beginning of a line.

4. Bold and italic are always properly closed.