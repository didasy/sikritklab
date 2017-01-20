package form

import (
	"strings"

	"strconv"

	"github.com/labstack/echo"
)

type Search struct {
	Page    int
	PerPage int
	OrderBy string
	Title   string
	Tags    string
}

func GetSearchForm(c echo.Context) *Search {
	search := &Search{
		Page:    0,
		PerPage: 10,
		OrderBy: "created_at DESC",
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

	if tags := c.QueryParam("tags"); tags != "" {
		tags = strings.TrimSpace(tags)
		if TagsRegExp.MatchString(tags) {
			search.Tags = tags
		}
	}

	if title := c.QueryParam("title"); title != "" {
		search.Title = title
	}

	return search
}
