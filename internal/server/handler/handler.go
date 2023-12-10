package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Update struct {
	Update_id      int
	Message        Message
	Callback_query CallbackQuery
}

type Message struct {
	Chat Chat
	Text string
}

type Chat struct {
	ID int
}

type CallbackQuery struct {
	ID   string
	From User
}

type User struct {
	ID         int
	Is_bot     bool
	First_name string
}

type CallbackQueryAnswer struct {
	Callback_query_id string
	Text              string
}

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &Update{}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error(err.Error())

			return
		}

		if err := json.Unmarshal(b, u); err != nil {
			log.Error(err.Error())

			return
		}

		log.Info(
			"new request",
			"chat id", u.Message.Chat.ID,
		)

		//url := "https://boards.4channel.org/po/thread/615798/stuff"
		//
		//thr, err := thread.Parse(url)
		//if err != nil {
		//	log.Error(err.Error())
		//
		//	return
		//}
		//
		//br, err := board.Parse("https://boards.4channel.org/n/")
		//
		//_ = thr
		//_ = br
		//
		//fmt.Println(br)

		//fmt.Println(thr)
		//fmt.Println(thr.Posts[0].Subject)
	}
}
