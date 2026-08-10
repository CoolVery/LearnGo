package classes

type Motorcycle struct {
	 Type (string)
	 Speed (float64)
	 FuelType (string)
}

func (motorcycle Motorcycle) CalculateTravelTime(distance float64) float64 {
		return distance / motorcycle.Speed

}