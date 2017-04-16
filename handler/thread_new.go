package handler

import (
	"net/http"
	"time"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/form"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

func ThreadNew(c echo.Context) (err error) {
	resp := &response.Response{}

	threadForm := &form.Thread{}
	err = c.Bind(threadForm)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = threadForm.Validate()
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	now := time.Now()

	threadID, err := database.NewThreadID()
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
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
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = tx.Save(thread)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = tx.Save(post)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	for _, tag := range tags {
		err = tx.Save(tag)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}
	}

	tx.Commit()

	resp.Message = thread
	return c.JSON(http.StatusCreated, resp)
}
