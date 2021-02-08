package leg

func getLeg(id int) *Leg {

	return &Leg{
		Id:          1,
		Origin:      "KISO",
		Destination: "KTEB",
	}
}

func getLegs() *[]Leg {

	return nil
}