package form

import (
	"strings"

	"regexp"
	"strconv"

	"github.com/labstack/echo"
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

func GetSearchForm(c echo.Context) *Search {
	search := &Search{
		Page:           0,
		PerPage:        10,
		OrderBy:        "CreatedAt",
		OrderDirection: "reverse",
	}

	if perPage := c.QueryParam("per_page"); perPage != "" {
		p, err := strconv.ParseInt(perPage, 10, 32)
		if err == nil {
			search.PerPage = int(p)
		}
	}

	if page := c.QueryParam("page"); page != "" {
		p, err := strconv.ParseInt(page, 10, 32)
		if err == nil {
			search.Page = int(p) * search.PerPage
		}
	}

	if tagsStr := c.QueryParam("tags"); tagsStr != "" {
		tagsStr = strings.TrimSpace(tagsStr)
		tags := strings.Split(tagsStr, ",")

		for _, tag := range tags {
			tag = SanitizeTagRegExp.ReplaceAllString(tag, "")
			if TagsRegExp.MatchString(tag) {
				search.Tags = append(search.Tags, tag)
			}
		}
	}

	if title := c.QueryParam("title"); title != "" {
		search.Title = title
	}

	// if both search exists, title takes precedence
	if search.Title != "" && len(search.Tags) > 0 {
		search.Tags = nil
	}

	return search
}
