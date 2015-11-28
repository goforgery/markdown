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
