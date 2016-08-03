package controllers

import (
    "github.com/revel/revel"
)

type App struct {
    *revel.Controller
}

func addHeaderCORS(c *revel.Controller) revel.Result {
    c.Response.Out.Header().Add("Access-Control-Allow-Origin","*")
    return nil
}

func init() {
    revel.InterceptFunc(addHeaderCORS, revel.AFTER, &App{})
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) saveToDb() revel.Result {
    return c.Render()
}
