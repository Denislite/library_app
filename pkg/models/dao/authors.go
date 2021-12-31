package dao

import (
	"database/sql"
	"errors"
	"github.com/Denislite/library_app/pkg/models"
)

type Model struct {
	DB *sql.DB
}

func (m *Model) InsertAuthor(surname, name, middleName, imagePath string) error {
	req := "INSERT INTO authors (surname, name, middle_name, image_path) VALUES(?,?,?,?)"
	_, err := m.DB.Exec(req, surname, name, middleName, imagePath)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) GetAllAuthors() ([]*models.Author, error) {
	req := "SELECT * FROM authors"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var authors []*models.Author

	for rows.Next() {
		author := &models.Author{}
		err = rows.Scan(&author.ID, &author.Surname, &author.Name, &author.Middle_name, &author.Image_link)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (m *Model) GetAuthorByID(id int) (*models.Author, error) {
	req := "SELECT * FROM authors WHERE id = ?"

	row := m.DB.QueryRow(req, id)
	author := &models.Author{}

	err := row.Scan(&author.ID, &author.Surname, &author.Name, &author.Middle_name, &author.Image_link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		} else {
			return nil, err
		}
	}

	return author, nil
}
