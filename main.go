package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	db "r3kk3/src/database"
	models "r3kk3/src/models"

	"github.com/gorilla/mux"
)

var database sql.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	jsonResponse, jsonError := json.Marshal(
		models.Response{"Hello user"})
	models.Return_json(w, jsonResponse, jsonError)
}

func AddWorker(w http.ResponseWriter, r *http.Request) {
	var new_worker models.Worker
	json.NewDecoder(r.Body).Decode(&new_worker)
	worker_id, err := db.InsertWorker(database, new_worker)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Add worker failed"})
		models.Return_json(w, jsonResponse, jsonError)
	} else {
		strid := strconv.Itoa(worker_id)
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Worker №" + strid + " added successfuly"})
		models.Return_json(w, jsonResponse, jsonError)
	}
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
	worker_id := mux.Vars(r)["id"]
	_, err := db.DeleteWorker(database, worker_id)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Delete worker failed"})
		models.Return_json(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Worker №" + worker_id + " delete successfuly"})
		models.Return_json(w, jsonResponse, jsonError)
	}
}

func GetWorker(w http.ResponseWriter, r *http.Request) {
	worker_id := mux.Vars(r)["id"]
	worker, err := db.GetWorker(database, worker_id)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Get worker failed"})
		models.Return_json(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(worker)
		models.Return_json(w, jsonResponse, jsonError)
	}
}

func GetWorkersByCompanyId(w http.ResponseWriter, r *http.Request) {
	company_id := mux.Vars(r)["company_id"]
	workers, err := db.GetWorkersByCompanyId(database, company_id)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Get workers by company id failed"})
		models.Return_json(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(workers)
		models.Return_json(w, jsonResponse, jsonError)
	}

}

func GetWorkersByDepartment(w http.ResponseWriter, r *http.Request) {
	department_name := mux.Vars(r)["name"]
	workers, err := db.GetWorkersByDepartment(database, department_name)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Get workers by department name failed"})
		models.Return_json(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(workers)
		models.Return_json(w, jsonResponse, jsonError)
	}
}

func ChangeWorker(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	worker_id := mux.Vars(r)["id"]
	json.NewDecoder(r.Body).Decode(&worker)
	nworker, err := db.ChangeWorker(database, worker, worker_id)
	if err != nil {
		jsonResponse, jsonError := json.Marshal(
			models.Response{"Change worker failed"})
		models.Return_json(w, jsonResponse, jsonError)
	} else {
		jsonResponse, jsonError := json.Marshal(nworker)
		models.Return_json(w, jsonResponse, jsonError)
	}
}

func main() {

	// Подключаемся к БД
	database = db.InitDB()
	time.Sleep(1 * time.Second)

	// Создаем таблицу паспортов
	db.CreatePassportTable(database)

	// Создаем таблицу отделов
	db.CreateDepartmentTable(database)

	// Создаем таблицу сотрудников
	db.CreateWorkerTable(database)

	// Инициализируем роуты
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/add_worker", AddWorker).Methods("POST")
	router.HandleFunc("/delete_worker/{id}", DeleteWorker).Methods("POST")
	router.HandleFunc("/get_worker/{id}", GetWorker).Methods("GET")
	router.HandleFunc("/get_workers_by_company_id/{company_id}", GetWorkersByCompanyId).Methods("GET")
	router.HandleFunc("/get_workers_by_department/{name}", GetWorkersByDepartment).Methods("GET")
	router.HandleFunc("/change_worker/{id}", ChangeWorker).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
