package custommiddleware

import (
	"net/http"
	"os"

	"github.com/JesusIslam/sikritklab/response"
	"github.com/haisum/recaptcha"
	"github.com/labstack/echo"
)

var (
	Secret string
	re     recaptcha.R
)

func init() {
	Secret = os.Getenv("RECAPTCHA_SECRET")

	re = recaptcha.R{
		Secret: Secret,
	}
}

func CheckCaptcha(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		resp := &response.Response{}

		// must be sent as post form g-recaptcha-response
		isValid := re.Verify(*c.Request())
		if !isValid {
			resp.Error = "Invalid captcha"
			return c.JSON(http.StatusForbidden, resp)
		}

		return next(c)
	}
}
