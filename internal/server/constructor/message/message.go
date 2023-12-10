package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Photo struct {
	Chat_id      int                  `json:"chat_id"`
	Photo        string               `json:"photo"`
	Caption      string               `json:"caption"`
	Reply_markup InlineKeyboardMarkup `json:"reply_markup"`
}

type Text struct {
	Chat_id      int                  `json:"chat_id"`
	Text         string               `json:"text"`
	Reply_markup InlineKeyboardMarkup `json:"reply_markup"`
}

type InlineKeyboardMarkup struct {
	Inline_keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text          string `json:"text"`
	Callback_data string `json:"callback_data"`
}

type Sendable interface {
	Send(apiUrl string) (err error)
}

func (ph *Photo) Send(apiUrl string) (err error) {
	body, err := json.Marshal(ph)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(body)

	resp, err := http.Post(apiUrl+"/sendPhoto", "application/json", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	fmt.Println(string(bb), err)

	return
}

func (t *Text) Send(apiUrl string) (err error) {
	body, err := json.Marshal(t)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(body)

	resp, err := http.Post(apiUrl+"/sendMessage", "application/json", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	fmt.Println(string(bb), err)

	return
}
