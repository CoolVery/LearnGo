package classes

type WeatherCondition int

const (
    Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
    Rain
    HeavyRain
    Snow
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}


func GetWeatherMultiplier(weather WeatherData) float64 {
	resultCoeff := 1.0
	switch weather.Condition {
	case Rain:
		resultCoeff += 0.125
	case HeavyRain:
		resultCoeff += 0.2
	case Snow:
		resultCoeff += 0.15
	}
	if weather.WindSpeed > 15 {
		resultCoeff += 0.1
	}
	return resultCoeff

}