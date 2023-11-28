package puppy

import (
	"github.com/ChrisShia/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func From10() {
	fmt.Println("Im from 1.0.0")
}
