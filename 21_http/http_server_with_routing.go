
/*
- import gorilla mux == routing
- justinas/alice == middleware
*/

type MyHandler struct{}

func (h *MyHandler) Search(resp ResponseWriter, req *Request) {
	// ...
}
func (h *MyHandler) AddItem(resp ResponseWriter, req *Request) {
	// ...
}
func main() {
	handler := &MyHandler{}
	// создаем маршрутизатор запросов
	mux := http.NewServeMux()
	mux.HandleFunc("/search", handler.Search)
	mux.HandleFunc("/add_item", handler.AddItem)
	// создаем и запускаем HTTP сервер
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}