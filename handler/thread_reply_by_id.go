package handler

import (
	"net/http"
	"time"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/form"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/asdine/storm/q"
	"github.com/labstack/echo"
)

func ThreadReplyByID(c echo.Context) (err error) {
	resp := &response.Response{}

	threadID := c.Param("id")

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

	tx, err := database.DB.Begin(true)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// trim expired posts (if more than 500 in this thread)
	// first get them all with newest on top
	posts := []*model.Post{}
	err = tx.Select(
		q.Eq("ThreadID", threadID),
	).OrderBy("CreatedAt").Reverse().Limit(500).Find(&posts)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// then delete them one by one (the oldest) if more than 500 exists
	// make it 499 posts max
	if len(posts) > 500 {
		for _, p := range posts[499:] {
			err = tx.DeleteStruct(p)
			if err != nil {
				tx.Rollback()
				resp.Error = err.Error()
				return c.JSON(http.StatusInternalServerError, resp)
			}
		}
	}

	// then create the post
	post := &model.Post{
		CreatedAt: time.Now(),
		Title:     postForm.Title,
		Content:   postForm.Content,
		ThreadID:  threadID,
	}

	err = tx.Save(post)
	if err != nil {
		tx.Rollback()
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	tx.Commit()

	resp.Message = post
	return c.JSON(http.StatusCreated, resp)
}
