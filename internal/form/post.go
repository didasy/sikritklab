package form

import (
	"fmt"


	"github.com/JesusIslam/sikritklab/internal/constant"
)

type Post struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content"`
}

func (p *Post) Validate() (err error) {
	if len(p.Content) < 1 || len(p.Content) > 2000 {
		err = fmt.Errorf(constant.WarningInvalidContent)
	}

	if len(p.Title) > 128 {
		err = fmt.Errorf(constant.WarningInvalidTitle)
	}

	return err
}
