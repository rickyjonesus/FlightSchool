package Database

import "fmt"

const (
	host     = "192.168.1.111"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "FlightSchool"
)

//
//
//Export
func GetConnectionInfo() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlInfo
}
