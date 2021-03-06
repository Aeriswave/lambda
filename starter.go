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
var guess int = rand.Intn(16)
var number int = rand.Intn(16)
var lowNum int = 0
var upNum int = 10
var ptik int = 0

func Ticker() {
	ticker = (ticker+1)%120 - ticker%120 + ticker
	if count == 0 {
		mdText = ""
	}
	if guess < 0 {
		guess = rand.Intn(16 + ticker)
	}

	if count <= 0 {
		mdText = fmt.Sprintf("Это %d?\n ", guess) + mdText
		count = -count + 1
	} else if guess != number {
		if guess < number {
			mdText = fmt.Sprintf("Нет, не %d, а больше\n", guess) + mdText
			lowNum = guess
			if lowNum >= upNum-1 {
				upNum = (lowNum + 1) << 1
				guess = upNum
			} else {
				guess = lowNum + rand.Intn(upNum-lowNum-1) + 1
			}
		} else {
			mdText = fmt.Sprintf("Нет, не %d, а меньше\n", guess) + mdText
			upNum = guess
			if lowNum >= upNum-1 {
				lowNum = (upNum>>1 - 1)
			}
			guess = upNum - rand.Intn(upNum-lowNum-1) - 1
		}
		count = -count
	} else {
		mdText = fmt.Sprintf("\nУгадал c %d попытки, это %d\n", count, number) + mdText
		count = 0
		guess = rand.Intn(16 + ticker)
		upNum = 0
		lowNum = 0
		number = rand.Intn(16 + ticker)
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
	var tmpDOWN string = "\nУгадай число.\n"
	Ticker()
	switch ticker % 2 {
	case ptik:
		tmpTOP += "Так, сделан шаг!\n"
	default:
		tmpTOP += fmt.Sprintf("%d-й тик!\n", (ticker >> 1))
	}
	tmpTOP += fmt.Sprintf("\nОбнови страницу, что бы продолжить...\nУгадай число быстрее робота: примерный диапазон %d...%d.\n", lowNum+1, upNum-1)

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
	tmpDOWN += "(c) AeriswavE\nИгра функцией на Яндекс.облакЕ\n"
	var tmpText string = tmpTOP + "\n\n" + mdText + "\n\n" + tmpDOWN

	m := Message{"AeriswavE", tmpText, 1234567, "профи.сайт/АТ"}
	return json.Marshal(m)
}

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}
