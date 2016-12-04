package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dottorblaster/tonnarello/database"

	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
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

	iris.Post("/insert", func(ctx *iris.Context) {
		pasta := Pasta{}
		err := ctx.ReadForm(&pasta)
		if err != nil {
			fmt.Printf("ERR")
		}

		pasta.ID = bson.NewObjectId()

		err = pastas.Insert(pasta)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Redirect("pasta/"+pasta.ID.Hex(), http.StatusSeeOther)
	})

	iris.Get("/pasta/:id", func(ctx *iris.Context) {
		objID := bson.ObjectIdHex(ctx.Param("id"))
		pasta := Pasta{}

		pastas.FindId(objID).One(&pasta)
		ctx.Render("pasta.html", Page{"Tonnarello • " + pasta.Label, pasta}, iris.RenderOptions{"gzip": true})
	})

	iris.Get("/raw/:id", func(ctx *iris.Context) {
		objID := bson.ObjectIdHex(ctx.Param("id"))
		pasta := Pasta{}

		pastas.FindId(objID).One(&pasta)
		ctx.Text(iris.StatusOK, pasta.Content)
	})

	iris.Listen(":4000")

	defer mongoSession.Close()
}
