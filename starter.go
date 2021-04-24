package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	//	jsText "lambda/lambda/github.com/aeriswave/jsonText/"
	//tst "lambda/lambda/tstmod"
	txt "github.com/aeriswave/jsonText"
)

//var nnn string = jsText.Text.Test2
var ticker int = 0
var mdText string = "Первый запуск\n\n"
var count int = 0
var lowNum int = -1
var upNum int = 50
var guess int = rand.Intn(upNum - lowNum - 1)
var number int = rand.Intn(upNum - lowNum - 1)
var ptik int = 0

func Main() {
	//	var ttt jsonText
	return
}

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
				upNum = (lowNum + 1) << 2
				guess = upNum
			} else {
				guess = lowNum + (upNum-lowNum)>>1
			}
		} else {
			mdText = fmt.Sprintf("Нет, не %d, а меньше\n", guess) + mdText
			upNum = guess
			if lowNum >= upNum-1 {
				lowNum = -1
			}
			guess = upNum - (upNum-lowNum)>>1
		}
		count = -count
	} else {
		mdText = fmt.Sprintf("\nУгадал c %d попытки, это %d\n", count, number) + mdText
		count = 0
		guess = rand.Intn(10 + lowNum + upNum>>1 + ticker)
		upNum = 0
		lowNum = -1
		number = rand.Intn(10 + lowNum + upNum + ticker)
	}
	return
}

type Message struct {
	Name   string
	Text   txt.TextTemplate
	Body   string
	Time   int64
	Ul     string
	Answer string
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
	tmpDOWN += "Игра функцией на Яндекс.облакЕ\n(c) Тряпицын Алексей\n"
	var tmpText string = tmpTOP + "\n\n" + mdText + "\n\n" + tmpDOWN

	var mmm txt.TextTemplate
	mmm.Set("Заголовок", "Текст", "Подтекст")
	m := Message{"Автор Тряпицын Алексей Васильевич", mmm.Get(), tmpText, 9057119603, "профи.сайт/АТ", fmt.Sprintf("Правильный ответ: %d", number)}

	return json.Marshal(m)
}
