package form

import (
	"fmt"
)

type Post struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content"`
}

func (p *Post) Validate() (err error) {
	if len(p.Content) < 1 || len(p.Content) > 2000 {
		err = fmt.Errorf("Invalid content: must be between 1 and 2000 characters long")
	}

	if len(p.Title) > 128 {
		err = fmt.Errorf("Invalid title: cannot be more than 128 characters long")
	}

	return err
}
