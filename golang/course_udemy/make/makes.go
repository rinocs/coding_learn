package main

import "fmt"

// type alias

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// use make

	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames = append(userNames, "Mario")
	userNames = append(userNames, "Maccio")
	fmt.Println(userNames)

	courseRatings := map[string]float64{}
	// memory reallocation every time we add an element
	courseRatings["english"] = 6.7
	courseRatings["go"] = 6.0
	courseRatings["test"] = 7

	fmt.Println(courseRatings)

	courseRatingsMake := make(map[string]float64, 3)

	courseRatingsMake["english"] = 6.7
	courseRatingsMake["go"] = 6.0
	courseRatingsMake["test"] = 7

	fmt.Println(courseRatingsMake)

	fMap := make(floatMap, 3)

	fMap["we"] = 1.2

	fMap.output()

	// loop on map

	for k, v := range courseRatingsMake {
		fmt.Println(k, v)
	}

}
