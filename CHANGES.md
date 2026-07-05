# Changelog

## [v1.3.1][v1.3.1] - 05 Jul 2026

* Fix markdown TOC rendering when `Options.RenderHierarchy` is enabled
* Fix `Options.HideParents` with `Options.RenderHierarchy` to keep parent headings and descriptions, while hiding their usage and flags
* Fix `Options.HideParents` in troff output to hide non-leaf commands entirely
* Rename `hello` example to `groups`

[v1.3.1]: https://github.com/gavv/cobradoc/releases/tag/v1.3.1

## [v1.3.0][v1.3.0] - 05 Jul 2026

* Add `Options.HideCommands` to exclude selected commands from generated docs
* Add `Options.HideParents` to omit non-leaf commands from generated docs
* Add `Options.HideCompletion` to omit the generated completion command tree
* Add `Options.RenderHierarchy` to render nested markdown subcommands as nested sections
* Update minimum Go version to 1.19
* Replace Makefile with Taskfile

[v1.3.0]: https://github.com/gavv/cobradoc/releases/tag/v1.3.0

## [v1.2.0][v1.2.0] - 17 May 2025

* Add `Options.HideHelp` that hides `help` command and `--help` option ([gh-2][gh-2])

[v1.2.0]: https://github.com/gavv/cobradoc/releases/tag/v1.2.0

[gh-2]: https://github.com/gavv/cobradoc/issues/2

## [v1.1.0][v1.1.0] - 28 Jan 2023

* Add `Options.Language` and use it for text transformations ([gh-1][gh-1])
* Fix markdown formatting ([gh-1][gh-1])

[v1.1.0]: https://github.com/gavv/cobradoc/releases/tag/v1.1.0

[gh-1]: https://github.com/gavv/cobradoc/issues/1

## [v1.0.0][v1.0.0] - 28 Jan 2023

* Initial release

[v1.0.0]: https://github.com/gavv/cobradoc/releases/tag/v1.0.0
