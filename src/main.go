package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
)

func main() {
	iris.UseTemplate(html.New(html.Config{
		Layout: "layout.html",
	})).Directory("./templates", ".html")

	iris.Static("/public", "./static", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("home.html", Page{"Tonnarello", Pasta{"null", "null", "null"}}, iris.RenderOptions{"gzip": true})
	})

	iris.Post("/insert", func (ctx *iris.Context) {
		pasta := Pasta{}
		err := ctx.ReadForm(&pasta)
		if err != nil {
			fmt.Printf("ERR")
		}

		// TODO
		// Map the Pasta to a Mongo document here, save it

		ctx.Write("200 ok")
	})

	iris.Listen(":4000")
}
