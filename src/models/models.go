package models

import (
	"encoding/json"
	"net/http"
)

type Passport struct {
	Id     int
	Type   string
	Number string
}

type Department struct {
	Id    int
	Name  string
	Phone string
}

type Response struct {
	Message string
}

type Worker struct {
	Id         int
	Name       string
	Surname    string
	Phone      string
	CompanyId  int
	Passport   Passport
	Department Department
}

func (w *Worker) OverlayFields(ow Worker) {
	if ow.Id != 0 {
		w.Id = ow.Id
	}
	if ow.Name != "" {
		w.Name = ow.Name
	}
	if ow.Surname != "" {
		w.Surname = ow.Surname
	}
	if ow.Phone != "" {
		w.Phone = ow.Phone
	}
	if ow.CompanyId != 0 {
		w.CompanyId = ow.CompanyId
	}
	if ow.Passport.Id != 0 {
		w.Passport.Id = ow.Passport.Id
	}
	if ow.Passport.Type != "" {
		w.Passport.Type = ow.Passport.Type
	}
	if ow.Passport.Number != "" {
		w.Passport.Number = ow.Passport.Number
	}
	if ow.Department.Id != 0 {
		w.Department.Id = ow.Department.Id
	}
	if ow.Department.Name != "" {
		w.Department.Name = ow.Department.Name
	}
	if ow.Department.Phone != "" {
		w.Department.Phone = ow.Department.Phone
	}
}

func Return_json(w http.ResponseWriter, jsonResponse []byte, jsonError error) {

	if jsonError != nil {
		jsonResponse, _ := json.Marshal(Response{"JSON convert error"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
