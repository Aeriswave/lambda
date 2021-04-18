package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type JSONString string

type Message struct {
	Name string
	Body string
	Time int64
	Ul   string
}

func Handler() ([]byte, error) {
	var tmpTOP string = ""
	var tmpDOWN string = "\nУгадай число от 0 до 9.\n"
	Ticker()
	switch ticker % 2 {
	case ptik:
		tmpTOP += "Так, сделан шаг!\n"
	default:
		tmpTOP += fmt.Sprintf("%d-й тик!\n", (ticker >> 1))
	}
	tmpTOP += "\nОбнови страницу, что бы продолжить...\n\n"

	switch time.Now().Weekday() {

	case time.Monday:
		tmpTOP = "Сегодня понедельник." + tmpTOP
	case time.Tuesday:
		tmpTOP = "Сегодня вторник." + tmpTOP
	case time.Wednesday:
		tmpTOP = "Сегодня среда." + tmpTOP
	case time.Thursday:
		tmpTOP = "Сегодня четверг." + tmpTOP
	case time.Friday:
		tmpTOP = "Сегодня пятница." + tmpTOP
	case time.Saturday:
		tmpTOP = "Сегодня суббота." + tmpTOP
	case time.Sunday:
		tmpTOP = "Сегодня воскресенье." + tmpTOP
	}

	var tmpText string = tmpTOP + "\n\n" + mdText + "\n\n" + tmpDOWN

	m := Message{"Автор сообщения", tmpText, 1294706395881547000, "ссылка"}
	return json.Marshal(m)
}

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

var ticker int = 0
var mdText string = "Первый запуск\n\n"
var count int = -1
var guess int = rand.Intn(10)
var number int = rand.Intn(10)
var ptik int = 0

func Ticker() {
	ticker = (ticker+1)%1000000000 - ticker%1000000000 + ticker
	if count == -1 {
		guess = ticker % 10
		number = rand.Intn(10)
		ptik = (ticker) % 2
	}
	if guess != number {
		if count < 0 {
			mdText = fmt.Sprintf("Это %d?", guess) + "\n\n" + mdText
			count = 1
		} else {
			if guess < number {
				mdText = fmt.Sprintf("Нет, не %d, а больше\n", guess) + mdText
				guess++
			} else {
				mdText = fmt.Sprintf("Нет, не %d, а меньше\n", guess) + mdText
				guess--
			}
			count = -1
		}
	} else {
		mdText = fmt.Sprintf("\nУгадал, это %d\n", number) + mdText
		count = -1
	}
	return
}
