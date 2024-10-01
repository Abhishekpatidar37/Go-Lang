package models

type Post struct {
	ID     uint `gorm:"primarykey"`
	Title  string
	Body   string
	UserId uint
}
