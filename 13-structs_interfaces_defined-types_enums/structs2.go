package main

import "time"

type page2 struct {
	number int
	header string
	text   string
}

type book2 struct {
	title       string
	autor       string
	releaseDate *time.Time
	pages       []page2
}
