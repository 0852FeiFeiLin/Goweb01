package active

import (
	"fmt"
	"goWeb01/dao"
	"goWeb01/entity"
	"html/template"
	"net/http"
)

//跳转页面
func Register_Get(w http.ResponseWriter, r *http.Request) { //响应数据，请求数据
	files, err := template.ParseFiles("./view/register.html")
	if err != nil {
		fmt.Println("预处理失败！")
		fmt.Println(err.Error())
		return
	} else {
		files.Execute(w, nil) //跳转到预处理页面
		/*记得配置路径哦*/
	}
	//要点进文件里
}
func Register_Post(w http.ResponseWriter, r *http.Request) {
	//预处理前端数据
	err := r.ParseForm()
	if err != nil {
		fmt.Println("数据格式化处理失败")
		fmt.Println(err.Error())
		return
	}
	//把前端数据村给后端
	name := r.FormValue("username")
	password := r.FormValue("password")
	//存进结构体里，通过结构体给后端
	user1 := entity.User{
		Name:     name,
		Password: password,
	}
	//调用add方法，传结构体进去
	user, err := dao.AddUser(user1)
	if user == -2 {
		fmt.Println("该用户已存在，注册失败")
		files, err := template.ParseFiles("./view/error.html")
		if err != nil {
			fmt.Println("返回错误页面失败")
			return
		}
		files.Execute(w, "该用户已存在哦！")
		return
	}
	if err != nil {
		fmt.Println("-----------------")
		fmt.Println("注册失败", err.Error())
		return
	} //如果没注册失败就成功，给程序员看的
	fmt.Println("注册成功")

	//注册成功就跳转到登录页面
	files, err := template.ParseFiles("./view/login.html")
	if err != nil {
		fmt.Println("页面预加载失败")
		fmt.Println(err.Error())
		return
	} else {
		files.Execute(w, nil)
		/*记得有跳转路径*/

	}

}
