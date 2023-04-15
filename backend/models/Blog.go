package models

import "fmt"

type Blog struct {
	ID       uint64 `gorm:"primaryKey;" json:"id"`
	Title    string `gorm:"size:255;" json:"title"`
	Content  string `gorm:"size:255;" json:"content"`
	CreateBy string `gorm:"size:255;" json:"create_by"`
}

func GeTAll() *[]Blog {
	var blogs []Blog
	DB.Find(&blogs)
	return &blogs
}

func Create(title string, content string, createby string) *Blog {
	entry := Blog{Title: title, Content: content, CreateBy: createby}
	DB.Create(&entry)
	fmt.Println("the_entry", entry)
	return &entry
}

func GetBlog(id uint64) *Blog {
	var blog Blog
	DB.Where("id = ? ", id).First(&blog)
	return &blog
}

func DeleteBlog(id uint64) {
	//delete blog
	DB.Where("id = ? ", id).Delete(&Blog{})
}
