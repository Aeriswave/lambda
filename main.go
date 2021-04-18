package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://functions.yandexcloud.net/d4ebb1qpmc2eo20avfan?requestId=1"
	fmt.Println("URL:>", url)
	/*
		type Identity struct {
			sourceIp  string //": "<адрес, с которого был сделан запрос>",
			userAgent string //": "<содержимое HTTP-заголовка User-Agent исходного запроса>"
		}
		type RequestContext struct {
			identity         Identity //": "<набор пар ключ:значение для аутентификации пользователя>",
			httpMethod       string   //": "<DELETE, GET, HEAD, OPTIONS, PATCH, POST или PUT>",
			requestId        string   //": "<ID запроса, генерируется в роутере>",
			requestTime      string   //": "<время запроса в формате CLF>",
			requestTimeEpoch string   //: "<время запроса в формате Unix>"
		}
		type JsonStruct struct {
			httpMethod                      string         //: "<название HTTP метода>",
			headers                         string         //: <словарь со строковыми значениями HTTP-заголовков>,
			multiValueHeaders               string         //: <словарь со списками значений HTTP-заголовков>,
			queryStringParameters           string         //": <словарь queryString-параметров>,
			multiValueQueryStringParameters string         //: <словарь списков значений queryString-параметров>,
			requestContext                  RequestContext //: <словарь с контекстом запроса>,
			body                            string         //": "<содержимое запроса>",
			isBase64Encoded                 bool           //": <true или false>
		}
		//m := JsonStruct {}
		//	json.MarshalJSON(m)
	*/
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
