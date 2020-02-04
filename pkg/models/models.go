package models

import (
	"database/sql"
	"errors"
	"time"
)

// ErrNoRecord returns error if no models found
var ErrNoRecord = errors.New("models: no matching record found")

//Snippet defines structure of each snippet in db
type Snippet struct {
	ID      int
	Title   sql.NullString //example of struct value that can accept NULL
	Content string
	Created time.Time
	Expires time.Time
}
