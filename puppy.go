package puppy

import (
	"fmt"

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

func From11() {
	fmt.Println("im from 1.1.0")
}
