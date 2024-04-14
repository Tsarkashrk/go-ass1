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
		log.Fatal("Open Connection Error:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping Connection Error:", err)
	}

	fmt.Println("DB OK")

	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorldHandler).Methods("GET")
	r.HandleFunc("/module", createModuleInfoHandler).Methods("POST")
	r.HandleFunc("/module/{id}", getModuleInfoHandler).Methods("GET")
	r.HandleFunc("/module/{id}", editModuleInfoHandler).Methods("PUT", "PATCH")
	r.HandleFunc("/module/{id}", deleteModuleInfoHandler).Methods("DELETE")

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

	moduleData := map[string]interface{}{
		"id":          1,
		"name":        "module 1",
		"description": "module 1 desc",
	}

	jsonData, err := json.Marshal(moduleData)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/module", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))

}