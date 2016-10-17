package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
)

func main() {
	iris.UseTemplate(html.New(html.Config{
		Layout: "layout.html",
	})).Directory("./templates", ".html")

	iris.Static("/public", "./static", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("home.html", Page{"Tonnarello", Pasta{"null", "null"}}, iris.RenderOptions{"gzip": true})
	})

	iris.Listen(":4000")
}
