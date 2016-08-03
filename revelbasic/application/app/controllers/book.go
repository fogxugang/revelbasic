package controllers

import (
    "github.com/sadhanandh/revelbasic"
    "github.com/sadhanandh/revelbasic/application/app/models"
    "github.com/revel/revel"
)

type Book struct {
    *revel.Controller
    revelbasic.MongoController
}

type Category struct {
    *revel.Controller
    revelbasic.MongoController
}

func (c Book) Index() revel.Result {
    return c.Render()
}

func (c Book) View(id string) revel.Result {
    b := models.GetBookById(c.MongoSession, id)
    if b.Id.Hex() != id {
        return c.NotFound("Could not find a book with '%s' as id.", id)
    }
book := b
return c.Render(book)
    //return c.Render(b)
}

func (c Book) Categories() revel.Result {
    b := models.GetCategories(c.MongoSession)
    return c.RenderJson(b)
}

func (c Book) Brands() revel.Result {
    b := models.GetBrands(c.MongoSession)
    return c.RenderJson(b)
}
