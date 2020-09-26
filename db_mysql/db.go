package db_mysql

import (
	"BeegoProject123/models"
	"astaxie/beego"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db  * sql.DB

func init(){

	config :=beego.AppConfig

	dbDriver :=config.String("db_driverName")
	dbUser :=config.String("db_user")
	dbPwd := config.String("db_pwd")
	dbIp := config.String("db_ip")
	dbPort :=config.String("da_pore")
	dbName :=config.String("da_name")
	conmint :=dbUser+":"+dbPwd+"@tcp("+dbIp+":"+dbPort+")/"+dbName+"?charset=utf8"
	fmt.Println(conmint)
	DB,err:=sql.Open(dbDriver,conmint)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("成功")
	Db = DB

}

func  InserUser(user models.Info)(int64 ,error){
	// 将用户密码进行hash ，使用md5 计算hash值
	hashMd5:= md5.New()
	hashMd5.Write([]byte(user.Pwd))
	bytes :=hashMd5.Sum(nil)
	user.Pwd = hex.EncodeToString(bytes)
	fmt.Println("用户名：",user.Nick,"密码:",user.Pwd)

	res ,err :=Db.Exec("insert into user(nick,pwd)  values(?,?)",user.Nick,user.Pwd)
	if err !=nil{
		return -1 ,err
	}
	id ,err :=res.RowsAffected()
	if err !=nil{
		return -1 ,err
	}
	return id , nil

}



// 查询
func QueryUser(){
	Db.QueryRow("select * from Beego_text")
}