package main

import (
	"fmt"
	"log"
	"strconv"

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

	var mtroom int

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

	var user_chat_id int64
	//REALIZATION OF MENU MEETING ROOM CHANGE BUTTON
	var (
		// Universal markup builders.

		selector = &tele.ReplyMarkup{}
		btnPrev  = selector.Data("1", "prev")
		btnNext  = selector.Data("2", "next")
	)
	
	/*b.Handle("/start_button", func(c tele.Context) error {
		return c.Send("Hello!", selector)
	})*/



	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	b.Handle(&btnPrev, func(c tele.Context) error {
		/*var (
			text = c.Text()
		)*/

		mtroom = 1



		c.Send("Выбрана переговорка №" +strconv.Itoa(mtroom))

		return nil

	})

	b.Handle(&btnNext, func(c tele.Context) error {
		/*var (
			text = c.Text()
		)*/

		mtroom = 2

		//c.Send("Пожалйста, выбери переговорку", selector)


		c.Send("Выбрана переговорка №" + strconv.Itoa(mtroom))

		return nil

	})
	//END OF REALIZATION


	b.Handle("/setroom", func(c tele.Context) error {
		c.Send("Пожалйста, выбери переговорку", selector)
		return nil
	})
	b.Handle("/show", func(c tele.Context) error {
		var (
			//user = c.Sender()
			text = c.Text()
		)
		c.Send("Свободные слоты в переговорку")
		var show string
		
		switch mtroom {
			case 1:
				show = `SELECT * FROM meetings_1
				WHERE in_meet = $1`
			case 2:
				show = `SELECT * FROM meetings_2
				WHERE in_meet = $1`
			case 0:
				c.Send("Вы не выбрали переговорку")
				return nil
		}
	
		rows, err := db.Query(show,false)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var id int
			var comment string
			var time string
			var in_meet bool
			var user_name string
			var user_chat_id string
			var priority int 
			if err := rows.Scan(&id, &comment, &user_name,&user_chat_id,&priority, &time,  &in_meet); err != nil {
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
			//user = c.Sender()
			text = c.Text()
		)
		c.Send("Свободные слоты в переговорку")
		var show string
		
		switch mtroom {
			case 1:
				show = `SELECT * FROM meetings_1
				WHERE in_meet = $1`
			case 2:
				show = `SELECT * FROM meetings_2
				WHERE in_meet = $1`
			case 0:
				c.Send("Вы не выбрали переговорку")
				return nil
		}
		rows, err := db.Query(show, true)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var id int
			var comment string
			var time string
			var in_meet bool
			var user_chat_id string
			var priority int 
			var user_name string
			if err := rows.Scan(&id, &comment, &user_name,  &user_chat_id,&priority,&time, &in_meet); err != nil {
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
				var dbcheck string
				switch mtroom {
				case 1:
					dbcheck = `SELECT user_name,in_time from meetings_1 WHERE in_time = $1`
				case 2:
					dbcheck = `SELECT user_name,in_time from meetings_2 WHERE in_time = $1`
				case 0:
					c.Send("Вы не выбрали переговорку")
					return nil
			}
	
				
				var user_name_check string
				var time string
				if row := db.QueryRow(dbcheck, user_time); row != nil {
					err := row.Scan(&user_name_check, &time)
					if err != sql.ErrNoRows {
						if user_name_check != c.Sender().Username {
							return c.Send("Сожалеем но на это время записаны не вы")
						}
						
					}
				}
				var data string
		
		switch mtroom {
			case 1:
				data = ` UPDATE meetings_1 
				SET in_meet = true, comment = $1,user_name = $2, user_chat_id = $3
				WHERE in_time = $4`
			case 2:
				data = ` UPDATE meetings_2 
				SET in_meet = true, comment = $1,user_name = $2, user_chat_id = $3
				WHERE in_time = $4`
			case 0:
				c.Send("Вы не выбрали переговорку")
				return nil
		}
				user_time = text

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
				if !user_time_valid {
					return c.Send("Возможно такое время не предусмотрено")
				}

				user_chat_id = c.Sender().ID
				if _, err = db.Exec(data, user_comment, c.Sender().Username,user_chat_id, user_time); err != nil {
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
		c.Send(c.Text())
		var user_time string
		c.Send("Вот список ваших броней")

		var (
			text = c.Text()
		)
		var show string
		switch mtroom {
		case 1:
			show = `SELECT * FROM meetings_1
					WHERE user_name = $1`
		case 2:
			show = `SELECT * FROM meetings_2
					WHERE user_name = $1`
		case 0:
			c.Send("Вы не выбрали переговорку")
			return nil
	}
		
		rows, err := db.Query(show, c.Sender().Username)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var id int
			var comment string
			var time string
			var in_meet bool
			var user_name string
			var user_chat_id string
			var priority int 
			if err := rows.Scan(&id, &comment, &user_name, &user_chat_id,&priority,&time, &in_meet); err != nil {
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
		c.Send("Из списка введите время , по которому будет удалена запись")
		b.Handle(tele.OnText, func(c tele.Context) error {

			// All the text messages that weren't
			// captured by existing handlers.

			comment := "Никем не занята"

			user_name := ""
			user_id := ""

			var (
				text = c.Text()
			)

			user_time = text
			//

			//PLACE FOR CALLBACK FUNCTION

			//
		var data string
		
		switch mtroom {
			case 1:
				data = ` UPDATE meetings_1 
				SET in_meet = true, comment = $1,user_name = $2, user_chat_id = $3
				WHERE in_time = $4`
			case 2:
				data = ` UPDATE meetings_2 
				SET in_meet = true, comment = $1,user_name = $2, user_chat_id = $3
				WHERE in_time = $4`
			case 0:
				c.Send("Вы не выбрали переговорку")
				return nil
		}
			user_name_check_bool := true


			var dbcheck string
			switch mtroom {
			case 1:
				dbcheck = `SELECT user_name,in_time from meetings_1 WHERE in_time = $1`
			case 2:
				dbcheck = `SELECT user_name,in_time from meetings_2 WHERE in_time = $1`
			case 0:
				c.Send("Вы не выбрали переговорку")
				return nil
		}

			
			var user_name_check string
			var time string
			if row := db.QueryRow(dbcheck, user_time); row != nil {
				err := row.Scan(&user_name_check, &time)
				if err != sql.ErrNoRows {
					if user_name_check != c.Sender().Username {
						user_name_check_bool = false
						return c.Send("Сожалеем но на это время записаны не вы")
					} else {
						user_name_check_bool = true
					}

				}
			}

			if user_name_check_bool {
				if _, err = db.Exec(data, comment, user_name, user_id,user_time); err != nil {
					c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует")
					return c.Send(err)
				} else {
					c.Send(err)
				}
				// Instead, prefer a context short-hand:
				return c.Send(c.Sender().Username + "Ты удалил запись на" + " " + user_time)
			} else {

				return c.Send(c.Sender().Username + "Эта запись не существует или её делали не вы" + " " + user_time)
			}

		})

		return nil
	})

	b.Start()
}
