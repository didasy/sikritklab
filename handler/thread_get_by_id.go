package handler

import (
	"net/http"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/asdine/storm/q"
	"github.com/labstack/echo"
)

func ThreadGetByID(c echo.Context) (err error) {
	resp := &response.Response{}

	threadID := c.Param("id")

	tx, err := database.DB.Begin(false)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// get the thread
	thread := &model.Thread{}
	err = tx.One("ID", threadID, thread)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// then get the first 500 posts sorted by created at oldest at top
	posts := []*model.Post{}
	err = tx.Select(
		q.Eq("ThreadID", threadID),
	).OrderBy("CreatedAt").Limit(500).Find(&posts)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// get the tags too
	tags := []*model.Tag{}
	err = tx.Find("ThreadID", threadID, &tags)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	tx.Commit()

	resp.Message = &model.ThreadPost{
		Thread: thread,
		Posts:  posts,
		Tags:   tags,
	}
	return c.JSON(http.StatusOK, resp)
}
