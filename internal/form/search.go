package form

import (
	"strings"

	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	SanitizeTagRegExp = regexp.MustCompile(`/\r\s\t\n/`)
)

type Search struct {
	Page           int
	PerPage        int
	OrderBy        string // unchangeable
	OrderDirection string // unchangeable
	Title          string
	Tags           []string
}

func GetSearchForm(c *gin.Context) *Search {
	search := &Search{
		Page:           0,
		PerPage:        10,
		OrderBy:        "CreatedAt",
		OrderDirection: "reverse",
	}

	if perPage := c.Query("per_page"); perPage != "" {
		p, err := strconv.ParseInt(perPage, 10, 32)
		if err == nil {
			search.PerPage = int(p)
		}
	}

	if page := c.Query("page"); page != "" {
		p, err := strconv.ParseInt(page, 10, 32)
		if err == nil {
			search.Page = int(p) * search.PerPage
		}
	}

	if tagsStr := c.Query("tags"); tagsStr != "" {
		tagsStr = strings.TrimSpace(tagsStr)
		tags := strings.Split(tagsStr, ",")

		for _, tag := range tags {
			tag = SanitizeTagRegExp.ReplaceAllString(tag, "")
			if TagsRegExp.MatchString(tag) {
				search.Tags = append(search.Tags, tag)
			}
		}
	}

	if title := c.Query("title"); title != "" {
		search.Title = title
	}

	// if both search exists, title takes precedence
	if search.Title != "" && len(search.Tags) > 0 {
		search.Tags = nil
	}

	return search
}
