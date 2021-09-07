package entity
//存放实体类的，数据表里面每一列对应一个属性

type User struct{ //首字母大写，表示可以在整个项目使用（java 里public）
	id int
	Name string  //（对应数据库字段）
	Password string  //（和user 表对应）

}
