package form

import (
	"fmt"
	"regexp"

	"github.com/JesusIslam/sikritklab/internal/constant"
)

var (
	TagsRegExp = regexp.MustCompile(`/[a-zA-Z0-9]{2,32}/`)
)

type Thread struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (t *Thread) Validate() (err error) {
	if len(t.Title) < 1 || len(t.Title) > 128 {
		err = fmt.Errorf(constant.WarningInvalidTitle)
	}

	if len(t.Content) < 1 || len(t.Content) > 2000 {
		err = fmt.Errorf(constant.WarningInvalidContent)
	}

	if len(t.Tags) < 1 {
		err = fmt.Errorf(constant.WarningInvalidTextEmpty)
	}

	for _, tag := range t.Tags {
		if !TagsRegExp.MatchString(tag) {
			err = fmt.Errorf(constant.WarningInvalidTextAlphanumerics)
			break
		}
	}

	return err
}
