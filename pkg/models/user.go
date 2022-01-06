package models

type User struct {
	ID           int
	Surname      string
	Name         string
	MiddleName   string
	PassportData string
	BirthdayDate string
	Email        string
	Address      string
	DutyCount    int
	BookCount    int
}
