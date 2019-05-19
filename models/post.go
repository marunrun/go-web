package models

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` 	//设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`   //设置多对多关系
}

