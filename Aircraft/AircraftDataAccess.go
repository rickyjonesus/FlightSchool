package Aircraft

import (
	"database/sql"
	"fmt"

	"github.com/rickyjonesus/FlightSchool/Database"
)

func AddAircraft(aircraftToAdd Aircraft) {

	psqlInfo := Database.GetConnectionInfo()

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := fmt.Sprintf("INSERT INTO \"Aircraft\" (\"TailNumber\",\"AircraftTypeId\") VALUES ('%v', %v)",
		aircraftToAdd.TailNumber,
		aircraftToAdd.AircraftTypeId)

	fmt.Println(sqlStatement)

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}
