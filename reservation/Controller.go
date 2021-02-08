package reservation

import (
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {

	res := getReservation(1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(res)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {
}


func Register() {

	http.HandleFunc("/Reservation/Update", Add)

	http.HandleFunc("/Reservation/Add", Add)

	http.HandleFunc("/Reservation", Get)

}