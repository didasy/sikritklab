package handler

import (
	"net/http"

	"github.com/JesusIslam/sikritklab/internal/database"
	"github.com/JesusIslam/sikritklab/internal/model"
	"github.com/JesusIslam/sikritklab/internal/response"
	"github.com/asdine/storm/q"
	"github.com/gin-gonic/gin"
)

func ThreadGetByID(c *gin.Context) {
	resp := &response.Response{}

	threadID := c.Param("id")

	tx, err := database.DB.Begin(false)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// get the thread
	thread := &model.Thread{}
	err = tx.One("ID", threadID, thread)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// then get the first 500 posts sorted by created at oldest at top
	posts := []*model.Post{}
	err = tx.Select(
		q.Eq("ThreadID", threadID),
	).OrderBy("CreatedAt").Limit(500).Find(&posts)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// get the tags too
	tags := []*model.Tag{}
	err = tx.Find("ThreadID", threadID, &tags)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	tx.Commit()

	resp.Message = &model.ThreadPost{
		Thread: thread,
		Posts:  posts,
		Tags:   tags,
	}
	c.JSON(http.StatusOK, resp)
}
