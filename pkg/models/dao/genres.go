package dao

import "github.com/Denislite/library_app/pkg/models"

func (m *Model) GetAllGenres() ([]*models.Genre, error) {
	req := "SELECT * FROM genre"
	rows, err := m.DB.Query(req)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var genres []*models.Genre

	for rows.Next() {
		genre := &models.Genre{}
		err = rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}
