#### WEB-терминал, сделанный на основе сервиса публичной функции yandex cloud   
Архитектура решения:   
№ Terminal.go: Отображение страницы:   
- WEB-адрес функции   
- Вызываемая при запросе страницы функция   
- Отдаваемая по протоколу http JSON-структура   
№ Data.go: Данные и функционал для формирования отображаемого текста   
- Структура данных   
- Перезапись и добавление данных   
- Предоставление данных   
№ Events.go: Обработка событий   
- запрос на обновление экрана терминала   
