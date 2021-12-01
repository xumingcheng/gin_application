package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xumingcheng/gin_application/common"
	"github.com/xumingcheng/gin_application/model"
	"github.com/xumingcheng/gin_application/response"
	"github.com/xumingcheng/gin_application/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context){
	var requestUser model.User
	c.Bind(&requestUser)
	name:=requestUser.Name
	telephone:=requestUser.Telephone
	password:=requestUser.Password
	//数据验证
	if len(telephone)!=11{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须是11位")
		fmt.Println(telephone,len(telephone))
		return
	}
	if len(password)<6{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码长度不能少于6位")
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(name)==0{
    name=utils.RandomString(10)
	}

	//判断手机号是否存在
    if isTelephoneExit(common.DB,telephone){
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}
	hashPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		response.Response(c,http.StatusInternalServerError,500,nil,"加密错误")
		return
	}
	newUser:=model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hashPassword),
	}
	common.DB.Create(&newUser)
	token,err:=common.ReleaseToken(newUser)
	if err!=nil{
        c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		return
	}
	response.Success(c,gin.H{"token":token,},"注册成功")

}
func Login(c *gin.Context){
	var requestUser model.User
	c.Bind(&requestUser)
	// name:=requestUser.Name
	telephone:=requestUser.Telephone
	password:=requestUser.Password
	//数据验证
	if len(telephone)!=11{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须是11位")
		fmt.Println(telephone,len(telephone))
		return
	}
	if len(password)<6{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码长度不能少于6位")
		return
	}
	var user model.User
	common.DB.Where("telephone=?",telephone).First(&user)
	if user.ID==0{
		c.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":422,
			"msg":"用户不存在",
		})
		return
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"code":500,
			"msg":"密码错误",
		})
		return
	}
	token,err:=common.ReleaseToken(user)
    if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		return
	}
	response.Success(c,gin.H{"token":token},"登陆成功")
}
func Info(c *gin.Context){
	user,_:=c.Get("user")
	response.Success(c,gin.H{"user":response.TouserDto(user.(model.User))},"响应成功")
}
func isTelephoneExit(db *gorm.DB, telephone string) bool {
    var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID!=0{
		return true
	}
	return false
}
