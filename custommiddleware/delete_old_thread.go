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

		db := database.New()

		yesterday := time.Now().Add(-24 * time.Hour)
		err = db.Table("threads").Where("created_at < ?", yesterday).Unscoped().Delete(&model.Post{}).Error
		if err != nil && err != database.NotFound {
			resp.Error = err.Error()
			return c.JSON(http.StatusServiceUnavailable, resp)
		}

		return next(c)
	}
}
