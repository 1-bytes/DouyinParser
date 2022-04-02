package middlewares

import (
	"DouyinParser/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// RequestId 为每个请求标记一个唯一性质的 ID.
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := utils.GetUUID()
		if err != nil {
			uuid = strconv.FormatInt(time.Now().UnixNano(), 10) + utils.GetRandomString(32)
			uuid = utils.Get32MD5Encode(uuid)
		}
		c.Header("request-id", uuid)
		c.Next()
	}
}
