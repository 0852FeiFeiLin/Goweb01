package dao

//操作是数据库

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goWeb01/entity"
)

func ConnertDb() (*sql.DB, error) { //连接数据库
	//导包后连接数据库
	open, err := sql.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/user1?charset= utf8") //数据库是user1
	if err != nil {
		fmt.Println("连接失败啦")
		fmt.Println(err.Error())

	}
	return open, nil //返回一个连接
}

//实例化结构体
//连接数据库查询
func QueryUesr(user entity.User) (*entity.User, error) { //参数和返回值都是uesr 结构体
	fmt.Println("前端传过来的参数", user)

	db, err := ConnertDb() //调用连接数据库函数
	if err != nil {
		return nil, err
	}
	//这个user1 是后端从数据库查询出来的结果
	user1 := entity.User{} //实例化entity里的User结构体

	//写sql 语句
	row := db.QueryRow("select name,password from user where name =? and password = ?", user.Name, user.Password) //条件表达式是结构体里的Name
	//将前端里的属性放在结构体里

	//将查询到的数据导入结构体 row.Scan(将查询的数据导入实例化的结构体)
	err = row.Scan(&user1.Name, &user1.Password) //查询的导入结构体
	fmt.Println("查出来的数据", user1)
	if err != nil {
		fmt.Println("导入失败")
		return nil, err
	}
	return &user1, nil //返回查询到的用户名密码
}

//注册添加功能
func AddUser(user entity.User) (int64, error) {
	fmt.Println("前端传过来的数据：", user)
	//建立数据库连接
	userdb, err := ConnertDb()
	if err != nil {
		return -1, err
	}
	row := userdb.QueryRow("select name,password  from user where name = ? and password = ?", user.Name, user.Password)
	u := entity.User{}
	row.Scan(&u.Name, &u.Password)
	if err != nil {
		fmt.Println("写入变量失败")
		return -1, err
	}
	if u.Name == user.Name {
		fmt.Println("该用户已存在")
		return -2, err
	}
	//sql语句插入数据
	exec, err := userdb.Exec("insert into user (name,password) values (?,?) ", user.Name, user.Password)
	if err != nil {
		return -1, err
	}
	//统计受影响行数
	affected, err := exec.RowsAffected()
	if err != nil {
		return -1, nil
	}
	return affected, err
}

//注销用户功能
func RemoveUser(user entity.User) (int64, error) {
	fmt.Println("前端传的要注销的数据", user)
	db, err := ConnertDb()
	if err != nil {
		fmt.Println("数据库连接失败")
		return -1, err //-1代表失败
	}
	//sql语句
	exec, err := db.Exec("delete from user where name = ? and  password = ?", user.Name, user.Password)
	if err != nil {
		return -1, err //-1代表失败
	}
	//统计受影响行数
	affected, err := exec.RowsAffected()
	if err != nil {
		return -1, err
	}
	return affected, err
}

/*//错误打包
func Errdemo(err error){
	if err != nil{

	}
}*/
