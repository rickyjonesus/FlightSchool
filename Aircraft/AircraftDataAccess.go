package Aircraft

import (
	"database/sql"

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

	sqlStatement := "INSERT INTO \"Aircraft\" (\"Id\",\"TailNumber\") VALUES (5, 'N3941FL')"

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}
