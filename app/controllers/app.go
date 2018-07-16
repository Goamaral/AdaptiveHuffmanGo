package controllers

import (
	"search_tree/tree"
	"strings"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var Root *tree.Node

func (c App) DataPage() revel.Result {

}

func (c App) Data() revel.Result {
	if Root == nil {
		Root = tree.CreateNode()
	}

	input := c.Params.Query.Get("data")

	words := strings.Split(input, " ")

	for _, word := range words {
		Root.AddWord(word)
	}

	c.Redirect("/search")
}

func (c App) SearchPage(request string) revel.Result {

}

func (c App) Search(request string) revel.Result {
	input := c.Params.Query.Get("data")

	suggestions := Root.GetSuggestions(input)

	c.Render(suggestions)
}
