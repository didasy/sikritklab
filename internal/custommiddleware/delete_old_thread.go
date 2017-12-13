package custommiddleware

import (
	"net/http"
	"time"

	"github.com/JesusIslam/sikritklab/internal/database"
	"github.com/JesusIslam/sikritklab/internal/model"
	"github.com/JesusIslam/sikritklab/internal/response"
	"github.com/asdine/storm/q"
	"github.com/gin-gonic/gin"
)

func DeleteOldThread() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &response.Response{}

		tx, err := database.DB.Begin(true)
		if err != nil {
			resp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		// get all threads
		threads := []*model.Thread{}
		err = tx.All(&threads)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		for _, thread := range threads {
			// get posts of each thread
			posts := []*model.Post{}
			err = tx.Select(
				q.Eq("ThreadID", thread.ID),
			).OrderBy("CreatedAt").Reverse().Find(&posts)
			if err != nil {
				tx.Rollback()
				resp.Error = err.Error()
				c.JSON(http.StatusInternalServerError, resp)
				return
			}

			// delete the thread if last post is older than yesterday
			yesterday := time.Now().Add(-24 * time.Hour)
			if posts[0].CreatedAt.Before(yesterday) {
				err = tx.DeleteStruct(thread)
				if err != nil {
					tx.Rollback()
					resp.Error = err.Error()
					c.JSON(http.StatusInternalServerError, resp)
					return
				}

				// delete all posts from the thread
				for _, post := range posts {
					err = tx.DeleteStruct(post)
					if err != nil {
						tx.Rollback()
						resp.Error = err.Error()
						c.JSON(http.StatusInternalServerError, resp)
						return
					}
				}

				// delete the tags of this thread too
				tags := []*model.Tag{}
				err = tx.Find("ThreadID", thread.ID, &tags)
				if err != nil {
					tx.Rollback()
					resp.Error = err.Error()
					c.JSON(http.StatusInternalServerError, resp)
					return
				}
				for _, tag := range tags {
					err = tx.DeleteStruct(tag)
					if err != nil {
						tx.Rollback()
						resp.Error = err.Error()
						c.JSON(http.StatusInternalServerError, resp)
						return
					}
				}
			}
		}

		tx.Commit()

		c.Next()
	}
}
