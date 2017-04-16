package form

import (
	"fmt"
	"regexp"
)

var (
	TagsRegExp = regexp.MustCompile(`\[a-zA-Z0-9]{2,32}\`)
)

type Thread struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (t *Thread) Validate() (err error) {
	if len(t.Title) < 1 || len(t.Title) > 128 {
		err = fmt.Errorf("Invalid title: must be between 1 and 128 characters long")
	}

	if len(t.Content) < 1 || len(t.Content) > 2000 {
		err = fmt.Errorf("Invalid content: must be between 1 and 2000 characters long")
	}

	if len(t.Tags) < 1 {
		err = fmt.Errorf("Invalid tags: must not be empty")
	}

	for _, tag := range t.Tags {
		if !TagsRegExp.MatchString(tag) {
			err = fmt.Errorf("Invalid tags: must be alphanumerics only")
			break
		}
	}

	return err
}
