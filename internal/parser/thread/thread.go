package thread

import (
	"github.com/glebpepega/chanreader/internal/parser/board"
	"github.com/glebpepega/chanreader/internal/parser/post"
	"github.com/glebpepega/chanreader/internal/server/constructor/message"
	"golang.org/x/net/html"
	"strconv"
)

type Thread struct {
	Board  board.Board
	Number string
	Posts  []post.Post
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

func (t *Thread) Render(apiUrl string, chatId int) (err error) {
	for i, p := range t.Posts {
		var s message.Sendable

		text := p.Subject + "\n" + p.DateTime + "\nNo." + strconv.Itoa(p.Number) + "\n\n" + p.Message

		keyboard := make([][]message.InlineKeyboardButton, 0)

		if i == len(t.Posts)-1 {
			bRow := []message.InlineKeyboardButton{
				{
					Text:          "Board",
					Callback_data: t.Board.Name,
				},
			}

			hRow := []message.InlineKeyboardButton{
				{
					Text:          "Home",
					Callback_data: "H",
				},
			}

			keyboard = append(keyboard, bRow, hRow)
		}

		replyMarkup := message.InlineKeyboardMarkup{
			Inline_keyboard: keyboard,
		}

		if p.FileLink != "" {
			s = &message.Photo{
				Chat_id:      chatId,
				Photo:        p.FileLink,
				Caption:      text,
				Reply_markup: replyMarkup,
			}

			continue
		}

		if p.FileLink == "" {
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
