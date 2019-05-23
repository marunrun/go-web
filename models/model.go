package models

import (
	"reflect"
	"time"
)

type Blog struct {
	Id int `orm:"PK"`
	Title string
	Content string
	Created time.Time
}

func GetLink()  {
	
}

func GetAll() (blogs []Blog)  {

}
