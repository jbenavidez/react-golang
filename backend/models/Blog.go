package models

type Blog struct {
	ID       uint64 `gorm:"primaryKey"`
	Title    string `gorm:"size:255"`
	Content  string `gorm:"size:255"`
	CreateBy string `gorm:"size:255"`
}

func (Blog) GeTAll() *[]Blog {
	var blogs []Blog
	DB.Find(&blogs)
	return &blogs
}

func Create(title string, content string, createby string) *Blog {
	entry := Blog{Title: title, Content: content, CreateBy: createby}
	DB.Create(&entry)
	return &entry
}
