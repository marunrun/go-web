package models

type User struct {
	Uid         int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name        string
	Profile     *Profile   `orm:"rel(one)"` 	// 一对一的关系
	Post        []*Post `orm:"reverse(many)"` 	// 设置一对多的反向关系
}
