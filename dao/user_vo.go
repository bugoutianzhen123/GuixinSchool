package dao

type User struct{
	ID string `gorm:"column:id;primaryKey"` //ID是学生的学号
	Name string `gorm:"column:name"`
	CreatedTime string `gorm:"column:created_time"`
}

func (User) TableName() string {
	return "user"
}