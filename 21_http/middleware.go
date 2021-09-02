// это функция - middleware, она преобразует один обработчик в другой
func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if !currentUser(req).IsAdmin {
			http.NotFound(resp, req)
			return
		}
		h(resp, req)
	}
}

func (h *MyHandler) AddItem(resp ResponseWriter, req *Request) {
	ctx := req.Context()
	user := ctx.Value("currentUser").(*MyUser)
	// ...
}

func (h *MyHandler) Search(resp ResponseWriter, req *Request) {
	ctx := req.Context()
	// ...
	// мы должны передавать контекст вниз по всем вызовам
	results, err := DoBusinessLogicRequest(ctx, query, limit)
	// ...
}

func authorize(h http.HandlerFunc, timeout time.Duration) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		// выполняем авторизацию пользователя
		user, err := DoAuthorizeUser(req)
		if err != nil {
			// если не удалось - возвращаем соответствующий HTTP статус
			resp.WriteHeader(403)
			return
		}
		// сохраняем пользователя в контекст
		ctx := context.WithValue(req.Context(), "currentUser", user)
		req = req.WithContext(ctx)
		h(resp, req)
	}
}

func withTimeout(h http.HandlerFunc, timeout time.Duration) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		// берем контекст запроса и ограничиваем его таймаутом
		ctx := context.WithTimeout(req.Context(), timeout)
		// обновляем контекст запроса
		req = req.WithContext(ctx)
		h(resp, req)
	}
}

func main() {
	handler := &MyHandler{}
	// создаем маршрутизатор запросов
	mux := http.NewServeMux()
	// mux.HandleFunc("/search", handler.Search)
	mux.HandleFunc("/search", withTimeout(handler.Search, 5*time.Second))
	// !!! мы обернули один из обработчиков в middleware
	// mux.HandleFunc("/add_item", adminOnly(handler.AddItem))
	mux.HandleFunc("/add_item", authorize(handler.AddItem))
}