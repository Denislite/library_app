package models

type Order struct {
	ID         int
	UserId     int
	BookId     int
	GiveDate   string
	ReturnDate string
	RealDate   string
	Cost       float64
}
