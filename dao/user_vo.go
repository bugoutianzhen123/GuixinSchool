package dao

type User struct{
	ID string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	CreatedTime string `gorm:"column:created_time"`
}

func (User) TableName() string {
	return "user"
}