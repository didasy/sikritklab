package custommiddleware

import (
	"net/http"
	"os"

	"github.com/JesusIslam/sikritklab/internal/constant"
	"github.com/JesusIslam/sikritklab/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/haisum/recaptcha"
)

var (
	Secret string
	re     recaptcha.R
)

func init() {
	Secret = os.Getenv(constant.EnvRecaptchaSecret)

	re = recaptcha.R{
		Secret: Secret,
	}
}

func CheckCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &response.Response{}

		if os.Getenv(constant.EnvEnableRecaptcha) == "true" {
			// must be sent as post form g-recaptcha-response
			isValid := re.Verify(*c.Request)
			if !isValid {
				resp.Error = constant.ErrorInvalidCaptcha
				c.JSON(http.StatusForbidden, resp)
				return
			}
		}

		c.Next()
	}
}
