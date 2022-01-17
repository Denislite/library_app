package dao

import (
	"fmt"
	"github.com/Denislite/library_app/pkg/models"
)

func (m *Model) GiveBook(id, days int, book, returnDate string, discount float64) error {
	req := `INSERT INTO usersBooks (user_id, book_id, give_date, return_date, default_price) SELECT ?, ?, 
            UTC_TIMESTAMP(), ?, books.price_per_day*?*? FROM books WHERE books.ID = ?`
	_, err := m.DB.Exec(req, id, book, returnDate, days, discount, book)
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
			&userBook.DefaultPrice, &userBook.DutyCount, &userBook.Returned, &userBook.Rating)
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

func (m *Model) UpdateDuty() error {
	req := "UPDATE usersBooks SET duty_count = duty_count + default_price*0.01 WHERE returned = FALSE AND return_date < UTC_TIMESTAMP()"
	_, err := m.DB.Exec(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (m *Model) GetUsersWithDuty() ([]*models.User, error) {
	req := "SELECT * FROM users WHERE (SELECT user_id FROM usersBooks WHERE return_date < UTC_TIMESTAMP())"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Surname, &user.Name, &user.MiddleName, &user.PassportData, &user.BirthdayDate,
			&user.Email, &user.Address, &user.Discount, &user.DutyCount, &user.BookCount)
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

func (m *Model) GetClosedOrders() ([]*models.Order, error) {
	req := "SELECT * FROM orders"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []*models.Order

	for rows.Next() {
		order := &models.Order{}
		err = rows.Scan(&order.ID, &order.UserId, &order.BookId, &order.GiveDate, &order.ReturnDate, &order.RealDate, &order.Cost)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (m *Model) GetActiveOrders() ([]*models.UserBooks, error) {
	req := "SELECT * FROM usersBooks WHERE returned = FALSE"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userBooks []*models.UserBooks

	for rows.Next() {
		userBook := &models.UserBooks{}
		err = rows.Scan(&userBook.UserID, &userBook.BookID, &userBook.GiveDate, &userBook.ReturnDate,
			&userBook.DefaultPrice, &userBook.DutyCount, &userBook.Returned, &userBook.Rating)
		if err != nil {
			return nil, err
		}
		userBooks = append(userBooks, userBook)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userBooks, nil
}
