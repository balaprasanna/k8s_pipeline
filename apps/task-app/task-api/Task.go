package main

// Task ...
// This is our Task Model
type Task struct {
	Id      int64  `gorm:"primary_key,column:id"`
	Name    string `gorm:"column:name" json:"Name" binding:"required"`
	Type    string `gorm:"column:type" json:"Type" binding:"required"`
	Status  uint   `gorm:"column:status"`
	Content string `gorm:"column:content" json:"Content" binding:"required"`
}

// TableName ...
// This is an extension to rename the table name to `task`.
func (Task) TableName() string {
	return "task"
}
