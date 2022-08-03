package main

import (
	"fmt"
	"log"

	//"os"
	"time"

	tele "gopkg.in/telebot.v3"

	_ "github.com/lib/pq"

	"database/sql"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "537j04222"
	dbname   = "postgres"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pref := tele.Settings{
		Token:  "5499433992:AAFHQ866_6-_YshxHPOp-oIpwv6X8XMxrPw",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/show", func(c tele.Context) error {

		var (
			//user = c.Sender()
			text = c.Text()
		)
		c.Send("Свободные слоты в переговорку")
		show := `SELECT * FROM meetings
					WHERE in_meet = $1`
		rows, err := db.Query(show, false)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var id int
			var comment string
			var time string
			var in_meet bool
			var user_name string
			if err := rows.Scan(&id, &comment, &user_name, &time, &in_meet); err != nil {
				log.Fatal(err)
			}
			text = time + " " + comment
			selector := &tele.ReplyMarkup{}
			btn := selector.Data(text, text)
			selector.Inline(
				selector.Row(btn),
			)

			c.Send(comment, selector)
		}
		return nil

	})

	b.Handle("/show_ordered", func(c tele.Context) error {
		var (
			text = c.Text()
		)
		c.Send("Записанные в  переговорки")
		show := `SELECT * FROM meetings
					WHERE in_meet = $1`
		rows, err := db.Query(show, true)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var id int
			var comment string
			var time string
			var in_meet bool
			var user_name string
			if err := rows.Scan(&id, &comment, &user_name, &time, &in_meet); err != nil {
				log.Fatal(err)
			}
			text = time + " " + comment
			selector := &tele.ReplyMarkup{}
			btn := selector.Data(text, text)
			selector.Inline(
				selector.Row(btn),
			)

			c.Send(user_name, selector)
		}
		return nil
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Я принимаю команды /cancel,/start,/show,/show_ordered!")
	})

	b.Handle("/start", func(c tele.Context) error {
		var user_time string
		var user_comment string
		c.Send("Желаете записаться? Оставьте комментарий")

		b.Handle(tele.OnText, func(c tele.Context) error {
			// All the text messages that weren't
			// captured by existing handlers.

			var (
				text = c.Text()
			)

			// Use full-fledged bot's functions
			// only if you need a result:
			user_comment = text
			c.Send(user_comment + "Твой комментарий")
			// Instead, prefer a context short-hand:
			c.Send("На какое время хочешь записаться ?")
			b.Handle(tele.OnText, func(c tele.Context) error {

				// All the text messages that weren't
				// captured by existing handlers.

				var (
					text = c.Text()
				)
				user_time = text
				// Instead, prefer a context short-hand:
				return c.Send(user_time + " " + "ТЫ запишешься на это время")
			})
			b.Handle(tele.OnText, func(c tele.Context) error {

				// All the text messages that weren't
				// captured by existing handlers.

				var (
					text = c.Text()
				)

				user_time = text

				data := `UPDATE meetings 
					SET in_meet = true 
					WHERE in_time = $1`
				user_time = text

				if _, err = db.Exec(data, user_time); err != nil {
					c.Send("Сожалеем но на это время уже кто-то записан")
					return c.Send(user_time)
				} else {
					c.Send("Удачно провёл запись")
				}
				// Instead, prefer a context short-hand:
				return c.Send(user_time + " " + "ТЫ запишешься на это время")
			})

			return nil
		})
		return nil
	})

	b.Handle("/cancel", func(c tele.Context) error {
		c.Send("С какого времени вы хотите убрать бронь")

		var (
			user = c.Sender()
			text = c.Text()
		)

		c.Update()
		c.Send(text)
		c.Send(user.FirstName)
		return nil
	})
	b.Start()
}
