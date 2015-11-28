# markdown

Simple markdown renderer and caching package for Forgery2 based on [blackfriday](http://github.com/russross/blackfriday).

* `.Render()` Function that renders the given filename as markdown.
* `.Clean()` Function that empties the cache of any rendered files.

## Install

	go get github.com/goforgery/markdown

## Use

```javascript
package main

import (
    "github.com/goforgery/forgery2"
    "github.com/goforgery/markdown"
)

func main() {
    app := f.CreateApp()
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.End(markdown.Render("./text.md"))
    })
    app.Listen(3000)
}
```

## Test

    go test
