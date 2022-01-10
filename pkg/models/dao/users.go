package dao

import (
	"database/sql"
	"errors"
	"github.com/Denislite/library_app/pkg/models"
)

func (m *Model) InsertUser(surname, name, middleName, passportData, birthdayDate, email, address string) error {
	req := "INSERT INTO users (surname, name, middle_name, passport_data, birthday_date, email, address) VALUES(?,?,?,?,?,?,?)"
	_, err := m.DB.Exec(req, surname, name, middleName, passportData, birthdayDate, email, address)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetAllUsers() ([]*models.User, error) {
	req := "SELECT * FROM users"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Surname, &user.Name, &user.MiddleName, &user.PassportData, &user.BirthdayDate, &user.Email, &user.Address, &user.DutyCount, &user.BookCount)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *Model) GetUserByID(id int) (*models.User, []*models.Book, error) {
	req := "SELECT * FROM users WHERE id = ?"

	row := m.DB.QueryRow(req, id)
	user := &models.User{}

	err := row.Scan(&user.ID, &user.Surname, &user.Name, &user.MiddleName, &user.PassportData, &user.BirthdayDate, &user.Email, &user.Address, &user.DutyCount, &user.BookCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, err
		} else {
			return nil, nil, err
		}
	}

	req = "SELECT * FROM books WHERE id IN (SELECT book_id FROM usersBooks WHERE user_id = ? AND returned = FALSE)"
	rows, err := m.DB.Query(req, id)

	defer rows.Close()

	var books []*models.Book

	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.AltName, &book.Genre, &book.Price, &book.Count, &book.ImageLink, &book.PricePerDay,
			&book.Year, &book.RegDate, &book.Rating)
		if err != nil {
			return nil, nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return user, books, nil
}
