package handler

import (
	"net/http"

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

	db := database.New()
	thread := &model.Thread{
		Title: threadForm.Title,
		Tags:  threadForm.Tags,
	}
	err = db.Table("threads").Create(thread).Error
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusServiceUnavailable, resp)
	}

	resp.Message = thread
	return c.JSON(http.StatusCreated, resp)
}
