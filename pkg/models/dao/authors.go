package dao

import (
	"database/sql"
	"github.com/Denislite/library_app/pkg/models"
)

type Model struct {
	DB *sql.DB
}

func (m *Model) InsertAuthor(surname, name, middle_name, image_path string) error {
	req := "INSERT INTO authors (surname, name, middle_name, image_path) VALUES(?,?,?,?)"
	_, err := m.DB.Exec(req, surname, name, middle_name, image_path)
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
