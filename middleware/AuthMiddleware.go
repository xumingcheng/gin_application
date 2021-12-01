package middleware

import (

	"github.com/gin-gonic/gin"
	"github.com/xumingcheng/gin_application/common"
	"github.com/xumingcheng/gin_application/model"
	"net/http"
	"strings"

)

func AuthMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		auth:="jiangzhou"
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString==""||!strings.HasPrefix(tokenString,auth+":"){//验证token不为空，并且以：jiangzhou: 为前缀
			c.JSON(http.StatusUnauthorized, gin.H{"code":  http.StatusUnauthorized, "msg": "权限不足"})
			c.Abort()
			return
		}
		index:=strings.Index(tokenString,auth+":") //找到token前缀对应的位置
		//真实token的值
		tokenString= tokenString[index+len(auth)+1:]//截取真实的token（开始位置为：索引开始的位置+关键字符的长度+1(:的长度为1)）
		token, claims, err := common.ParseToken(tokenString)

		if err!=nil ||!token.Valid{//解析错误或者过期等
			c.JSON(http.StatusUnauthorized, gin.H{"code":  http.StatusUnauthorized, "msg": "证书无效"})
			c.Abort()
			return
		}
		userId:=claims.UserId
		var user model.User
		common.DB.First(&user,userId)
		if user.ID==0{//如果没有读取到内容，说明token值有错误
			c.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			c.Abort()
			return
		}
		c.Set("user",user)//将key-value的值存储到context中。
		c.Next()
	}
}

