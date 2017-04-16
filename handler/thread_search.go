package handler

import (
	"net/http"
	"regexp"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/form"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/labstack/echo"
)

// Default order by created_at desc
func ThreadSearch(c echo.Context) (err error) {
	resp := &response.Response{}

	search := form.GetSearchForm(c)

	tx, err := database.DB.Begin(false)
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	searched := false
	threads := []*model.Thread{}

	if search.Title != "" {
		// get all threads
		tmp := []*model.Thread{}
		err = tx.AllByIndex("CreatedAt", &tmp, storm.Reverse())
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}

		// check and collect with regex
		searchRegex := regexp.MustCompile(search.Title)
		for _, thread := range tmp {
			if searchRegex.MatchString(thread.Title) {
				threads = append(threads, thread)
			}
		}

		searched = true
	}

	if len(search.Tags) > 0 {
		// get all tags first
		tags := []*model.Tag{}
		err = tx.Select(
			q.In("Tag", search.Tags),
		).Find(&tags)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}

		// collect the thread ID
		threadIDMap := map[string]bool{}
		for _, tag := range tags {
			threadIDMap[tag.ThreadID] = true
		}

		// turn into slice
		threadIDs := []string{}
		for id, _ := range threadIDMap {
			threadIDs = append(threadIDs, id)
		}

		// search all threads with these ids
		err = tx.Select(
			q.In("ID", threadIDs),
		).OrderBy("CreatedAt").Reverse().Skip(search.Page).Limit(search.PerPage).Find(&threads)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}

		searched = true
	}

	if !searched {
		// get the threads normally if no search
		err = tx.Select(nil).OrderBy("CreatedAt").Reverse().Skip(search.Page).Limit(search.PerPage).Find(&threads)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}
	}

	threadPosts := []*model.ThreadPost{}
	for _, thread := range threads {
		// then get first posts of each one
		posts := []*model.Post{}
		err = tx.Select(
			q.Eq("ThreadID", thread.ID),
		).OrderBy("CreatedAt").Limit(1).Find(&posts)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}

		// then get the tags
		tags := []*model.Tag{}
		err = tx.Find("ThreadID", thread.ID, &tags)
		if err != nil {
			tx.Rollback()
			resp.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, resp)
		}

		// then combine them
		threadPosts = append(threadPosts, &model.ThreadPost{
			Thread: thread,
			Posts:  posts,
			Tags:   tags,
		})
	}

	tx.Commit()

	// resp.Message = threadPosts
	return c.JSON(http.StatusOK, resp)
}
