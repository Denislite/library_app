package models

type UserBooks struct {
	UserID       int
	BookID       int
	GiveDate     string
	ReturnDate   string
	DefaultPrice string
	DutyCount    string
	Returned     bool
	Rating       int
}
