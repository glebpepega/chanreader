package handler

import (
	"bytes"
	"encoding/json"
	"github.com/glebpepega/chanreader/internal/parser"
	"github.com/glebpepega/chanreader/internal/server/constructor/home"
	"github.com/glebpepega/chanreader/internal/server/constructor/page"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type Update struct {
	Update_id      int           `json:"update_id"`
	Message        Message       `json:"message"`
	Callback_query CallbackQuery `json:"callback_query"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	Id int `json:"id"`
}

type CallbackQuery struct {
	Id   string `json:"id"`
	From User   `json:"from"`
	Data string `json:"data"`
}

type User struct {
	Id         int    `json:"id"`
	Is_bot     bool   `json:"is_bot"`
	First_name string `json:"first_name"`
}

type CallbackQueryAnswer struct {
	Callback_query_id string `json:"callback_query_id"`
}

const (
	boardName = iota
	threadNum
)

func (cq *CallbackQuery) Answer(apiUrl string) (err error) {
	cqa := CallbackQueryAnswer{
		Callback_query_id: cq.Id,
	}

	body, err := json.Marshal(cqa)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(body)

	resp, err := http.Post(apiUrl+"/answerCallbackQuery", "application/json", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return
}

func New(log *slog.Logger, apiUrl string) http.HandlerFunc {
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

		var chatId int

		if u.Callback_query.Id != "" {
			if u.Callback_query.Data == "" {
				log.Error("empty callback data")

				return
			}

			if err := u.Callback_query.Answer(apiUrl); err != nil {
				log.Error(err.Error())

				return
			}

			chatId = u.Callback_query.From.Id

			log.Info(
				"new request",
				"chat id", chatId,
			)

			if u.Callback_query.Data == "H" {
				if err := home.New(apiUrl, chatId); err != nil {
					log.Error(err.Error())
				}

				return
			}

			var p page.Page

			if strings.Contains(u.Callback_query.Data, "thread") {
				sl := strings.Split(u.Callback_query.Data, "/thread/")

				if len(sl) != 2 {
					log.Error("unexpected callback data")

					return
				}

				p = &parser.Thread{
					Board: parser.Board{
						Name: sl[boardName],
					},
					Number: sl[threadNum],
				}
			} else {
				p = &parser.Board{
					Name: u.Callback_query.Data,
				}
			}

			if err := page.New(apiUrl, chatId, p); err != nil {
				log.Error(err.Error())
			}

			return
		}

		msg := u.Message.Text

		if msg == "/start" || msg == "/home" {
			chatId = u.Message.Chat.Id

			log.Info(
				"new request",
				"chat id", chatId,
			)

			if err := home.New(apiUrl, chatId); err != nil {
				log.Error(err.Error())

				return
			}
		}
	}
}
