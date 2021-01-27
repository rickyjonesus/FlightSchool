package main

import (
	"fmt"
	"net/http"

	"log"

	_ "github.com/lib/pq"
	"github.com/rickyjonesus/FlightSchool/aircraft"
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

	aircraft.Register()
	//http.HandleFunc("/Aircraft/Update", aircraft.Add)

	//http.HandleFunc("/Aircraft/Add", aircraft.Add)

	//http.HandleFunc("/Aircraft", aircraft.Get)

	log.Fatal(http.ListenAndServe(":81", nil))

}
