# links

This tool examines HTML files and reports invalid links.

## Installation

	go get github.com/mewmew/playground/cmd/links

## Usage

	links PATH...

## Examples

1. Report any invalid links in the set of files [a.html][] and [b.html][].

		cd testdata/
		links *.html
		// Output:
		// invalid fragment id to "b.html#a" in file "a.html".
		// invalid link target "c.html" in file "b.html".

[a.html]: https://raw.githubusercontent.com/mewmew/playground/master/cmd/links/testdata/a.html
[b.html]: https://raw.githubusercontent.com/mewmew/playground/master/cmd/links/testdata/b.html

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
