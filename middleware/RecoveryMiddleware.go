package middleware
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xumingcheng/gin_application/response"
)

func RecoverMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if err:=recover();err!=nil{
				response.Fail(c,nil,fmt.Sprint(err))
				c.Abort()
				return
			}
		}()
	}
}
