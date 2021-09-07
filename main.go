package main

import (
	"fmt"
	"goWeb01/active"
	"net/http"
	"os"
)

//main 是所有程序的入口
func main() {
	//跳转成功就配置路径  （和Execute 配对）
	http.HandleFunc("/",active.Index) //配置路径  "/"代表根目录是项目启动时自动访问的目录，代码会去找active 目录下的Index方法
	//写了这个方法，参数二就必须是处理器，
	http.HandleFunc("/login",active.Login)//点击发送，就会去找参数2函数，相当于处理的路径
	//设置静态文件路径
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
	//跳转注册页面配置路径
	http.HandleFunc("/register1.html",active.Register_Get)  //点击按钮触发get方法跳转页面
	http.HandleFunc("/register",active.Register_Post)

	//注销用户页面
	http.HandleFunc("/remove.html",active.Logout_Get)
	http.HandleFunc("/logout",active.Logout_Post)

	//新建后端服务器  参数：（地址）  监听器
	err := http.ListenAndServe(":8081", nil)
	//处理报错
	if err != nil{
		fmt.Println("服务器创建失败")
		fmt.Println(err.Error())
		//如果创建失败，关闭服务器
		os.Exit(0) //参数为0 代表退出
		/*服务器测试是否成功，浏览器输入localhost:8090*/
		//如果成功就会跳转到我们的页面
	}




}