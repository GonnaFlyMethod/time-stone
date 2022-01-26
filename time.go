package main

type Time struct {
	hours, seconds, minutes int
}

func CreateNewTimeObj(hours, seconds, minutes int) *Time {
	return &Time{hours, seconds, minutes}
}
