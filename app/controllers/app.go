package controllers

import (
	"SearchTree/tree"
	"strings"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var Root *tree.Node

func (c App) DataPage() revel.Result {
	return c.Render()
}

func (c App) Data() revel.Result {
	if Root == nil {
		Root = tree.CreateNode()
	}

	input := c.Params.Get("data")

	words := strings.Split(input, " ")

	for _, word := range words {
		Root.AddWord(word)
	}

	return c.Redirect("/search")
}

func (c App) SearchPage(request string) revel.Result {
	return c.Render()
}

func (c App) Search(request string) revel.Result {
	input := c.Params.Get("data")

	var suggestions []string
	var has_results bool

	if Root != nil {
		suggestions = Root.GetSuggestions(input)
	}

	if len(suggestions) == 0 {
		has_results = false
	} else {
		has_results = true
	}

	return c.Render(suggestions, has_results)
}
