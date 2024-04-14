package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Tsarkashrk/go-ass1/internal/data"
	"github.com/gorilla/mux"
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
	defer db.Close() 

	model := data.NewDBModel(db)

	err = model.Insert(&moduleInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting ModuleInfo into database: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ModuleInfo inserted successfully"))
}

func getModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid module ID", http.StatusBadRequest)
		return
	}

	idUint := uint(id)

	db := getDBConnection()
	defer db.Close()

	model := data.NewDBModel(db)
	moduleInfo, err := model.Retrieve(idUint) 
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving ModuleInfo from database: %v", err), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(moduleInfo)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}


func editModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid module ID", http.StatusBadRequest)
		return
	}

	var moduleInfo data.ModuleInfo
	err = json.NewDecoder(r.Body).Decode(&moduleInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	db := getDBConnection()
	defer db.Close()

	model := data.NewDBModel(db)
	existingModuleInfo, err := model.Retrieve(uint(id)) 
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving ModuleInfo: %v", err), http.StatusInternalServerError)
		return
	}

	err = model.Update(existingModuleInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating ModuleInfo in database: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ModuleInfo updated successfully"))
}

func deleteModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid module ID", http.StatusBadRequest)
		return
	}

	db := getDBConnection()
	defer db.Close()

	model := data.NewDBModel(db)
	err = model.Delete(uint(id)) 
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting ModuleInfo: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ModuleInfo deleted successfully"))
}
