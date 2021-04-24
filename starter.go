package main

import (
	"encoding/json"
	"fmt"

	"math/rand"
	"time"

	txt "github.com/Aeriswave/TextFormer"
)

var ticker int = 0
var count int = 0
var lowNum int = -1
var upNum int = 50
var guess int = rand.Intn(upNum - lowNum - 1)
var number int = rand.Intn(upNum - lowNum - 1)
var ptik int = 0

func main() {
	var tmp txt.TextString = ""
	var MSG txt.TextTemplate = txt.TextTemplate{
		Top:         tmp,
		TopSplit:    tmp,
		Middle:      tmp,
		BottomSplit: tmp,
		Bottom:      tmp}

	//	var iMSG txt.IText = &MSG
	MSG.Set("Угадай число.", "Для обновления необходимо перезагрузить страницу", "")
	var bb string = *Ticker()
	MSG.Middle = txt.TextString(bb)
	fmt.Printf(MSG.Get())
	//	fmt.Printf(MSG.Get())
	//	fmt.Printf(string(MSG.Middle))
	return
}

func Ticker() *string {
	var tikerText string = "Новый тикер\n\n"

	ticker = (ticker+1)%120 - ticker%120 + ticker
	if count == 0 {
		tikerText = ""
	}
	if guess < 0 {
		guess = rand.Intn(16 + ticker)
	}

	if count <= 0 {
		tikerText = fmt.Sprintf("Это %d?\n ", guess) + tikerText
		count = -count + 1
	} else if guess != number {
		if guess < number {
			tikerText = fmt.Sprintf("Нет, не %d, а больше\n", guess) + tikerText
			lowNum = guess
			if lowNum >= upNum-1 {
				upNum = (lowNum + 1) << 2
				guess = upNum
			} else {
				guess = lowNum + (upNum-lowNum)>>1
			}
		} else {
			tikerText = fmt.Sprintf("Нет, не %d, а меньше\n", guess) + tikerText
			upNum = guess
			if lowNum >= upNum-1 {
				lowNum = -1
			}
			guess = upNum - (upNum-lowNum)>>1
		}
		count = -count
	} else {
		tikerText = fmt.Sprintf("\nУгадал c %d попытки, это %d\n", count, number) + tikerText
		count = 0
		guess = rand.Intn(10 + lowNum + upNum>>1 + ticker)
		upNum = 0
		lowNum = -1
		number = rand.Intn(10 + lowNum + upNum + ticker)
	}
	return &tikerText
}

type Message struct {
	Name string
	Text string
	//	Body   string
	Time   int64
	Ul     string
	Answer string
}

func ShowText() ([]byte, error) {
	var MSG txt.TextString = " "
	var iMSG txt.IText
	iMSG = &MSG
	iMSG.Set("Угадай число.", "Для обновления необходимо перезагрузить страницу", "")
	var middleText *string = Ticker()
	switch ticker % 2 {
	case ptik:
		iMSG.Set("Так, сделан шаг!")
	default:
		iMSG.Set(fmt.Sprintf("%d-й тик!", (ticker >> 1)))
	}
	iMSG.AddTopUD(
		fmt.Sprintf(
			"Угадай число быстрее робота: примерный диапазон %d...%d.", lowNum+1, upNum-1))
	iMSG.AddTopUD("Обнови страницу, что бы продолжить...")

	switch time.Now().Weekday() {
	case time.Monday:
		iMSG.AddTopUD("Сегодня понедельник.")
	case time.Tuesday:
		iMSG.AddTopUD("Сегодня вторник.")
	case time.Wednesday:
		iMSG.AddTopUD("Сегодня среда.")
	case time.Thursday:
		iMSG.AddTopUD("Сегодня четверг.")
	case time.Friday:
		iMSG.AddTopUD("Сегодня пятница.")
	case time.Saturday:
		iMSG.AddTopUD("Сегодня суббота.")
	case time.Sunday:
		iMSG.AddTopUD("Сегодня воскресенье.")
	}
	iMSG.Set("", "Игра функцией на Яндекс.облакЕ\n(c) Тряпицын Алексей\n", *middleText)
	m := Message{"Автор Тряпицын Алексей Васильевич", iMSG.Get(), 9057119603, "профи.сайт/АТ", fmt.Sprintf("Правильный ответ: %d", number)}

	return json.Marshal(m)
}
