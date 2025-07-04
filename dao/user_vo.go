package dao

type User struct{
	ID string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

func (User) TableName() string {
	return "user"
}