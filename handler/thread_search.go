package handler

import (
	"net/http"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/form"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

// Search by title, tag1-5
// tag cannot be wild card, title default to wildcard
// Default order by created_at desc
func ThreadSearch(c echo.Context) (err error) {
	resp := &response.Response{}

	search := form.GetSearchForm(c)

	// get the thread and first 1 post of each thread
	threadsWithFirstPost := []response.M{}

	tx := database.New().Begin()

	threads := []*model.Thread{}
	err = tx.Table("threads").Where("MATCH (tags) AGAINTS (?)", search.Tags).Find(&threads).Error
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusServiceUnavailable, resp)
	}

	for _, thread := range threads {
		post := &model.Post{}
		err = tx.Table("posts").Where("thread_id = ?", thread.ID).First(post).Error
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusServiceUnavailable, resp)
		}
		if err == database.NotFound {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusNotFound, resp)
		}

		threadsWithFirstPost = append(threadsWithFirstPost, response.M{
			"thread":     thread,
			"first_post": post,
		})
	}

	tx.Commit()

	resp.Message = threadsWithFirstPost
	return c.JSON(http.StatusOK, resp)
}
