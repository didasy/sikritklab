package handler

import (
	"math/rand"
	"net/http"
	"time"

	"strconv"

	"github.com/JesusIslam/sikritklab/database"
	"github.com/JesusIslam/sikritklab/model"
	"github.com/JesusIslam/sikritklab/response"
	"github.com/labstack/echo"
)

var (
	BaseThreadPath = "/thread/id/"
)

func ThreadRandom(c echo.Context) (err error) {
	resp := &response.Response{}

	rand.Seed(time.Now().UnixNano())

	// get count of threads
	n := 0
	n, err = database.DB.Count(&model.Thread{})
	if err != nil {
		resp.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = BaseThreadPath + strconv.FormatInt(rand.Int63n(int64(n)), 10)
	return c.JSON(http.StatusOK, resp)
}
