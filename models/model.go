package models

import (
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
