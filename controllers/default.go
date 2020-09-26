package controllers

import (
	"BeegoProject123/db_mysql"
	"BeegoProject123/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller // 匿名字段
}

/*
 方法的重写
 */

// Get类型的请求与数据解析

func (c *MainController) Get() {

	// 1. 获取请求数据
	c.GetString("user") // 返回字符串
	usre :=c.Ctx.Input.Query("user")
	pwd := c.Ctx.Input.Query("pwd")

	// 2.使用固定数据进行校验
	if usre != "sianernian" || pwd !="123"{
		// 代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("数据错误"))
		c.Ctx.WriteString(" 错误")
	}
	// 校验正确
	c.Ctx.ResponseWriter.Write([]byte("数据成功"))

	//c.Data["Website"] = "www.baidu.com"
	//c.Data["Email"] = "1916430401@qq.com"
	//c.TplName = "index.tpl"


	//c.Ctx.WriteString("hello")
	//c.Ctx.WriteString("你好")
}



// func (c *MainController)Post(){
//
//	// 1.接收 post请求的参数
//	name :=c.Ctx.Request.FormValue("name")
//	age :=c.Ctx.Request.FormValue("age")
//	sex :=c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//
//	// 2 . 进行数据校验
//	if name!="zhangzian" || age !="18"{
//		c.Ctx.ResponseWriter.Write([]byte("失败"))
//		return
//	}
//	c.Ctx.WriteString("成功")
//
//	for i :=0; i<10 ; i++  {
//		fmt.Printf("第%d次打印\n",i)
//	}
//}


//func (c *MainController)Post(){
//	//1. 解析 前端 json 格式
//
//	var person models.Person
//                            // 读取数据  字节切片
//	dataByte , err :=ioutil.ReadAll(c.Ctx.Request.Body)
//
//	if err !=nil {
//		c.Ctx.WriteString("解析错误")
//		return
//	}
//	                    // 字节切片
//	err =json.Unmarshal(dataByte,&person) //从前端上解析 json
//	if err !=nil{
//		c.Ctx.WriteString("失败")
//		return
//	}
//	fmt.Println("姓名",person.Name)
//	fmt.Println("年龄",person.Age)
//	fmt.Println("性别",person.Sex)
//	c.Ctx.WriteString("成功")
//
//0-

func (c *MainController)Post(){
	fmt.Println(c == nil)
	fmt.Println(c.Ctx == nil)
	fmt.Println(c.Ctx.Request == nil)

	// 解析前端 json 数据

	     // 读取数据
	dataByte , err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("失败")
		return
	}
	var info models.Info

	err =json.Unmarshal(dataByte,&info) // 解析json
	if err !=nil{
		c.Ctx.WriteString("失败")
		return
	}
	//fmt.Println("姓名：",info.Name)
	//fmt.Println("生日：",info.Birthday)
	//fmt.Println("地址：",info.Address)
	//fmt.Println("昵称：",info.Nick)

	id ,err :=db_mysql.InserUser(info)
	if err !=nil{
		fmt.Println(err.Error())
		c.Ctx.WriteString("错误")
		return
	}
	fmt.Println(id)


	result :=models.ResponseResult{
		Code:    0,
		Message: "成功",
		Data:    nil,
	}
	c.Data["json"] = &result
	c.ServeJSON()

	c.Ctx.WriteString("成功")

}