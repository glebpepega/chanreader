package board

import (
	"github.com/glebpepega/chanreader/internal/parser/post"
	"github.com/glebpepega/chanreader/internal/parser/thread"
	"github.com/glebpepega/chanreader/pkg/node"
	"golang.org/x/net/html"
)

type Board struct {
	Threads []thread.Thread
}

func (b *Board) ProcessChildNodes(n *html.Node) (err error) {
	for _, a := range n.Attr {
		if a.Val == "post op" {

			var (
				p post.Post
				t thread.Thread
			)

			if err = post.Parse(&p, n); err != nil {
				return
			}

			t.Posts = append(t.Posts, p)

			b.Threads = append(b.Threads, t)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err = b.ProcessChildNodes(c); err != nil {
			return
		}
	}

	return
}

func Parse(url string) (board Board, err error) {
	n, err := node.Get(url)
	if err != nil {
		return
	}

	err = board.ProcessChildNodes(n)

	return
}
