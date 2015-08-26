# markdown

Simple markdown renderer and caching package for Forgery2 based on [blackfriday](http://github.com/russross/blackfriday).

	go get github.com/goforgery/markdown

* `.Render()` Function that renders the given filename as markdown.
* `.Clean()` Function that empties the cache of any rendered files.

## Use

    package main

    import(
        "fmt"
        "github.com/goforgery/markdown"
    )

    func init() {
        html := markdown.Render("./text.md")
        fmt.Println(html)
        markdown.Clean()
    }
