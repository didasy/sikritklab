package form

import (
	"fmt"

	"github.com/JesusIslam/sikritklab/internal/constant"
	"github.com/mvdan/xurls"
)

type Post struct {
	Title   string `json:"title,omitempty"`
	Image   string `json:"image,omitempty"`
	Content string `json:"content"`
}

func (p *Post) Validate() (err error) {
	if len(p.Image) > 0 {
		if len(p.Image) > 1024 {
			err = fmt.Errorf(constant.WarningInvalidImage)
			return
		}
		if !xurls.Strict().MatchString(p.Image) {
			err = fmt.Errorf(constant.WarningInvalidImage)
		}
	}

	if len(p.Content) < 1 || len(p.Content) > 2000 {
		err = fmt.Errorf(constant.WarningInvalidContent)
	}

	if len(p.Title) > 128 {
		err = fmt.Errorf(constant.WarningInvalidTitle)
	}

	return err
}
