package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hola Mundo")

	start := time.Now()

	fmt.Println(start.Weekday())

}
