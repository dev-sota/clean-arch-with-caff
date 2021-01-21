package entity

// User .
type User struct {
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func (User) TableName() string {
	return "user"
}
