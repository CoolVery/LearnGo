package classes

type Car struct {
	 Type (string)
	 Speed (float64)
	 FuelType (string)
}

func (car Car) CalculateTravelTime(distance float64) float64 {
	
	return distance / car.Speed
}