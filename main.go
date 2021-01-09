package main

import (
	"fmt"
	//"net/http"

	_ "github.com/lib/pq"
	//	log "github.com/sirupsen/logrus"
	"github.com/rickyjonesus/FlightSchool/Aircraft"
)

func main() {

	fmt.Println("Program Starting")
	//log.AddHook(logruseq.NewSeqHook("https://atlas-seq.azurewebsites.net"))

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	ac := Aircraft.Aircraft{0, "N843AL"}

	Aircraft.AddAircraft(ac)

	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	//})

	//http.ListenAndServe(":80", nil)

}
