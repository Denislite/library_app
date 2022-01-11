package models

type Book struct {
	ID          int
	Name        string
	AltName     string
	Genre       string
	Price       float64
	Count       int
	ImageLink   string
	PricePerDay float64
	Year        int
	RegDate     string
	Rating      int
}
