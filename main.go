package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")

	if len(params) != 5 {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	num1, err1 := strconv.ParseFloat(params[3], 64)
	num2, err2 := strconv.ParseFloat(params[4], 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Неверные параметры", http.StatusBadRequest)
		return
	}

	var result float64
	switch params[2] {
	case "add":
		result = num1 + num2
	case "subtract":
		result = num1 - num2
	case "multiply":
		result = num1 * num2
	case "divide":
		if num2 == 0 {
			http.Error(w, "На ноль делить нельзя", http.StatusBadRequest)
			return
		}
		result = num1 / num2
	default:
		http.Error(w, "Неизвестная операция", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Результат: %.2f", result)
}

func main() {
	http.HandleFunc("/calc/", calculateHandler)

	fmt.Println("Запуск сервера на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
