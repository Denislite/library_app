package dao

import (
	"database/sql"
	"errors"
	"github.com/Denislite/library_app/pkg/models"
)

func (m *Model) InsertBook(name, altName, imageLink, genre, price, count, pricePerDay, year string, authors []string) error {
	req := `INSERT INTO books (name, alt_name, genre, price, count, image_path, price_per_day, year, reg_date) 
			VALUES(?,?,?,?,?,?,?,?,UTC_TIMESTAMP())`
	result, err := m.DB.Exec(req, name, altName, genre, price, count, imageLink, pricePerDay, year)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, value := range authors {
		req = "INSERT INTO bookAuthors (book_id, author_id) VALUES(?, ?)"
		_, err := m.DB.Exec(req, id, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) GetAllBooks() ([]*models.Book, error) {
	req := "SELECT * FROM books"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*models.Book

	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.AltName, &book.Genre, &book.Price, &book.Count, &book.ImageLink,
			&book.PricePerDay, &book.Year, &book.RegDate, &book.Rating)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (m *Model) GetBookByID(id int) (*models.Book, []*models.Author, error) {
	req := "SELECT * FROM books WHERE id = ?"

	row := m.DB.QueryRow(req, id)
	book := &models.Book{}
	err := row.Scan(&book.ID, &book.Name, &book.AltName, &book.Genre, &book.Price, &book.Count, &book.ImageLink,
		&book.PricePerDay, &book.Year, &book.RegDate, &book.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, err
		} else {
			return nil, nil, err
		}
	}

	req = "SELECT * FROM authors WHERE id IN (SELECT author_id FROM bookAuthors WHERE book_id = ?)"
	rows, err := m.DB.Query(req, id)

	defer rows.Close()

	var authors []*models.Author

	for rows.Next() {
		author := &models.Author{}
		err = rows.Scan(&author.ID, &author.Surname, &author.Name, &author.MiddleName, &author.ImageLink)
		if err != nil {
			return nil, nil, err
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return book, authors, nil
}

func (m *Model) GetTopBooks() ([]*models.Book, error) {
	req := "SELECT * FROM books ORDER BY rating DESC LIMIT 3"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*models.Book

	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.AltName, &book.Genre, &book.Price, &book.Count, &book.ImageLink,
			&book.PricePerDay, &book.Year, &book.RegDate, &book.Rating)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (m *Model) GetAvailableBooks(id int) ([]*models.Book, error) {
	req := `SELECT * FROM books WHERE count > 0 AND id NOT IN (SELECT book_id FROM usersBooks where user_id = ? 
            AND returned = FALSE)`
	rows, err := m.DB.Query(req, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*models.Book

	for rows.Next() {
		book := &models.Book{}
		err = rows.Scan(&book.ID, &book.Name, &book.AltName, &book.Genre, &book.Price, &book.Count, &book.ImageLink,
			&book.PricePerDay, &book.Year, &book.RegDate, &book.Rating)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
