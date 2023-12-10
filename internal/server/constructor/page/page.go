package page

import (
	"errors"
	"github.com/glebpepega/chanreader/internal/parser"
	"github.com/glebpepega/chanreader/pkg/node"
	"golang.org/x/net/html"
)

type ChildNodesProcessor interface {
	ProcessChildNodes(n *html.Node) (err error)
}

type Renderable interface {
	Render(apiUrl string, chatId int) (err error)
}

type Page interface {
	ChildNodesProcessor
	Renderable
}

func New(apiUrl string, chatId int, p Page) (err error) {
	var chanUrl string

	switch v := p.(type) {
	case *parser.Board:
		chanUrl = "https://boards.4channel.org" + v.Name
	case *parser.Thread:
		addr := v.Board.Name + "/thread/" + v.Number
		chanUrl = "https://boards.4channel.org" + addr
	default:
		err = errors.New("unexpected type")

		return
	}

	n, err := node.Get(chanUrl)
	if err != nil {
		return
	}

	if err = p.ProcessChildNodes(n); err != nil {
		return
	}

	err = p.Render(apiUrl, chatId)

	return
}
