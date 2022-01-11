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
	Discount     float64
	DutyCount    float64
	BookCount    int
}
