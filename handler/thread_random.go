package handler

import (
	"math/rand"
	"net/http"
	"time"

	"strconv"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

var (
	BaseThreadPath = "/thread/id/"
)

func ThreadRandom(c echo.Context) (err error) {
	resp := &response.Response{}

	rand.Seed(time.Now().UnixNano())

	db := database.New()
	var n int64
	err = db.Table("threads").Count(&n).Error
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusServiceUnavailable, resp)
	}
	if n < 1 {
		resp.Error = "No thread found"
		return c.JSON(http.StatusNotFound, resp)
	}

	resp.Message = BaseThreadPath + strconv.FormatInt(rand.Int63n(n), 10)
	return c.JSON(http.StatusOK, resp)
}
