package main

import (
  "fmt"
  "time"
  "math/rand"
)
// Счётчик просмотров
var ticker int = 0
// Счётчик попыток угадать число
var count int = -1
var guess int = -1
var number int = rand.Intn(10)
var ptik int = 0

// Вызывается при загрузке страницы или обновлении страницы в браузере клиента
func Ticker() {
    // увеличение счётчика
    ticker=(ticker+1)%1000000000-ticker%1000000000+ticker

    switch ticker%4 {
    case 0:
        mdText="Так, сделай шаг!\n"+mdText
    case 1:
        mdText="\nОбновляй страницу, для продолжения...\n"+mdText
    case 2:
        mdText=fmt.Sprintf("%d\nТик!\n", ticker)+mdText
    case 3:
        mdText="\nУгадай число от 0 до 9.\n\n"+mdText
    default:
        WeekDayTicker() // Отображение информации о дне недели
    }
}
func Guess() {
    // Первая попытка угадать число
    if count ==-1 {
        guess=ticker%10
        count++
        number = rand.Intn(10)
    }
    ptik = (ticker)%4 // назначение формата отображения
    if guess != number {
        mdText=fmt.Sprintf("Это %d?\n", guess)+mdText
        if guess < number {
                mdText=fmt.Sprintf("Нет, не %d, а больше\n\n", guess, guess)+mdText
                guess++
            } else {
                mdText=fmt.Sprintf("Нет, не %d, а меньше\nЭто %d?\n\n", guess, guess)+mdText
                guess--
            }
            count++
        }  else  {
          mdText=fmt.Sprintf("\nУгадал, это %d\n\n", number)+mdText
          count=-1
    }

    }

 }

func WeekDayTicker() {
    switch time.Now().Weekday() {
    case time.Monday:
        mdText+="Сегодня понедельник.\n"
    case time.Tuesday:
        mdText+="Сегодня вторник.\n"
    case time.Wednesday:
        mdText+="Сегодня среда.\n"
    case time.Thursday:
        mdText+="Сегодня четверг.\n"
    case time.Friday:
        mdText+="Сегодня пятница.\n"
    case time.Saturday:
        mdText+="Сегодня суббота.\n"
    case time.Sunday:
        mdText+="Сегодня воскресенье.\n"
    }
  return
}


package main

import (
  "encoding/json"
)

type JSONString string

type Message struct {
    Name string
    Body string
    Time int64
    Ul string
}

func Handler() ([]byte, error) {
WeekDayTicker()
  m := Message{"Alice", mdText, 1294706395881547000, "www.ru"}
  return json.Marshal(m)
}

func (j JSONString) MarshalJSON() ([]byte, error) {
    return []byte(j), nil
}


package main
// Переменная для отображения на экране
var mdText string = "Первый запуск\n\n"
// Структура выводимого текста
type TextShaper struct {
    header string
    hi string
    text string
    low string
    titres string
}

// Clock методы ввода вывода в терминал
type Text interface {
	Replace(string) // Перезаписать текст
	AddToStart(string) // Добавить в начало текста
	AddToEnd(string) // Добавить в конец текста
}


