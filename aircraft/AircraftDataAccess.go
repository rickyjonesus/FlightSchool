package aircraft

import (
	"fmt"
	"os"

	"github.com/randallmlough/pgxscan"
	database "github.com/rickyjonesus/FlightSchool/database"
)

//AddAircraft ... Adds an aircraft to the database
func AddAircraft(aircraftToAdd Aircraft) int {
	// psqlInfo := Database.GetConnectionInfo()

	// conn, err := pgx.Connect(context.Background(), psqlInfo.host)
	// if err ! nil {
	/// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background() )

	// err = conn.(context.Background(),addAircraftSQL).can(&name, &weight)
	// if err = nil {
	// 		fmt.Fprintf(os.Stderr, "Queryow failed: %v\n", err)
	// 	os.Exit(1)

	// sqlStatement := fmt.Sprintf(adAircraftSQL,
	// 	aircraftToAdd.TailNumber,
	// 	aircraftToAdd.AircraftTypId)

	// 	fmt.Println(sqlStatement)

	// _, err = db.ExecsqlStatement)

	// i err != nil {
	// panic(err)

	return 1
}

//
func GetAircraftByTail(tailNumber string) Aircraft {

	var retAircraft Aircraft

	conn := database.GetConnection()
	defer conn.Close()

	row := conn.QueryRow(getAircraftByTailSQL, tailNumber)
	//.Scan(&retAircraft.Id, &retAircraft.TailNumber, &retAircraft.AircraftTypeId)
	pgxscan.NewScanner(row).Scan(retAircraft)
	

	fmt.Println(retAircraft.Id, retAircraft.TailNumber)

	return retAircraft
}

func GetAircraft() *[]Aircraft {

	conn := database.GetConnection()
	defer conn.Close()

	rows, err := conn.Query(getAllAircraftSQL)

	var retAircraft []Aircraft


	if err := pgxscan.NewScanner(rows).Scan(&retAircraft); err != nil {
	
	}
	if err != nil {
	
		os.Exit(1)
	}

	return &retAircraft
}
