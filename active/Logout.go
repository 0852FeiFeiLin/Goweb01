package active

import (
	"fmt"
	"goWeb01/dao"
	"goWeb01/entity"
	"html/template"
	"net/http"
)

func Logout_Get(w http.ResponseWriter ,r *http.Request) {
	files, err := template.ParseFiles("./view/logout.html")
	if err != nil {
		fmt.Println("页面加载失败")
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("跳转至注销用户页面")
		files.Execute(w, nil)
	}
}
func  Logout_Post(w http.ResponseWriter,r *http.Request){
	//格式化解析
	err := r.ParseForm()
	Removeerr(err)
	//把前端数据传给结构体，然后通过结构体给数据库，数据库删除
	name := r.FormValue("username")
	password := r.FormValue("password")
	user2 := entity.User{
		Name: name,
		Password: password,
	}
	//调用RemoveUser 方法
	_, err = dao.RemoveUser(user2)
	if err != nil {
		fmt.Println("注销用户失败")
		return
	}else{
		fmt.Println("注销用户成功")
	}
	//注销成功跳转页面
	files, err := template.ParseFiles("./view/fanhui.html")
	if err != nil {
		fmt.Println("页面跳转失败")
		return
	}
	files.Execute(w,nil)

}
//错误打包
func Removeerr(err error){
	if err != nil{
		fmt.Println("页面预解析出错啦")
		fmt.Println(err.Error())
		return
	}
}