package home

import "github.com/glebpepega/chanreader/internal/server/constructor/message"

func New(apiUrl string, chatId int) (err error) {
	var (
		ph       message.Photo
		keyboard [][]message.InlineKeyboardButton
	)

	ph.Chat_id = chatId

	ph.Photo = "https://mangabrog.files.wordpress.com/2015/08/ytbheadeer.jpg"

	row := []message.InlineKeyboardButton{
		{
			Text:          "/mu",
			Callback_data: "/mu",
		},
		{
			Text:          "/lit",
			Callback_data: "/lit",
		},
		{
			Text:          "/tv",
			Callback_data: "/tv",
		},
	}

	keyboard = append(keyboard, row)

	ph.Reply_markup.Inline_keyboard = keyboard

	err = ph.Send(apiUrl)

	return
}
