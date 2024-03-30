package main

import (
	"fmt"
	"strconv"
)

func main() {
	var temperature string
	var converter string
	fmt.Print("Please Add a tempature followed with Converter Type:")
	fmt.Scanf("%s %s", &temperature, &converter)
	if temperature[len(temperature)-1:] == converter {
		fmt.Println("Already in Same Unit")
		return
	}
	var temparatureUnit int
	temparatureUnit, err := strconv.Atoi(temperature[:len(temperature)-1])
	if err != nil {
		fmt.Println("Please Add a Valid Temperature Value!")
		return
	}
	switch converter {
	case "C":
		var C int
		C = (temparatureUnit - 32) * 5 / 9
		fmt.Printf("Temperature Converted into Celsius: %d C", C)
	case "F":
		var F int
		F = (temparatureUnit * 9 / 5) + 32
		fmt.Printf("Temperature Converted into Farenheit: %d F", F)
	}
}
