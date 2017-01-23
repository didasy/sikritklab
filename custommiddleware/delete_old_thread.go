package custommiddleware

import (
	"net/http"
	"time"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

func DeleteOldThread(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		resp := &response.Response{}

		tx := database.New().Begin()

		// get all posts
		posts := []*model.Post{}
		err = tx.Table("posts").Find(&posts).Error
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusServiceUnavailable, resp)
		}

		// get last post of each thread
		// thread_id : post
		threadPostMap := map[int64]*model.Post{}
		for _, post := range posts {
			// if already exists, check if created_at is newer
			// if newer, replace
			if lastPost := threadPostMap[post.ThreadID]; lastPost != nil {
				if lastPost.CreatedAt.Before(post.CreatedAt) {
					threadPostMap[post.ThreadID] = post
				}
			} else {
				threadPostMap[post.ThreadID] = post
			}
		}

		// from each thread, check which has last post more than 24h ago
		threadsToBeDeleted := []int64{}
		yesterday := time.Now().Add(-24 * time.Hour)
		for threadID, post := range threadPostMap {
			if post.CreatedAt.Before(yesterday) {
				threadsToBeDeleted = append(threadsToBeDeleted, threadID)
			}
		}

		if len(threadsToBeDeleted) > 0 {
			// delete all of those threads
			err = tx.Table("threads").Where("id IN (?)", threadsToBeDeleted).Unscoped().Delete(&model.Thread{}).Error
			if err != nil && err != database.NotFound {
				tx.Rollback()
				resp.Error = err.Error()
				return c.JSON(http.StatusServiceUnavailable, resp)
			}

			// delete all posts of those threads
			err = tx.Table("posts").Where("thread_id in (?)", threadsToBeDeleted).Unscoped().Delete(&model.Post{}).Error
			if err != nil && err != database.NotFound {
				tx.Rollback()
				resp.Error = err.Error()
				return c.JSON(http.StatusServiceUnavailable, resp)
			}
		}

		tx.Commit()

		return next(c)
	}
}
