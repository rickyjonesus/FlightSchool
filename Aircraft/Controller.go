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

	tailNumber := r.URL.Query().Get("tailNumber")

	w.Header().Set("Content-Type", "applicatio/json")
	w.WriteHeader(http.StatusCreated)
	
	if(tailNumber == ""){
		fmt.Println("Get all aircraft")	
		ac  := GetAircraft()
		json.NewEncoder(w).Encode(ac)
	} else {
		ac := GetAircraftByTail(tailNumber)
		json.NewEncoder(w).Encode(ac)
	}
}



func Register() {

	http.HandleFunc("/Aircraft/Update", Add)

	http.HandleFunc("/Aircraft/Add", Add)

	http.HandleFunc("/Aircraft", Get)

}
