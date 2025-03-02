# Markdown Test File

## Headings Demo

# Heading 1
## Heading 2
### Heading 3
#### Heading 4
##### Heading 5
###### Heading 6

## Text Formatting

Regular text looks like this.

**Bold text** looks like this.

*Italic text* looks like this.

***Bold and italic text*** looks like this.

~~Strikethrough text~~ looks like this.

`Inline code` looks like this.

> Blockquotes look like this.
> 
> They can span multiple lines.

## Lists

### Unordered List
* Item 1
* Item 2
  * Nested item 2.1
  * Nested item 2.2
* Item 3

### Ordered List
1. First item
2. Second item
   1. Nested item 2.1
   2. Nested item 2.2
3. Third item

### Task List
- [x] Completed task
- [ ] Incomplete task
- [ ] Another task

## Tables

| Name | Age | Occupation |
|------|-----|------------|
| John | 32  | Developer  |
| Jane | 28  | Designer   |
| Bob  | 45  | Manager    |

### Table With Alignment

| Left-aligned | Center-aligned | Right-aligned |
|:-------------|:--------------:|-------------:|
| Content      | Content        | Content      |
| Left         | Center         | Right        |

## Code Blocks

```python
def hello_world():
    print("Hello, world!")

hello_world()
```

```javascript
function sayHello() {
  console.log("Hello, world!");
}

sayHello();
```

## Links

[Visit GitHub](https://github.com)

---

## Advanced Elements

### Definition List

Term 1
: Definition 1

Term 2
: Definition 2

### Footnotes

Here's a sentence with a footnote[^1].

[^1]: This is the footnote.

## Nested Elements

> This is a blockquote containing a list
> 
> * Item 1
> * Item 2
>
> And a table
>
> | Column 1 | Column 2 |
> |----------|----------|
> | Value 1  | Value 2  |

## Escape Characters

\* This is not italic \*

\\ Backslash

\` Backtick

## HTML in Markdown

<details>
  <summary>Click to expand</summary>
  
  This content appears when expanded.
  
  It can include *formatted* text and other **markdown** elements.
</details>