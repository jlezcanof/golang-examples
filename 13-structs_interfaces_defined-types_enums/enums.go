package main

import "fmt"

func main() {
	//day := Friday
	//fmt.Println(Friday.String())
	day := Monday
	fmt.Printf(day.String())
}

type Weekday int

const (
	Monday    Weekday = 0
	Tuesday   Weekday = 1
	Wednesday Weekday = 2
	Thursday  Weekday = 3
	Friday    Weekday = 4
	Saturday  Weekday = 5
	Sunday    Weekday = 6
)

func (d Weekday) String() string {
	names := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	return names[d]
}
