package goserbench

import "time"

//go:generate ffjson structdef-ffjson.go
type FFJSONA struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}
