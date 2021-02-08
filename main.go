package main

import (
	"fmt"
	"net/http"

	"log"

	_ "github.com/lib/pq"
	"github.com/rickyjonesus/FlightSchool/aircraft"
	"github.com/rickyjonesus/FlightSchool/leg"
	"github.com/rickyjonesus/FlightSchool/reservation"
	//	"github.com/micro/go-micro/v2"
)

func main() {

	fmt.Println("Program Starting")

	aircraft.Register()

	leg.Register()

	reservation.Register()

	log.Fatal(http.ListenAndServe(":81", nil))

}
