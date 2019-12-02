package mysql

import (
	"database/sql"
	"github.com/joshuaherrera/snippetbox/pkg/models"
)

// SnippetModel defines wrapper for sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert new snippet into db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Exec takes sql statement and params to pass into it.
	// returns sql.Result object which contains info about
	// what happened on execution
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// get id of newly inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	// id is of type int64 so convert to int before return
	return int(id), nil
}

// Get snippet with id from db
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// returns pointer to sql.Row object that has res from db
	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}
	// reads into our snippet object. must past pointers and give exactly
	// same number of columns as arguments
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// Latest retrieves 10 most recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// defer close after checking for errors to avoid a panic closing a
	// nil resultset. Needed to make sure connections remain in pool
	defer rows.Close()

	snippets := []*models.Snippet{}

	// use rows.Next to iterate through rows in the resultset so
	// we can utilize rows.Scan on each record found
	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// must check for errors encountered in the iteration. Don't
	// assume successful iteration over resultset
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
