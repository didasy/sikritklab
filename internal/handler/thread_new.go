package handler

import (
	"net/http"
	"time"

	"github.com/JesusIslam/sikritklab/internal/database"
	"github.com/JesusIslam/sikritklab/internal/form"
	"github.com/JesusIslam/sikritklab/internal/model"
	"github.com/JesusIslam/sikritklab/internal/response"
	"github.com/gin-gonic/gin"
)

func ThreadNew(c *gin.Context) {
	var err error
	resp := &response.Response{}

	threadForm := &form.Thread{}
	err = c.Bind(threadForm)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err = threadForm.Validate()
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	now := time.Now()

	threadID, err := database.NewThreadID()
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// create the thread
	thread := &model.Thread{
		ID:        threadID,
		CreatedAt: now,
		Title:     threadForm.Title,
	}

	// and the first post
	post := &model.Post{
		CreatedAt: now,
		ThreadID:  threadID,
		Content:   threadForm.Content,
	}

	// and the tags
	tags := []*model.Tag{}
	for _, tag := range threadForm.Tags {
		tags = append(tags, &model.Tag{
			CreatedAt: now,
			Tag:       tag,
			ThreadID:  threadID,
		})
	}

	// save them all
	tx, err := database.DB.Begin(true)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	err = tx.Save(thread)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	err = tx.Save(post)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, tag := range tags {
		err = tx.Save(tag)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
	}

	tx.Commit()

	resp.Message = thread
	c.JSON(http.StatusCreated, resp)
}
