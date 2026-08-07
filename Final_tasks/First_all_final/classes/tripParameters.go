package classes

type TripParameters struct {
	Distance float64
	Duration float64
}

const priceKm = 10.0
const pricePerMinute = 2.0

func (tripParameters *TripParameters) CalculateBasePrice() float64 {
	result := tripParameters.Distance * priceKm + tripParameters.Duration * pricePerMinute
	return result
}
