package active

import (

	"fmt"
	"goWeb01/dao"
	"goWeb01/entity"
	"html/template"
	"net/http"
)

//新建页面 跳转页面
/*处理器   用来处理前端传过来的信息，还可以把后台的信息串到前端*/
func Index(w http.ResponseWriter ,r *http.Request){ //固定参数 参数一：响应数据  参数二：后台接收到的数据
	files, err := template.ParseFiles("./view/login.html")//参数是用来预处理加载页面
		                            //代表进服务器跳转到的页面
	if err != nil{
		fmt.Println("页面预加载失败")
		w.Write([]byte("出错了！未打开登录页面"))
		/*w就是响应，用于后台数据响应到前端去*/
		fmt.Println(err.Error())
		return
	}else{ //没错就成功跳转
		             /*就是在这里面写东西，然后显示在模板里*/
		files.Execute(w,nil)//跳转到预处理页面
		//下一步配置跳转路径（让服务器知道我们要跳转的地方）
	}
}
//Login方法实现登录逻辑（目的把前端的数据和后台查询到的数据进行比较）
func Login(w http.ResponseWriter, r *http.Request ){ //固定参数：后参数1后台响应给前端的数据，参数2是后台接收到前端的数据
	/* r  可以获取前端 ， w可以发送信息到前端 */
	err := r.ParseForm()  //格式化解析表单数据
	//判断错误（从前端串的数据）
	if err != nil {
		//如果错误，就跳转但错误页面
		//files, _ := template.ParseFiles("./view/error.html")
		fmt.Println("数据格式化处理失败")
		fmt.Println(err.Error())
		return
	}
	//r 获取前端信息
	name := r.FormValue("username") //通过前端的name关键字获取前端数据
	password := r.FormValue("password")
	if name == "" || password == ""{
		files, _ := template.ParseFiles("./view/error.html")
		files.Execute(w,"用户名或密码不能为空哦")
		return
	}
	//获取完前端数据 接收（结构体存）
	user := entity.User{Name:name,Password: password}//把前端数据封装进了后端结构体


	//前端传给结构体后，就要调用dao的方法传参进去，结构体就是参数
	queryuesr, err2 := dao.QueryUesr(user) //调用dao的 QueryUser的方法
	if err2 != nil {  //前端数据错误，就查询不到数据库数据
		files, _ := template.ParseFiles("./view/error.html")
		files.Execute(w,"登录失败啦！请检查用户名密码")
		fmt.Println("查询失败")
		fmt.Println("该用户密码错误or该用户不存在")
		fmt.Println(err2.Error())
		return
	}//如果成功
	//把前端后端密码做比较    （如果数据库没有查询到数据，他是直接不走这个判断的，所以输出错误在上面查询失败哪里输出）
	if user.Name == queryuesr.Name && user.Password == queryuesr.Password{
		//判断前端用户名密码和数据库查出来的是否一致
		files, _ := template.ParseFiles("./view/page1.html")
		files.Execute(w,"登录成功")

		fmt.Println("登录成功")
		return
	}else{

		fmt.Println("登录失败，检查用户名密码")
		return
	}
	//前端传的数据==在user结构体里面，后端
}
