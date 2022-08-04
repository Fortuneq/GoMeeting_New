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
		c.Send("Все слоты в которые кто-то записан")
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
			c.Send(user_comment + " " + "Твой комментарий")
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
				return c.Send(user_time + " " + "Ты записан на это время")
			})
			b.Handle(tele.OnText, func(c tele.Context) error {

				// All the text messages that weren't
				// captured by existing handlers.

				var (
					text = c.Text()
				)

				user_time = text
				//

				//PLACE FOR CALLBACK FUNCTION

				//
				//dbcheck := `SELECT * FROM meetings WHERE in_time = $1`
				data := ` UPDATE meetings 
					SET in_meet = true, comment = $1,user_name = $2
					WHERE in_time = $3`
				user_time = text

				/*if rows, _ := db.Exec(dbcheck, user_time); err != nil {

					c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует или кто-то уже записан")
					return c.Send(rows)
				} else {
					fmt.Println(rows)
					c.Send(rows)
				}*/

				user_time_valid := false
				for i := 1; i < 10; i++ {
					if user_time == "19:30" {
						user_time_valid = false
						break
					}

					if user_time == fmt.Sprintf("1%d:%d0", i, 0) {
						user_time_valid = true
						break
					}

					if user_time == fmt.Sprintf("1%d:%d0", i, 3) {
						user_time_valid = true
						break
					}
				}
				if user_time_valid != true {
					return c.Send("Возможно такое время не предусмотрено")
				}

				if _, err = db.Exec(data, user_comment, c.Sender().Username, user_time); err != nil {
					c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует")
					return c.Send(err)
				} else {
					c.Send(err)
				}
				// Instead, prefer a context short-hand:
				return c.Send(c.Sender().Username + "ТЫ запиcан на" + " " + user_time)
			})

			return nil
		})
		return nil
	})

	b.Handle("/cancel", func(c tele.Context) error {
		var user_time string
		c.Send("С какого времени вы хотите убрать бронь")
		b.Handle(tele.OnText, func(c tele.Context) error {

			// All the text messages that weren't
			// captured by existing handlers.

			comment := ""

			user_name := ""

			var (
				text = c.Text()
			)

			user_time = text
			//

			//PLACE FOR CALLBACK FUNCTION

			//
			data := `UPDATE meetings 
					SET in_meet = false, comment = $1,user_name = $2
					WHERE in_time = $3`
			//user_time = text

			if _, err = db.Exec(data, comment, user_name, user_time); err != nil {
				c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует")
				return c.Send(err)
			} else {
				c.Send(err)
			}
			// Instead, prefer a context short-hand:
			return c.Send(c.Sender().Username + "Ты удалил запись на" + " " + user_time)
		})

		return nil
	})
	b.Start()
}
