package handler

import (
	"net/http"
	"strconv"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

func ThreadGetByID(c echo.Context) (err error) {
	resp := &response.Response{}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	tx := database.New().Begin()

	// get the thread
	thread := &model.Thread{}
	err = tx.Table("threads").Where("id = ?", id).First(thread).Error
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

	// then get the first 500 posts
	posts := []*model.Post{}
	err = tx.Table("posts").Where("thread_id = ?", id).Order(database.ThreadDisplayOrder).Limit(database.PostLimitPerThread).Find(&posts).Error
	if err == database.NotFound {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusNotFound, resp)
	}
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusServiceUnavailable, resp)
	}
	tx.Commit()

	resp.Message = response.M{
		"thread": thread,
		"posts":  posts,
	}
	return c.JSON(http.StatusOK, resp)
}
