package aircraft

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {

	ac := Aircraft{Id: 0, TailNumber: "N843AL", AircraftTypeId: 1}

	AddAircraft(ac)

	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func Get(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get aircraft")
	ac := GetAircraft("")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ac)

}
