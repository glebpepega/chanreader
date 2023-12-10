package board

import (
	"errors"
	"github.com/glebpepega/chanreader/internal/parser/post"
	"github.com/glebpepega/chanreader/internal/parser/thread"
	"github.com/glebpepega/chanreader/internal/server/constructor/message"
	"golang.org/x/net/html"
	"strconv"
)

type Board struct {
	Name    string
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

func (b *Board) Render(apiUrl string, chatId int) (err error) {
	var s message.Sendable

	for i, t := range b.Threads {
		if len(t.Posts) == 0 {
			err = errors.New("empty thread")

			return
		}

		p := t.Posts[0]

		text := p.Subject + "\n" + p.DateTime + "\n\n" + p.Message
		threadNum := strconv.Itoa(p.Number)

		var keyboard [][]message.InlineKeyboardButton

		row := []message.InlineKeyboardButton{
			{
				Text:          "No." + threadNum,
				Callback_data: b.Name + "/thread/" + threadNum,
			},
		}

		keyboard = append(keyboard, row)

		if i == len(b.Threads)-1 {
			row = []message.InlineKeyboardButton{
				{
					Text:          "Home",
					Callback_data: "H",
				},
			}
			keyboard = append(keyboard, row)
		}

		replyMarkup := message.InlineKeyboardMarkup{
			Inline_keyboard: keyboard,
		}

		if t.Posts[0].FileLink != "" {
			s = &message.Photo{
				Chat_id:      chatId,
				Photo:        p.FileLink,
				Caption:      text,
				Reply_markup: replyMarkup,
			}
		}

		if t.Posts[0].FileLink == "" {
			s = &message.Text{
				Chat_id:      chatId,
				Text:         text,
				Reply_markup: replyMarkup,
			}
		}

		if err = s.Send(apiUrl); err != nil {
			return
		}
	}

	return
}
