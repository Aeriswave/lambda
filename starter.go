package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var ticker int = 0
var mdText string = "Первый запуск\n\n"
var count int = 0
var guess int = -1
var number int = rand.Intn(10)
var ptik int = 0

func Ticker() {
	ticker = (ticker+1)%1000000000 - ticker%1000000000 + ticker
	if count == 0 {
		mdText = ""
	}
	if (guess < 0) || (guess > 9) {
		guess = rand.Intn(10)
	}
	if count <= 0 {
		mdText = fmt.Sprintf("Это %d?", guess) + "\n\n" + mdText
		count = -count + 1
	} else if guess != number {
		if guess < number {
			mdText = fmt.Sprintf("Нет, не %d, а больше\n", guess) + mdText
			guess++
		} else {
			mdText = fmt.Sprintf("Нет, не %d, а меньше\n", guess) + mdText
			guess--
		}
		count = -count
	} else {
		mdText = fmt.Sprintf("\nУгадал c %d попытки, это %d\n", count, number) + mdText
		count = 0
		guess = -1
		number = rand.Intn(10)
	}
	return
}

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
	tmpTOP += "\nОбнови страницу, что бы продолжить...\nУгадай число от 0 до 9.\n"

	switch time.Now().Weekday() {

	case time.Monday:
		tmpTOP = "Сегодня понедельник.\n" + tmpTOP
	case time.Tuesday:
		tmpTOP = "Сегодня вторник.\n" + tmpTOP
	case time.Wednesday:
		tmpTOP = "Сегодня среда.\n" + tmpTOP
	case time.Thursday:
		tmpTOP = "Сегодня четверг.\n" + tmpTOP
	case time.Friday:
		tmpTOP = "Сегодня пятница.\n" + tmpTOP
	case time.Saturday:
		tmpTOP = "Сегодня суббота.\n" + tmpTOP
	case time.Sunday:
		tmpTOP = "Сегодня воскресенье.\n" + tmpTOP
	}

	var tmpText string = tmpTOP + "\n\n" + mdText + "\n\n" + tmpDOWN

	m := Message{"Автор сообщения", tmpText, 1294706395881547000, "ссылка"}
	return json.Marshal(m)
}

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}
