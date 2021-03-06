package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/varunbhayana/rate-limiting/service"
	"github.com/varunbhayana/rate-limiting/util/cycle_util"
)

type RedisModel struct {
	Time  int64 `json:"Time"`
	Count int64 `json:"Count"`
}

func RateLimit() func(*gin.Context) {
	return func(c *gin.Context) {
		cycle_util.DegdCall(
			1*time.Second,
			c,
			func() (int, interface{}) {
				userId := c.GetHeader("user-id")
				applicationId := c.GetHeader("application-id")
				if userId == "" || applicationId == "" {
					return 400, "Bad Request"
				}

				return service.RateLimiter.RateLimit(userId, applicationId)
			},
		)
	}

}
