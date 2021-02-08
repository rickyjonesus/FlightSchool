package leg

import (
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {

	leg := getLeg(1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(leg)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {
}

func Register() {

	http.HandleFunc("/Leg/Update", Update)

	http.HandleFunc("/Leg/Add", Add)

	http.HandleFunc("/Leg", Get)

}
