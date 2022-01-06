package models

type Book struct {
	ID          int
	Name        string
	AltName     string
	Genre       string
	Price       int
	Count       int
	ImageLink   string
	PricePerDay int
	Year        int
	RegDate     string
	Rating      int
}
