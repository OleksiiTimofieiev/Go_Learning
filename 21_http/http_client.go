package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/*
- http == hyper text protocol
- GET, POST, PUT, DELETE \r\n => 200, 404
- /etc/ssl => ssl certs
- req.with context
- http.Client is thread safe, supports pool of connections
- config of connection == http.Transport
- framework for authorization
*/

type AddRequest struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Создаем буфер (io.Reader) из которого клиент возьмет тело запроса
var body bytes.Buffer

func main() {
	// создаем HTTP клиент
	client := &http.Client{}

	// строим нужный URL
	reqArgs := url.Values{}
	reqArgs.Add("query", "go syntax")
	reqArgs.Add("limit", "5")
	reqUrl, _ := url.Parse("https://google.com")
	reqUrl.Path = "/search"
	reqUrl.RawQuery = reqArgs.Encode()

	// создаем GET-запрос
	req, _ := http.NewRequest("GET", reqUrl.String(), nil)
	fmt.Println(req)

	// выполняем запрос
	req.Header.Add("User-Agent", `Mozilla/5.0 Gecko/20100101 Firefox/39.0`)
	resp, _ := client.Do(req)

	// Запрос в виде Go структуры
	addReq := &AddRequest{
		Id:    123,
		Title: "for loop",
		Text:  "...",
	}

	json.NewEncoder(body).Encode(addReq)

	// создаем POST-запрос
	req, err := http.NewRequest("POST", "https://site.ru/add_item", body)
	// выполняем запрос
	resp, err = client.Do(req)
	if err != nil {
		// или другая уместная обработка
		log.Fatal(err)
	}

	// если ошибки не было - нам необходимо "закрыть" тело ответа
	// иначе при повторном запросе будет открыто новое сетевое соединение
	defer resp.Body.Close()
	// проверяем HTTP status ответа
	if resp.StatusCode != 200 {
		// обработка HTTP статусов зависит от приложения
		return fmt.Errorf("unexpected http status: %s", resp.Status)
	}
	// возможно проверяем какие-то заголовки
	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		log.Fatalf("unexpected content-type: %s", ct)

		fmt.Println(ioutil.ReadAll(resp.Body))
	}
}
