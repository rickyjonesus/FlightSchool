package main

import (
	"fmt"
	"net/http"

	"github.com/rickyjonesus/FlightSchool/aircraft"

	_ "github.com/lib/pq"
	//	"github.com/micro/go-micro/v2"
)

func main() {

	fmt.Println("Program Starting")
	//log.AddHook(logruseq.NewSeqHook("https://atlas-seq.azurewebsites.net"))

	// create a new service
	//service := micro.NewService(
	//micro.Name("helloworld"),
	//)

	// initialise flags
	//service.Init()

	//service.Run()

	http.HandleFunc("/Aircraft/Add", aircraft.Add)

	http.HandleFunc("/Aircraft", aircraft.Get)

	http.ListenAndServe(":80", nil)

}
