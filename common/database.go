package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func InitDB()*gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.Get("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",username,password,host,port,database,charset,url.QueryEscape(loc))
	fmt.Println(args)
	db,err:=gorm.Open(mysql.Open(args),&gorm.Config{})
    if err!=nil{
		panic("failed to connect to err:"+err.Error())
	}
     DB=db
	 return db
}