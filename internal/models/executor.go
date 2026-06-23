package models

type ExecutorModel struct {
	ID          int64
	Coordinates *Point
}

type Point struct {
	Lattitude float64
	Longitude float64
}
