package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dottorblaster/tonnarello/database"

	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
)

func main() {
	mongoSession, pastas, err := database.NewPastasConnection()
	if err != nil {
		log.Fatal(err)
	}

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

		pasta.Id = bson.NewObjectId()

		err = pastas.Insert(pasta)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Redirect("pasta/" + pasta.Id.Hex(), http.StatusSeeOther)
	})

	iris.Get("/pasta/:id", func(ctx *iris.Context) {
		objId := bson.ObjectIdHex(ctx.Param("id"))
		pasta := &Pasta{}

		pastas.FindId(objId).One(pasta)
		ctx.Render("pasta.html", Page{"Tonnarello", pasta}, iris.RenderOptions{"gzip": true})
	})

	iris.Listen(":4000")

	defer mongoSession.Close()
}
