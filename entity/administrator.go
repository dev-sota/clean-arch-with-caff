package entity

// Administrator .
type Administrator struct {
	ID       int    `gorm:"column:id"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (Administrator) TableName() string {
	return "administrator"
}
