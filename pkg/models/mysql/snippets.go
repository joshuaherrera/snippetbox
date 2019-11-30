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
	return nil, nil
}
// Latest retrieves 10 most recent snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}