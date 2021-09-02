package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type MyHandler struct {
	// все нужные вам объекты: конфиг, логер, соединение с базой и т.п.
}

// реализуем интерфейс `http.Handler`
func (h *MyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/search" {
		// разбираем аргументы
		args := req.URL.Query()
		query := args.Get("query")
		limit, err := strconv.Atoi(args.Get("limit"))
		if err != nil {
			panic("bad limit") // по-хорошему нужно возвращать HTTP 400
		}
		// выполняем бизнес-логику
		results, err := DoBusinessLogicRequest(query, limit)
		if err != nil {
			resp.WriteHeader(404)
			return
		}
		// устанавливаем заголовки ответа
		resp.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp.WriteHeader(200)
		// сериализуем и записываем тело ответа
		json.NewEncoder(resp).Encode(results)
	}
}

func main() {
	// создаем обработчик
	handler := &MyHandler{}
	// создаем HTTP сервер
	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// запускаем сервер, это заблокирует текущую горутину
	server.ListenAndServe()
}
