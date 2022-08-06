package main

import (
	"fmt"
	"log"
	"strconv"

	//"os"
	"time"

	"gopkg.in/telebot.v3"
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
	dt := time.Now()
		our_time := dt.Format("15:04")

	go timechange(dt,our_time)



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
		selector = &tele.ReplyMarkup{}
		btnPrev  = selector.Data("1", "prev")
		btnNext  = selector.Data("2", "next")
	)

	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	b.Handle(&btnPrev, func(c tele.Context) error {

		mtroom = 1

		c.Edit("Выбрана переговорка №" + strconv.Itoa(mtroom))

		return nil
	})

	b.Handle(&btnNext, func(c tele.Context) error {
		mtroom = 2

		c.Edit("Выбрана переговорка №" + strconv.Itoa(mtroom))

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
		//c.Send("Свободные слоты в переговорку")
		show := show_msg(mtroom, c)
		if show == "nil" {
			return nil
		}

		rows, err := db.Query(show, false)
		if err != nil {
			log.Fatal(err)
		}

		selector := &telebot.ReplyMarkup{ResizeKeyboard: true}
		row := telebot.Row{}
		row2 := telebot.Row{}
		row3 := telebot.Row{}
		row4 := telebot.Row{}
		row5 := telebot.Row{}
		row6 := telebot.Row{}
		row7 := telebot.Row{}
		row8 := telebot.Row{}
		for rows.Next() {
			id, comment, time, in_meet, user_name, user_chat_id, priority := params()
			if err := rows.Scan(&id, &comment, &user_name, &user_chat_id, &priority, &time, &in_meet); err != nil {
				log.Fatal(err)
			}
			text = time 
			unique := fmt.Sprintf("Id:%d", id)
			btn := selector.Data(text,"task", unique)

			b.Handle(&btn, func(c tele.Context) error {
				c.Edit("Через свободные записи нельзя записаться , для этого есть /start")
		
				return nil
			})
			if len(row)<=1{
				row = append(row, btn)
			} else if len(row2)<=1{
				row2 = append(row2, btn)
				}else if len(row3)<=1{
					row3 = append(row3, btn)
				}else if len(row4)<=1{
					row4 = append(row4, btn)
				}else if len(row5)<=1{
					row5 = append(row5, btn)
				}else if len(row6)<=1{
					row6 = append(row6, btn)
				}else if len(row7)<=1{
					row7 = append(row7, btn)
				}else if len(row8)<=1{
					row8 = append(row8, btn)
				}
		}

		selector.Inline(
			row,
			row2,
			row3,
			row4,
			row5,
			row6,
			row7,
			row8,
		)


		/*selector.Data(text, text, text)
		c.Data() */
		c.Send("Все свободные слоты в переговорку", selector)
		return nil
	})

	b.Handle("/show_ordered", func(c tele.Context) error {
		var (
			//user = c.Sender()
			text = c.Text()
		)
		c.Send("Занятые слоты в переговорку")

		show := show_msg(mtroom, c)
		if show == "nil" {
			return nil
		}

		rows, err := db.Query(show, true)
		if err != nil {
			log.Fatal(err)
		}

		selector := &telebot.ReplyMarkup{ResizeKeyboard: true}
		row := telebot.Row{}
		row2 := telebot.Row{}
		row3 := telebot.Row{}
		row4 := telebot.Row{}
		row5 := telebot.Row{}
		row6 := telebot.Row{}
		row7 := telebot.Row{}
		row8 := telebot.Row{}

		
		for rows.Next() {
			id, comment, time, in_meet, user_name, user_chat_id, priority := params()
			if err := rows.Scan(&id, &comment, &user_name, &user_chat_id, &priority, &time, &in_meet); err != nil {
				log.Fatal(err)
			}
			text = time + " " + user_name + " " + comment
			unique := fmt.Sprintf("Id:%d", id)
			btn := selector.Data(text, "task", unique)
			if len(row)<=1{
				row = append(row, btn)
			} else if len(row2)<=1{
				row2 = append(row2, btn)
				}else if len(row3)<=1{
					row3 = append(row3, btn)
				}else if len(row4)<=1{
					row4 = append(row4, btn)
				}else if len(row5)<=1{
					row5 = append(row5, btn)
				}else if len(row6)<=1{
					row6 = append(row6, btn)
				}else if len(row7)<=1{
					row7 = append(row7, btn)
				}else if len(row8)<=1{
					row8 = append(row8, btn)
				}
		}
		selector.Inline(
			row,
			row2,
			row3,
			row4,
			row5,
			row6,
			row7,
			row8,
		)
		/*selector.Data(text, text, text)
		c.Data() */
		c.Send("Все Занятые слоты", selector)
		return nil
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Я принимаю команды /cancel,/start,/show,/show_ordered!")
	})

	b.Handle("/setadmin", func(c tele.Context) error {
		c.Send("Создание админа")


		c.Send("Если хотите стать админом, введите пороль")
			
			
		b.Handle(tele.OnText, func(c tele.Context) error {
			
				text := c.Text()
				password := "123"
			user_input := text 
			password_valid := false
			for i := range user_input{
				if user_input[i] == password[i]{
					password_valid = true
				} else{
					password_valid = false 
					break
				}
			}


			if password_valid {
				
				data := update_admin(mtroom, c)
				if data == "nil" {
					return nil
				}
				
				if _, err = db.Exec(data, c.Sender().Username); err != nil {
					c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует")
					return c.Send(err)
				} else {
					c.Send(err)
				}

				c.Send("Теперь вы админ")
			} else {
				c.Send("Что то не так с паролем")
			}

			return nil
			})
			

		return nil
	})


	b.Handle("/start", func(c tele.Context) error {
		var admin_prioritet int
		var user_slots_true bool = true
		var meetroom_count int
		if mtroom == 0{
			c.Send("Сначала выберите переговорку", selector)
		}
		
		var user_time string
		var user_comment string
		c.Send("Желаете записаться? Оставьте комментарий")

		b.Handle(tele.OnText, func(c tele.Context) error {

			var (
				text = c.Text()
			)
			user_comment = text
			c.Send(user_comment + " " + "Твой комментарий")
			// Instead, prefer a context short-hand:
			c.Send("На какое время хочешь записаться ?")
			b.Handle(tele.OnText, func(c tele.Context) error {
				var (
					text = c.Text()
				)
				user_time = text

				return c.Send(user_time + " " + "Ты записан на это время")
			})
			b.Handle(tele.OnText, func(c tele.Context) error {
				var (
					text = c.Text()
				)

				user_time = text
				// Проверка админ ли юзер
				prior := check_priority(mtroom, c)
				if prior == "nil" {
					return nil
				}

				


				if row := db.QueryRow(prior, c.Sender().Username); row != nil {
					err := row.Scan(&admin_prioritet)
					if err != sql.ErrNoRows {
						if admin_prioritet == 2{
							fmt.Println("admin rabotaet v systeme")
							
						}
					}
				}

					show := show_user(mtroom, c)
					if show == "nil" {
						return nil
					}

					rows, err := db.Query(show, c.Sender().Username)
					if err != nil {
						log.Fatal(err)
						fmt.Println(err)
					}
					for rows.Next() {
						id, comment, time, in_meet, user_name, user_chat_id, priority := params()
						if err := rows.Scan(&id, &comment, &user_name, &user_chat_id, &priority, &time, &in_meet); err != nil {
							log.Fatal(err)
							fmt.Println(err)
						}else {
							meetroom_count++
						}
						if meetroom_count == 4  && priority != 2 {
							c.Send("Вы не можете выбрать больше 4 слотов на запись в переговорку")
							user_slots_true = false
							return nil
						}
					}
				// Проверка админ ли юзер (конец)

				dbcheck := dbcheck_msg(mtroom, c)
				if dbcheck == "nil" {
					return nil
				}


				var user_name_check string
				var time string
				if row := db.QueryRow(dbcheck, user_time); row != nil {
					err := row.Scan(&user_name_check, &time)
					if err != sql.ErrNoRows {
						if user_name_check != c.Sender().Username && user_name_check != "" {
							return c.Send("Сожалеем но на это время записаны не вы, а" +" "+ user_name_check)
						}

					}
				}

				var data string
				if user_slots_true {
					data := data_msg(mtroom, c)
					if data == "nil" {
						return nil
					}
				} else {
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
				if _, err = db.Exec(data, user_comment, c.Sender().Username, user_chat_id, user_time); err != nil {
					c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует")
					return c.Send(err)
				} else {
					c.Send(err)
				}




				

				// Instead, prefer a context short-hand:
				return c.Send(c.Sender().Username + " " +"ТЫ запиcан на" + " " + user_time)
			})

			return nil
		})
		return nil
	})

	b.Handle("/cancel", func(c tele.Context) error {
		var user_time string

		var admin_prioritet int 
		if mtroom == 0{
			c.Send("Сначала выберите переговорку", selector)
		}
		var (
			text = c.Text()
		)
		//проверка на работу админа в системе
		prior := check_priority(mtroom, c)
				if prior == "nil" {
					return nil
				}

				


				if row := db.QueryRow(prior, c.Sender().Username); row != nil {
					err := row.Scan(&admin_prioritet)
					if err != sql.ErrNoRows {
						if admin_prioritet == 2{
							fmt.Println("admin rabotaet v systeme")
							
						}
					}
				}
		// конец проверки на админа
		show := show_user(mtroom, c)
		if show == "nil" {
			return nil
		}

		rows, err := db.Query(show, c.Sender().Username)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
			
		}

		selector := &telebot.ReplyMarkup{ResizeKeyboard: true}
		row := telebot.Row{}
		row2 := telebot.Row{}
		row3 := telebot.Row{}
		row4 := telebot.Row{}
		row5 := telebot.Row{}
		row6 := telebot.Row{}
		row7 := telebot.Row{}
		row8 := telebot.Row{}

		for rows.Next() {
			id, comment, time, in_meet, user_name, user_chat_id, priority := params()
			if err := rows.Scan(&id, &comment, &user_name, &user_chat_id, &priority, &time, &in_meet); err != nil {
				log.Fatal(err)
				fmt.Println(err)
				
			}
			text = time + " " + user_name + " " + comment
			unique := fmt.Sprintf("Id:%d", id)
			btn := selector.Data(text, "task", unique)
			if len(row)<=1{
				row = append(row, btn)
			} else if len(row2)<=1{
				row2 = append(row2, btn)
				}else if len(row3)<=1{
					row3 = append(row3, btn)
				}else if len(row4)<=1{
					row4 = append(row4, btn)
				}else if len(row5)<=1{
					row5 = append(row5, btn)
				}else if len(row6)<=1{
					row6 = append(row6, btn)
				}else if len(row7)<=1{
					row7 = append(row7, btn)
				}else if len(row8)<=1{
					row8 = append(row8, btn)
				}
		}
		selector.Inline(
			row,
			row2,
			row3,
			row4,
			row5,
			row6,
			row7,
			row8,
		)
		/*selector.Data(text, text, text)
		c.Data() */
		c.Send("Вот список занятых вами слотов", selector)

	
		c.Send("Из списка введите время , по которому будет удалена запись")
		b.Handle(tele.OnText, func(c tele.Context) error {
			comment := "Никем не занята"

			user_name := ""
			user_id := 0

			var (
				text = c.Text()
			)

			user_time = text
			//

			//PLACE FOR CALLBACK FUNCTION

			//
			data := data_msg_fasle(mtroom, c)
			if data == "nil" {
				return nil
			}


			














			user_name_check_bool := true

			dbcheck := dbcheck_msg(mtroom, c)
			if dbcheck == "nil" {
				return nil
			}

			var user_name_check string
			var time string
			switch admin_prioritet{
				case 0:
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

				case 2:
			}

			

			if user_name_check_bool {
				if _, err = db.Exec(data, comment, user_name, user_id, user_time); err != nil {
					c.Send("Произошла какая-то ошибка, возможно такого слота не сущетсвует")
					fmt.Println(err)
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

	if mtroom != 0 {
		go heartBeat(mtroom , db,our_time, b)
	}

	b.Start()
}
func params() (id int, comment string, time string,
	in_meet bool,
	user_name string,
	user_chat_id string,
	priority int) {
	return
}

func show_msg(mtroom int, c tele.Context) (show string) {
	switch mtroom {
	case 1:
		show = `SELECT * FROM meetings_1
			WHERE in_meet = $1`
	case 2:
		show = `SELECT * FROM meetings_2
			WHERE in_meet = $1`
	case 0:
		c.Send("Вы не выбрали переговорку")
		show = "nil"
	}
	return show
}

func show_msg_for_notif(mtroom int) (show string) {
	switch mtroom {
	case 1:
		show = `SELECT * FROM meetings_1
			WHERE in_meet = $1`
	case 2:
		show = `SELECT * FROM meetings_2
			WHERE in_meet = $1`
	case 0:
		show = "nil"
	}
	return show
}

func show_user(mtroom int, c tele.Context) (show string) {
	switch mtroom {
	case 1:
		show = `SELECT * FROM meetings_1
			WHERE user_name = $1`
	case 2:
		show = `SELECT * FROM meetings_2
			WHERE user_name = $1`
	case 0:
		c.Send("Вы не выбрали переговорку")
		show = "nil"
	}
	return show
}

func dbcheck_msg(mtroom int, c tele.Context) (dbcheck string) {
	switch mtroom {
	case 1:
		dbcheck = `SELECT user_name,in_time from meetings_1 WHERE in_time = $1`

	case 2:
		dbcheck = `SELECT user_name,in_time from meetings_2 WHERE in_time = $1`
	case 0:
		c.Send("Вы не выбрали переговорку")
		dbcheck = "nil"
	}
	return dbcheck
}

func check_priority(mtroom int,c tele.Context)(prior string){
	switch mtroom {
	case 1:
		prior = `SELECT priority from meetings_1 WHERE user_name = $1`

	case 2:
		prior = `SELECT priority from meetings_2 WHERE user_name = $1`
	case 0:
		c.Send("Вы не выбрали переговорку")
		prior = "nil"
	}
	return prior
}

func data_msg(mtroom int, c tele.Context) (data string) {
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
		data = "nil"
	}
	return data
}

func data_msg_fasle(mtroom int, c tele.Context) (data string) {
	switch mtroom {
	case 1:
		data = ` UPDATE meetings_1 
		SET in_meet = false, comment = $1,user_name = $2, user_chat_id = $3
		WHERE in_time = $4`

	case 2:
		data = ` UPDATE meetings_2 
		SET in_meet = false, comment = $1,user_name = $2, user_chat_id = $3
		WHERE in_time = $4`
	case 0:
		c.Send("Вы не выбрали переговорку")
		data = "nil"
	}
	
	return data
}

func update_admin(mtroom int,c tele.Context)(data string){
	switch mtroom {
	case 1:
		data = ` UPDATE meetings_1 
		SET priority = 2
		WHERE user_name = $1`

	case 2:
		data = ` UPDATE meetings_2 
		SET priority = 2
		WHERE user_name = $1`
	case 0:
		c.Send("Вы не выбрали переговорку")
		data = "nil"
	}
	
	return data
}


func notif_users(mtroom int, db * sql.DB,our_time string, b * tele.Bot){

	show := show_msg_for_notif(mtroom)

	rows, err := db.Query(show, true)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		id, comment, time, in_meet, user_name, user_chat_id, priority := params()
		if err := rows.Scan(&id, &comment, &user_name, &user_chat_id, &priority, &time, &in_meet); err != nil {
			log.Println(err)
		}
		user_chat_id_int,_ := strconv.Atoi(user_chat_id)
		if our_time == time{
			b.Send(&tele.User{ID: int64(user_chat_id_int)}, "У вас сейчас запись")
		}
	}

}

func heartBeat(mtroom int, db * sql.DB,our_time string, b * tele.Bot) {
    for range time.Tick(time.Second * 10) {
        notif_users(mtroom , db,our_time, b)
    }
}

func timechange(dt time.Time,our_time string){
	for range time.Tick(time.Second * 10) {
        dt = time.Now()
		our_time = dt.Format("15:04")
    }
}