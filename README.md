# cobradoc [![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gavv/cobradoc) [![Build](https://github.com/gavv/cobradoc/workflows/build/badge.svg)](https://github.com/gavv/cobradoc/actions) [![GitHub release](https://img.shields.io/github/release/gavv/cobradoc.svg)](https://github.com/gavv/cobradoc/releases)

Alternative documentation generator for golang [Cobra](https://github.com/spf13/cobra).

Highlights
----------

* Supports markdown and manpage (troff) formats
* Supports command groups
* Generates single page for the whole command tree

Reference
---------

API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/gavv/cobradoc#section-documentation).

Example
-------

[_example](_example) directory demonstrates usage of this package.

It contains:

* [example cobra tool](_example/example.go)
* [generated manual page](_example/_manpage.md)
* [generated markdown](_example/_markdown.md)

Usage
-----

Generate markdown page:

```go
import "github.com/gavv/cobradoc"

err := cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Markdown, cobradoc.Options{
	Name:             "my-tool",
	Header:           "My page header",
	ShortDescription: "My tool description",
})
if err != nil {
	panic(err)
}
```

Generate manual page:

```go
import "github.com/gavv/cobradoc"

err := cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Troff, cobradoc.Options{
	Name:             "my-tool",
	Header:           "My page header",
	ShortDescription: "My tool description",
	ExtraSections: []cobradoc.ExtraSection{
		{
   			Title: cobradoc.BUGS,
			Text:  "Please report bugs via GitHub",
		},
	},
})
if err != nil {
    panic(err)
}
```

Credits
-------

This package is inspired by [cobraman](https://github.com/rayjohnson/cobraman) by Ray Johnson, but uses single-page approach and adds support for command groups.

Authors
-------

See [here](AUTHORS.md).

License
-------

[MIT](LICENSE)
