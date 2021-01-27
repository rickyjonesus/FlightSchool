package Database

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

const (
	host     = "192.168.1.111"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "FlightSchool"
)

func GetConnectionObject() pgx.ConnConfig {
	var ret pgx.ConnConfig

	ret.Host = host
	ret.Port = port
	ret.User = user
	ret.Password = password
	ret.Database = dbname

	return ret

}

//
//
//Export
func GetConnectionInfo() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlInfo
}

//
func GetConnection() *pgx.Conn {
	psqlInfo := GetConnectionObject()

	conn, err := pgx.Connect(psqlInfo)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}



	return conn
}
