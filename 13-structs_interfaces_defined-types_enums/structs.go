package main

import "time"

var page struct {
	number int
	header string
	text   string
}

var book struct {
	title       string
	autor       string
	releaseDate *time.Time
	pages       []struct {
		number int
		header string
		text   string
	}
}
