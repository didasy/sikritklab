package handler

import (
	"math/rand"
	"net/http"
	"time"

	"strconv"

	"github.com/JesusIslam/sikritklab/internal/database"
	"github.com/JesusIslam/sikritklab/internal/model"
	"github.com/JesusIslam/sikritklab/internal/response"
	"github.com/gin-gonic/gin"
)

var (
	BaseThreadPath = "/thread/id/"
)

func ThreadRandom(c *gin.Context) {
	resp := &response.Response{}

	rand.Seed(time.Now().UnixNano())

	// get count of threads
	n := 0
	n, err := database.DB.Count(&model.Thread{})
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Message = BaseThreadPath + strconv.FormatInt(rand.Int63n(int64(n)), 10)
	c.JSON(http.StatusOK, resp)
}
