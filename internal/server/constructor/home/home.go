package home

import "github.com/glebpepega/chanreader/internal/server/constructor/message"

func New(apiUrl string, chatId int) (err error) {
	var (
		ph       message.Photo
		keyboard [][]message.InlineKeyboardButton
	)

	ph.Chat_id = chatId

	ph.Photo = "https://mangabrog.files.wordpress.com/2015/08/ytbheadeer.jpg"

	row1 := []message.InlineKeyboardButton{
		{
			Text:          "/a",
			Callback_data: "/a",
		},
		{
			Text:          "/g",
			Callback_data: "/g",
		},
		{
			Text:          "/p",
			Callback_data: "/p",
		},
		{
			Text:          "/biz",
			Callback_data: "/biz",
		},
	}

	row2 := []message.InlineKeyboardButton{
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
		{
			Text:          "/x",
			Callback_data: "/x",
		},
	}

	keyboard = append(keyboard, row1, row2)

	ph.Reply_markup.Inline_keyboard = keyboard

	err = ph.Send(apiUrl)

	return
}
