package models

type UserBooks struct {
	UserID       int
	BookID       int
	ReturnDate   string
	DefaultPrice string
	Returned     bool
	Rating       int
}
