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
		c.Send("Все Готовые переговорки")
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
			if err := rows.Scan(&id, &comment, &time, &in_meet); err != nil {
				log.Fatal(err)
			}
			text = time + " " + comment
			selector := &tele.ReplyMarkup{}
			btn := selector.Data(text, text)
			selector.Inline(
				selector.Row(btn),
			)

			c.Send(comment,selector)
		}
		return nil

	})

	b.Handle("/show_ordered", func(c tele.Context) error {
		var (
			//user = c.Sender()
			text = c.Text()
		)
		c.Send("Все зарегестрированные переговорки")
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
			if err := rows.Scan(&id, &comment, &time, &in_meet); err != nil {
				log.Fatal(err)
			}
			text = time 
			selector := &tele.ReplyMarkup{}
			btn := selector.Data(text, text)
			selector.Inline(
				selector.Row(btn),
			)

			c.Send(comment,selector)
		}
		return nil
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Я принимаю команды /cancel,/start,/show!")
	})

	b.Handle("/start", func(c tele.Context) error {
		selector := &tele.ReplyMarkup{}

		btn1 := selector.Data("11:00", "11:00")
		btn2 := selector.Data("11:30", "11:30")
		btn3 := selector.Data("12:00", "12:00")
		btn4 := selector.Data("12:30", "12:30")
		btn5 := selector.Data("13:00", "13:00")
		btn6 := selector.Data("13:30", "13:30")
		btn7 := selector.Data("14:00", "14:00")
		btn8 := selector.Data("14:30", "14:30")
		btn9 := selector.Data("15:00", "15:00")

		//selector.InlineKeyboard = append(selector.InlineKeyboard,btn1)
		selector.Inline(
			selector.Row(btn1, btn2, btn3),
			selector.Row(btn4, btn5, btn6),
			selector.Row(btn7, btn8, btn9),
		)

		// On inline button pressed (callback)
		c.Send("Выберите время для записи", selector)
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

	/*b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		var (
			//user = c.Sender()
			text = c.Text()
		)

		// Use full-fledged bot's functions
		// only if you need a result:

		// Instead, prefer a context short-hand:
		return c.Send(text)
	})
*/
	b.Start()
}
