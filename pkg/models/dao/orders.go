package dao

import (
	"errors"
	"fmt"
	"github.com/Denislite/library_app/pkg/models"
)

func (m *Model) GiveBook(id, days int, book, returnDate string, discount float64) error {
	user, _, err := m.GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if user.BookCount == 5 {
		return errors.New("maximum book count")
	}

	req := `INSERT INTO usersBooks (user_id, book_id, give_date, return_date, default_price) SELECT ?, ?, 
            UTC_TIMESTAMP(), ?, books.price_per_day*?*? FROM books WHERE books.ID = ?`
	_, err = m.DB.Exec(req, id, book, returnDate, days, discount, book)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req = "UPDATE books SET count = count - 1 WHERE id = ?"
	_, err = m.DB.Exec(req, book)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req = `UPDATE users SET book_count = book_count + 1, discount = ?, 
           duty_count = (duty_count + (SELECT price_per_day FROM books WHERE id = ?)*?*users.discount) WHERE id = ?`
	_, err = m.DB.Exec(req, discount, book, days, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (m *Model) TakeBook(id int, book, rating, fine string, discount float64) error {
	req := `UPDATE users SET book_count = book_count - 1, duty_count = duty_count - (SELECT default_price FROM 
            usersBooks WHERE user_id = ? AND book_id = ? AND returned = FALSE), discount = ? WHERE id = ?`
	_, err := m.DB.Exec(req, id, book, discount, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req = `INSERT INTO orders (user_id, book_id, give_date, return_date, real_date, cost) SELECT user_id,
           book_id, give_date, return_date, UTC_TIMESTAMP(), default_price + ? + duty_count FROM usersBooks WHERE 
           user_id = ? AND book_id = ? AND returned = FALSE`
	_, err = m.DB.Exec(req, fine, id, book)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req = "UPDATE usersBooks SET returned = TRUE, rating = ? WHERE user_id = ? AND book_id = ? AND returned = FALSE"
	_, err = m.DB.Exec(req, rating, id, book)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req = "SELECT * FROM usersBooks WHERE book_id = ?"
	rows, err := m.DB.Query(req, book)

	if err != nil {
		return err
	}

	defer rows.Close()

	var userBooks []*models.UserBooks

	for rows.Next() {
		userBook := &models.UserBooks{}
		err = rows.Scan(&userBook.UserID, &userBook.BookID, &userBook.GiveDate, &userBook.ReturnDate,
			&userBook.DutyCount, &userBook.DefaultPrice, &userBook.Returned, &userBook.Rating)
		if err != nil {
			return err
		}
		userBooks = append(userBooks, userBook)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	var bookRating int

	count := 0
	for _, rating := range userBooks {
		bookRating += rating.Rating
		count += 1
	}
	bookRating = bookRating / count

	req = "UPDATE books SET rating = ?, count = count + 1 WHERE id = ?"
	_, err = m.DB.Exec(req, bookRating, book)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
