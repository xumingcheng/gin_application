package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xumingcheng/gin_application/common"
	"github.com/xumingcheng/gin_application/route"
	"os"
)

func main() {
  InitConfig()
  common.InitDB()
  r:=gin.Default()
  r=route.CollectRoute(r)
  port:=viper.GetString("server.port")
  if port!=""{
	  r.Run(":"+port)
  }else {
	  r.Run()
  }

}
func InitConfig(){
	workdir,_:=os.Getwd()//获取目录的相应的路径
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdir+"/config")
	fmt.Println(workdir)
	err:=viper.ReadInConfig()//加载文件
	if err!=nil{
		panic(err)
	}
}

