package thread

import (
	"github.com/glebpepega/chanreader/internal/parser/post"
	"github.com/glebpepega/chanreader/pkg/node"
	"golang.org/x/net/html"
)

type Thread struct {
	Posts []post.Post
}

func (t *Thread) ProcessChildNodes(n *html.Node) (err error) {
	for _, a := range n.Attr {
		if a.Val == "post op" ||
			a.Val == "post reply" {

			var p post.Post

			if err = post.Parse(&p, n); err != nil {
				return
			}

			t.Posts = append(t.Posts, p)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err = t.ProcessChildNodes(c); err != nil {
			return
		}
	}

	return
}

func Parse(url string) (thread Thread, err error) {
	n, err := node.Get(url)
	if err != nil {
		return
	}

	err = thread.ProcessChildNodes(n)

	return
}
