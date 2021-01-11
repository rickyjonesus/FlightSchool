package aircraft

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
	"github.com/rickyjonesus/FlightSchool/Database"
)

//AddAircraft ... Adds an aircraft to the database
func AddAircraft(aircraftToAdd Aircraft) {
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

}

//
func GetAircraft(tailNumbr string) Aircraft {
	psqlInfo := Database.GetConnectionObject()

	conn, err := pgx.Connect(psqlInfo)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()

	var retAircraft Aircraft

	err = conn.QueryRow(getAllAircraftSQL).Scan(&retAircraft.Id, &retAircraft.TailNumber, &retAircraft.AircraftTypeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(retAircraft.Id, retAircraft.TailNumber)

	return retAircraft
}
