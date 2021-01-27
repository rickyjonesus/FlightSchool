package aircraft

const addAircraftSQL = "INSERT INTO \"Aircraft\" (\"TailNumber\",\"AircraftTypeId\") VALUES ('%v', %v)"
const getAllAircraftSQL = "select * from \"Aircraft\""

const getAircraftByTailSQL = "select * from \"Aircraft\" where \"TailNumber\"=$1"
