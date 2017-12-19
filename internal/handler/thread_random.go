package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/asdine/storm"

	"github.com/JesusIslam/sikritklab/internal/database"
	"github.com/JesusIslam/sikritklab/internal/model"
	"github.com/JesusIslam/sikritklab/internal/response"
	"github.com/gin-gonic/gin"
)

func ThreadRandom(c *gin.Context) {
	resp := &response.Response{}

	rand.Seed(time.Now().UnixNano())

	// get all threads
	threads := []*model.Thread{}

	// get count of threads
	err := database.DB.AllByIndex("CreatedAt", &threads, storm.Reverse())
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	lengthOfThreads := len(threads)
	randomIndex := rand.Int63n(int64(lengthOfThreads))
	randomThreadID := threads[randomIndex].ID

	resp.Message = randomThreadID
	c.JSON(http.StatusOK, resp)
}
