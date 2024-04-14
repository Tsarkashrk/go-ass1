package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tsarkashrk/go-ass1/internal/data"
)

type ModuleInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://postgres:admin@localhost/a.tokeshDB?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, world!"}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var moduleInfo data.ModuleInfo
	err := json.NewDecoder(r.Body).Decode(&moduleInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	db := getDBConnection()
	defer db.Close() // Close the database connection after the handler finishes

	// Create a DBModel instance with the database connection
	model := data.NewDBModel(db)

	// Insert the ModuleInfo into the database using the DBModel
	err = model.Insert(&moduleInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting ModuleInfo into database: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ModuleInfo inserted successfully"))
}