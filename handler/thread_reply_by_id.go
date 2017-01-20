package handler

import (
	"net/http"
	"strconv"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/form"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

func ThreadReplyByID(c echo.Context) (err error) {
	resp := &response.Response{}

	threadID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	postForm := &form.Post{}
	err = c.Bind(postForm)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = postForm.Validate()
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	}

	db := database.New()
	post := &model.Post{
		Title:    postForm.Title,
		Content:  postForm.Content,
		ThreadID: threadID,
	}
	err = db.Table("posts").Create(post).Error
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusServiceUnavailable, resp)
	}

    resp.Message = post
	return c.JSON(http.StatusCreated, resp)
}
