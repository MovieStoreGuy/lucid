package main

import (
	"fmt"

	"github.com/MovieStoreGuy/lucid/services"
)

func main() {
	fmt.Println("Getting registered services: ", services.GetFactoryInstance())
}
