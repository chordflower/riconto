---
title: Extensions
tags:
  - scratch
  - ideas
  - markdown extensions
---

## Markdown Extensions ##

This is a list of possible available markdown extensions:

- [goldmark-supersubscript](https://github.com/bowman2001/goldmark-supersubscript)
- [goldmark-figure](https://github.com/MangoUmbrella/goldmark-figure)
- [goldmark-callout](https://gitlab.com/staticnoise/goldmark-callout)
- extension.Table
- extension.Strikethrough
- extension.TaskList
- extension.DefinitionList
- extension.Footnote
- extension.Typographer
- [goldmark-highlighting](https://github.com/yuin/goldmark-highlighting)
- [goldmark-emoji](https://github.com/yuin/goldmark-emoji)
- [goldmark-mathjax](https://github.com/litao91/goldmark-mathjax)
- [goldmark-anchor](https://github.com/abhinav/goldmark-anchor)
- [goldmark-toc](https://github.com/abhinav/goldmark-toc)

### Directive Based Extensions ###

These are [directive](https://talk.commonmark.org/t/generic-directives-plugins-syntax/444) based extensions.

#### Inline Directive Syntax ####

The syntax for inline directives:

```
:name[content]{key=val}
```

Exactly one colon, followed by the name which is the identifier for the extension and must be a string without spaces, content is text (unlike the linked proposed extension) and the `{key=val key2="val 2"}` contain generic attributes (i.e. key-value pairs) and are optional.

Aka the regular expression:

```regexp
:([^\[]+)\[([^\]]+)\](?:\{([^\}]+)\})?
```

Which extracts three groups:

1. The name (required);
2. The content (required);
3. The key-values (grouped together) (optional).

#### Leaf Block Directives ####

The syntax for leaf block directives:

```
::name[content]{key=val}
```

To be recognized as a directive, this has to form an otherwise empty paragraph. But as opposed to inline directives, there are two colons now, the brackets [] are optional as well, and spaces may be interspersed for readability.

Aka the regular expression:

```regex
:([^\[]+)(?:\[([^\]]+)\])?(?:\{([^\}]+)\})?
```

Which extracts three groups:

1. The name (required);
2. The content (optional);
3. The key-values (grouped together) (optional).

#### List of directives ####

- `::include[<markdown_file_to_include>]` => Parses the given markdown file and includes it in the current document ast. (Careful with recursive includes!);
- `::embed[<thing_to_embed>]{type=TYPE}` => Includes the given url using oEmbed, uses [go-oembed](https://github.com/dyatlov/go-oembed), to return oEmbed information;
- `::toc` => Includes a table of contents;
