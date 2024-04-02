package models

import "suggestions.api/db"

type Suggestion struct {
	Id       int64  `json:"-"`
	Username string `json:"username"`
	Text     string `binding:"required" form:"text" json:"text"`
}

func (s *Suggestion) Save() error {
	query := "INSERT INTO suggestions(username, text) VALUES (?, ?)"

	stmt, err := db.Database.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(s.Username, s.Text)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	s.Id = id
	return err
}

func GetAllSuggestions() ([]Suggestion, error) {
	query := "SELECT * FROM suggestions"
	rows, err := db.Database.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suggestions []Suggestion

	for rows.Next() {
		var suggestion Suggestion
		err := rows.Scan(&suggestion.Id, &suggestion.Username, &suggestion.Text)
		if err != nil {
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}
