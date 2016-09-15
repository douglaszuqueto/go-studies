package Domains

import (
	"encoding/json"
	"net/http"

	Repository "../../../Applications/Users/Repositories"

	"github.com/gorilla/mux"
)

func Index(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(Repository.All()); err != nil {
		panic(err)
	}
}

func Show(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	println("User ID:", userId)
	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(Repository.GetById(userId)); err != nil {
		panic(err)
	}
}

func Store(res http.ResponseWriter, req *http.Request) {
	println("Store User");
}

func Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	println("Update User:", userId)
}

func Remove(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	println("Delete User:", userId)
}
