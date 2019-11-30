package models

import (
	"errors"
	"time"
)

// ErrNoRecord returns error if no models found
var ErrNoRecord = errors.New("models: no matching record found")

//Snippet defines structure of each snippet in db
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
