package post

import (
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

type Post struct {
	Subject  string
	DateTime string
	Number   int
	Message  string
	FileLink string
}

func Parse(post *Post, n *html.Node) (err error) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "blockquote" {
			for bqc := c.FirstChild; bqc != nil; bqc = bqc.NextSibling {
				if bqc.Data == "a" && bqc.FirstChild != nil {
					post.Message += bqc.FirstChild.Data
				}

				if bqc.Data == "span" && len(bqc.Attr) > 0 && bqc.Attr[0].Val == "quote" {
					for sc := bqc.FirstChild; sc != nil; sc = sc.NextSibling {
						if sc.Type == html.TextNode {
							post.Message += sc.Data
						}
					}
				}

				if bqc.Data == "br" {
					post.Message += "\n"
				}

				if bqc.Type == html.TextNode {
					post.Message += bqc.Data
				}
			}
		}

		var flNode bool

		for _, a := range c.Attr {
			if a.Val == "subject" && c.FirstChild != nil {
				post.Subject = c.FirstChild.Data
			}

			if a.Val == "dateTime" && c.FirstChild != nil {
				post.DateTime = c.FirstChild.Data
			}

			if a.Val == "Reply to this post" && c.FirstChild != nil {
				post.Number, err = strconv.Atoi(c.FirstChild.Data)
			}

			if a.Val == "fileThumb" {
				flNode = true
			}
		}

		if flNode {
			for _, a := range c.Attr {
				if a.Key == "href" {
					post.FileLink = strings.TrimPrefix(a.Val, `//`)
				}
			}
		}

		if (post.Subject != "" &&
			post.DateTime != "" &&
			post.Number != 0 &&
			post.Message != "" &&
			post.FileLink != "") ||
			err != nil {
			return
		}

		err = Parse(post, c)
	}

	post.Message = strings.TrimSuffix(post.Message, "\n")

	return
}
