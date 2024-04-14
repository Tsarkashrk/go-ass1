package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://postgres:admin@localhost/a.tokeshDB?sslmode=disable")
	if err != nil {
		log.Fatal("Ошибка при открытии соединения:", err)
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка при проверке соединения:", err)
	}

	fmt.Println("Соединение с базой данных установлено успешно")

	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorldHandler).Methods("GET")
	r.HandleFunc("/module", createModuleInfoHandler).Methods("POST")
	// r.HandleFunc("/module/{id}", getModuleInfoHandler).Methods("GET")
	// r.HandleFunc("/module/{id}", editModuleInfoHandler).Methods("PUT", "PATCH")
	// r.HandleFunc("/module/{id}", deleteModuleInfoHandler).Methods("DELETE")

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

	moduleData := map[string]interface{}{
		"id":          1,
		"name":        "Модуль 1",
		"description": "Описание модуля 1",
	}

	jsonData, err := json.Marshal(moduleData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Отправляем POST запрос на сервер
	resp, err := http.Post("http://localhost:8080/module", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от сервера
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Выводим ответ от сервера
	fmt.Println("Response:", string(body))

}