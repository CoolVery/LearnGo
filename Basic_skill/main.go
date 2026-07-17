package main

import (
    "time"
    "fmt"
    "math"
)

func main() {
    var dateTime string
    var firstName, secondName, patronymic string
    var sumWorkFirst, sumWorkSecond, sumWorkThird float64
    fmt.Scanln(&dateTime)
    fmt.Scanln(&firstName)
    fmt.Scanln(&secondName)
    fmt.Scanln(&patronymic)
    fmt.Scanln(&sumWorkFirst)
    fmt.Scanln(&sumWorkSecond)
    fmt.Scanln(&sumWorkThird)
    dateTimeConverted, err := time.Parse("02.01.2006", dateTime)
    if err == nil {
     
    
    startWorkTime := dateTimeConverted.Add(time.Hour * 24 * 15)
    sumWork := sumWorkFirst + sumWorkSecond + sumWorkThird
    totalKopecksFloat := math.Round(sumWork * 100)
	totalKopecks := int(totalKopecksFloat)

	rubles := totalKopecks / 100
	kopecks := totalKopecks % 100
    fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %d руб. %d коп.\n\nС уважением,\nГл. бух. Иванов А.Е.", secondName, firstName, patronymic, startWorkTime.Format("02.01.2006"), rubles, kopecks)
    }
}