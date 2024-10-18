package main

import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,%s!", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/api/user", handler)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
