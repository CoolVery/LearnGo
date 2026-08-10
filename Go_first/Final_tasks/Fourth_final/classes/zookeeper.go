package classes

import (
	"fmt"
	. "github.com/CoolVery/LearnGo.git/interfaces"
)
type ZooKeeper struct {

}

func (zooKeeper ZooKeeper)Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!.\n", animal.GetName(), animal.MakeSound())
}